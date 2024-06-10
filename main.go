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
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatalf("error loading config; %v", err)
	}
	influx.SetInfluxWriter(cfg.Influx)
	var stats1 = map[string]*sys.Stats{}
	for _, block := range cfg.Blocks {
		stats1[block.Name] = sys.GetStats(cfg.Connect, block.Name)
	}
	time.Sleep(10 * time.Second)
	var stats2 = map[string]*sys.Stats{}
	for _, block := range cfg.Blocks {
		stats2[block.Name] = sys.GetStats(cfg.Connect, block.Name)
	}
	for block1, stat1 := range stats1 {
		for block2, stat2 := range stats2 {
			if block1 == block2 {
				diff := stat1.Diff(stat2)
				log.Printf("Diff for: %s\n", block2)
				diff.Output()
				influx.AddIoStat(block2, map[string]interface{}{
					"ReadsCompleted":  diff.ReadsCompleted,
					"ReadTime":        diff.ReadTime,
					"WritesCompleted": diff.WritesCompleted,
					"WriteTime":       diff.WriteTime,
					"IOTime":          diff.IOTime,
					"WeightedIOTime":  diff.WeightedIOTime,
				})
			}
		}
	}
	influx.Flush()
}
