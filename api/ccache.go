package api

import (
	"net/http"

	"github.com/raintank/metrictank/api/middleware"
	"github.com/raintank/metrictank/api/models"
	"github.com/raintank/metrictank/api/response"
)

func (s *Server) ccacheDelete(ctx *middleware.Context, req models.CCacheDelete) {
	resp := models.NewCCacheDeleteResp()

	for _, pattern := range req.Patterns {
		nodes, err := s.MetricIndex.Find(req.OrgId, pattern, 0)
		if err != nil {
			response.Write(ctx, response.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		for _, node := range nodes {
			for _, def := range node.Defs {
				res := s.Cache.DelMetric(def.Id)
				resp.Nodes = append(resp.Nodes, res.Deleted...)
			}
		}
	}
	response.Write(ctx, response.NewMsgp(200, resp))
}
