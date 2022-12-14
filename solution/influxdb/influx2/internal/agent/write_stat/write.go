package write_stat

import (
	"context"
	"influx2/config"
	"influx2/internal/app"
	"influx2/internal/stat_map"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/yiaw/yiaw-go/log"
)

type writeAgent struct {
	cfg config.InfluxDB
}

func (wa *writeAgent) Run(client influxdb2.Client) error {
	log.INF("Welcome Write Stat Item ...\n")

	WriteAPI := client.WriteAPIBlocking(wa.cfg.OrgName, wa.cfg.Bucket)

	allStat := stat_map.AllMakeStatInfo()
	timer := time.NewTicker(time.Second * 10)
	i := 0
	for {
		select {
		case <-timer.C:
			startTime := time.Now()
			// Insert ReqResp
			var points []*write.Point

			for _, stat := range allStat {
				stat.FakeData()
				points = append(points, stat.NewPoint()...)
				stat.Clear()
			}

			err := WriteAPI.WritePoint(context.TODO(), points...)
			if err != nil {
				log.ERR("insert stat failed. err=%v\n", err)
			}

			endTime := time.Since(startTime)
			elapsedTime := float64(endTime.Milliseconds()) / float64(1000)
			log.INF("[%d] points count:%d, elapsed:%0.3fms\n", i, len(points), elapsedTime)
		}
		i++
	}
}

func NewAgentService(cfg config.InfluxDB) app.AgentService {
	return &writeAgent{cfg: cfg}
}
