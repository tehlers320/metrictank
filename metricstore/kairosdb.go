package metricstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ctdk/goas/v2/logger"
	"github.com/raintank/raintank-metric/metricdef"
	"net/http"
	"time"
)

// Kairosdb client
type Kairosdb struct {
	client *http.Client
	host   string
}

func NewKairosdb(host string) (*Kairosdb, error) {
	logger.Debugf("initializing kairosdb client to %s", host)
	return &Kairosdb{
		client: &http.Client{Timeout: (10 * time.Second)},
		host:   host,
	}, nil
}

// Datapoint instances are persisted back to kairosdb via AddDatapoints
type Datapoint struct {
	Name      string            `json:"name"`
	Timestamp int64             `json:"timestamp"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

func (kdb *Kairosdb) SendMetrics(metrics *[]metricdef.IndvMetric) error {
	// marshal metrics into datapoint structs
	datapoints := make([]Datapoint, len(*metrics))
	for i, m := range *metrics {
		tags := make(map[string]string, len(m.Extra))
		for k, v := range m.Extra {
			tags[k] = fmt.Sprintf("%v", v)
		}
		tags["org_id"] = fmt.Sprintf("%v", m.OrgId)
		datapoints[i] = Datapoint{
			Name:      m.Metric,
			Timestamp: m.Time * 1000,
			Value:     m.Value,
			Tags:      tags,
		}
	}
	err := kdb.AddDatapoints(datapoints)
	if err != nil {
		logger.Infof("failed to send metrics to kairosdb -- retrying")
		// start a ticker and a goroutine to keep trying to submit the
		// datapoints
		ticker := time.NewTicker(time.Second * time.Duration(30))
		go func(t *time.Ticker, kdb *Kairosdb, datapoints []Datapoint) {
			// TODO: have this write out or somehow save the
			// outstanding data if we're shut down
			for range t.C {
				e := kdb.AddDatapoints(datapoints)
				if e == nil {
					// we're done
					t.Stop()
					logger.Infof("saved delayed datapoints to kairosdb")
					return
				} else {
					logger.Debugf("failed to save outstanding datapoints to kairosdb again - message was: %s", e.Error())
				}
			}
		}(ticker, kdb, datapoints)
		return err
	}
	return nil
}

// AddDatapoints add datapoints to configured kairosdb instance
func (kdb *Kairosdb) AddDatapoints(datapoints []Datapoint) error {

	json, err := json.Marshal(datapoints)
	if err != nil {
		return err
	}
	resp, err := kdb.client.Post(kdb.host+"/api/v1/datapoints", "application/json", bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	if resp.Status != "204 No Content" {
		return errors.New("Response was non-200: " + resp.Status)
	}
	return nil
}

func (kdb *Kairosdb) Type() string {
	return "Kairosdb"
}
