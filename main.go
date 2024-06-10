package main

import (
	"github.com/jchavannes/sysstat/sys"
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	config, err := sys.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var stats1 = map[string]*sys.Stats{}
	for _, block := range config.Blocks {
		stats1[block.Name] = sys.GetStats(config.Connect, block.Name)
	}
	time.Sleep(10 * time.Second)
	var stats2 = map[string]*sys.Stats{}
	for _, block := range config.Blocks {
		stats2[block.Name] = sys.GetStats(config.Connect, block.Name)
	}
	for block1, stat1 := range stats1 {
		for block2, stat2 := range stats2 {
			if block1 == block2 {
				diff := stat1.Diff(stat2)
				log.Printf("Diff for: %s\n", block2)
				diff.Output()
			}
		}
	}
}
