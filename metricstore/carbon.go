package metricstore

import (
	"github.com/ctdk/goas/v2/logger"
	"github.com/marpaia/graphite-golang"
	"github.com/raintank/raintank-metric/metricdef"
	"strconv"
	"time"
)

// Kairosdb client
type Carbon struct {
	Graphite *graphite.Graphite
}

func NewCarbon(host string, port int) (*Carbon, error) {
	graphite, err := graphite.NewGraphite(host, port)
	if err != nil {
		return nil, err
	}
	return &Carbon{Graphite: graphite}, nil
}

func (carbon *Carbon) SendMetrics(metrics *[]metricdef.IndvMetric) error {
	// marshal metrics into datapoint structs
	datapoints := make([]graphite.Metric, len(*metrics))
	for i, m := range *metrics {
		datapoints[i] = graphite.Metric{
			Name:      m.Id,
			Timestamp: m.Time,
			Value:     strconv.FormatFloat(m.Value, 'f', -1, 64),
		}
	}
	err := carbon.Graphite.SendMetrics(datapoints)

	if err != nil {
		logger.Infof("failed to send metrics to carbon -- retrying")
		ticker := time.NewTicker(time.Second * time.Duration(30))
		go func(t *time.Ticker, carbon *Carbon, datapoints []graphite.Metric){
			// see TODO in metricstore/kairosdb.go
			for range t.C {
				e := carbon.Graphite.SendMetrics(datapoints)
				if e == nil {
					// we're done
					t.Stop()
					logger.Infof("saved delayed datapoints to carbon")
					return
				} else {
					logger.Debugf("failed to save outstanding datapoints to carbon again - message was: %s", e.Error())
				}
			}
		}(ticker, carbon, datapoints)
		return err
	}

	return nil
}

func (carbon *Carbon) Type() string {
	return "Carbon"
}
