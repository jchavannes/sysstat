package influx

func AddIoStat(block string, metricValues map[string]interface{}) {
	writer := getInfluxWriter()
	writer.Write(Point{
		Measurement: NameIoStat,
		Fields:      metricValues,
		Tags: map[string]string{
			"block": block,
		},
	})
}
