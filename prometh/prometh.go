package prometh

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func Push(pushAddr string, name string, epoch int) (err error) {
	jobName := fmt.Sprintf("filscan_syncer_delay_second")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: jobName})
	gauge.Set(float64(epoch))
	err = push.New(pushAddr, jobName).Grouping("module", "filscan").Grouping("syncer", name).Collector(gauge).Push()

	return err

}
