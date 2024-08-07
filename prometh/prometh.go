package prometh

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func Push(pushAddr string, name string, delay int64) (err error) {
	jobName := fmt.Sprintf("filscan_syncer_delay_height")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: jobName})
	gauge.Set(float64(delay))
	err = push.New(pushAddr, jobName).Grouping("module", "filscan").Grouping("syncer", name).Collector(gauge).Push()

	return err
}
