package api

import (
	"net/http"
	"sync"

	"github.com/raintank/metrictank/api/middleware"
	"github.com/raintank/metrictank/api/models"
	"github.com/raintank/metrictank/api/response"
	"github.com/raintank/metrictank/cluster"
	"github.com/raintank/worldping-api/pkg/log"
)

func (s *Server) ccacheDelete(ctx *middleware.Context, req models.CCacheDelete) {
	resp := models.CCacheDeleteResp{}

	if req.Propagate {
		deletedSeries, deletedArchives, errors := s.ccacheDeletePropagate(&req)
		resp.DeletedSeries = deletedSeries
		resp.DeletedArchives = deletedArchives
		resp.PeerErrors = errors
	}

	for _, pattern := range req.Patterns {
		nodes, err := s.MetricIndex.Find(req.OrgId, pattern, 0)
		if err != nil {
			response.Write(ctx, response.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		for _, node := range nodes {
			for _, def := range node.Defs {
				res := s.Cache.DelMetric(def.Id)
				resp.DeletedSeries += res.Series
				resp.DeletedArchives += res.Archives
			}
		}
	}
	response.Write(ctx, response.NewMsgp(200, resp))
}

func (s *Server) ccacheDeletePropagate(req *models.CCacheDelete) (int, int, int) {
	// we never want to propagate more than once to avoid loops
	req.Propagate = false

	peers := cluster.Manager.MemberList()
	var deletedSeries, deletedArchives, errors int = 0, 0, 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, peer := range peers {
		if peer.IsLocal() {
			continue
		}
		wg.Add(1)
		go func(peer cluster.NodeIf) {
			mu.Lock()
			defer mu.Unlock()
			deletedPeerSeries, deletedPeerArchives, err := s.ccacheDeleteRemote(req, peer)
			if err != nil {
				errors++
			} else {
				deletedSeries += deletedPeerSeries
				deletedArchives += deletedPeerArchives
			}
			wg.Done()
		}(peer)
	}
	wg.Wait()

	return deletedSeries, deletedArchives, errors
}

func (s *Server) ccacheDeleteRemote(req *models.CCacheDelete, peer cluster.NodeIf) (int, int, error) {
	log.Debug("HTTP metricDelete calling %s/ccache/delete", peer.GetName())
	buf, err := peer.Post("/ccache/delete", *req)
	if err != nil {
		log.Error(4, "HTTP ccacheDelete error querying %s/ccache/delete: %q", peer.GetName(), err)
		return 0, 0, err
	}

	resp := models.CCacheDeleteResp{}
	buf, err = resp.UnmarshalMsg(buf)
	if err != nil {
		log.Error(4, "HTTP ccacheDelete error unmarshaling body from %s/ccache/delete: %q", peer.GetName(), err)
		return 0, 0, err
	}

	return resp.DeletedSeries, resp.DeletedArchives, nil
}
