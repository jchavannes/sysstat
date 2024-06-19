package main

import (
	"github.com/jchavannes/sysstat/config"
	"github.com/jchavannes/sysstat/influx"
	"github.com/jchavannes/sysstat/sys"
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	cfg, err := config.GetConfig(".config/sysstat/config.yaml")
	if err != nil {
		log.Fatalf("error loading config; %v", err)
	}
	hasInflux := cfg.Influx.Url != ""
	if hasInflux {
		influx.SetInfluxWriter(cfg.Influx)
	}
	ticker := time.NewTicker(60 * time.Second)
	var statsPrev = make(map[string]*sys.Stats)
	log.Printf("Starting sysstat\n")
	for loop := 0; ; loop++ {
		var statsNew = make(map[string]*sys.Stats)
		for _, block := range cfg.Blocks {
			statsNew[block.Name], err = sys.GetStats(cfg.Connect, block.Name)
			if err != nil {
				log.Fatalf("fatal error getting stats for block: %s; %v\n", block.Name, err)
			}
		}
		for blockNew, statNew := range statsNew {
			for blockPrev, statPrev := range statsPrev {
				if blockNew == blockPrev {
					diff := statPrev.Diff(statNew)
					if hasInflux || loop%10 == 1 {
						log.Printf("Diff for: %s\n", blockNew)
						diff.Output()
					}
					if !hasInflux {
						continue
					}
					influx.AddIoStat(blockNew, map[string]interface{}{
						"ReadsCompleted":           diff.ReadsCompleted,
						"ReadTime":                 diff.ReadTime,
						"SectorsRead":              diff.SectorsRead,
						"WritesCompleted":          diff.WritesCompleted,
						"WriteTime":                diff.WriteTime,
						"SectorsWritten":           diff.SectorsWritten,
						"IOTime":                   diff.IOTime,
						"WeightedIOTime":           diff.WeightedIOTime,
						"ReadsCompletedPerSecond":  float32(diff.ReadsCompleted) / 60.0,
						"ReadTimePercent":          float32(diff.ReadTime) / 60000.0,
						"SectorsReadPerSecond":     float32(diff.SectorsRead) / 60.0,
						"WritesCompletedPerSecond": float32(diff.WritesCompleted) / 60.0,
						"WriteTimePercent":         float32(diff.WriteTime) / 60000.0,
						"SectorsWrittenPerSecond":  float32(diff.SectorsWritten) / 60.0,
						"IOTimePercent":            float32(diff.IOTime) / 60000.0,
						"WeightedIOTimePercent":    float32(diff.WeightedIOTime) / 60000.0,
					})
				}
			}
		}
		if hasInflux {
			influx.Flush()
		}
		statsPrev = statsNew
		<-ticker.C
	}
}
