package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/jchavannes/sysstat/config"
	"time"
)

const (
	NameIoStat = "io_stat"

	TagSource = "source"
)

type Point struct {
	Measurement string
	Fields      map[string]interface{}
	Tags        map[string]string
}

type Writer struct {
	Api    api.WriteAPI
	Config config.Influx
}

func (i *Writer) Write(p Point) {
	point := influxdb2.NewPoint(p.Measurement, p.Tags, p.Fields, time.Now())
	if i.Config.Source != "" {
		point.AddTag(TagSource, i.Config.Source)
	}
	i.Api.WritePoint(point)
}

var _influxWriter *Writer

func SetInfluxWriter(influxConfig config.Influx) {
	c := influxdb2.NewClient(influxConfig.Url, influxConfig.Token)
	_influxWriter = &Writer{
		Api:    c.WriteAPI(influxConfig.Org, influxConfig.Bucket),
		Config: influxConfig,
	}
}

func getInfluxWriter() *Writer {
	return _influxWriter
}

func Flush() {
	getInfluxWriter().Api.Flush()
}
