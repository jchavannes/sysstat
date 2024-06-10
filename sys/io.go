package sys

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	ReadsCompleted    int64
	ReadsMerged       int64
	SectorsRead       int64
	ReadTime          int64
	WritesCompleted   int64
	WritesMerged      int64
	SectorsWritten    int64
	WriteTime         int64
	IOInProgress      int64
	IOTime            int64
	WeightedIOTime    int64
	DiscardsCompleted int64
	DiscardsMerged    int64
	SectorsDiscarded  int64
	DiscardTime       int64
	FlushesCompleted  int64
	FlushTime         int64
}

func (s *Stats) Output() {
	ps("METRIC", "UNITS", "VALUE")
	pm("ReadsCompleted", "requests", s.ReadsCompleted)
	pm("ReadsMerged", "requests", s.ReadsMerged)
	pm("SectorsRead", "sectors", s.SectorsRead)
	pm("ReadTime", "ms", s.ReadTime)
	pm("WritesCompleted", "requests", s.WritesCompleted)
	pm("WritesMerged", "requests", s.WritesMerged)
	pm("SectorsWritten", "sectors", s.SectorsWritten)
	pm("WriteTime", "ms", s.WriteTime)
	pm("IOInProgress", "requests", s.IOInProgress)
	pm("IOTime", "ms", s.IOTime)
	pm("WeightedIOTime", "ms", s.WeightedIOTime)
	pm("DiscardsCompleted", "requests", s.DiscardsCompleted)
	pm("DiscardsMerged", "requests", s.DiscardsMerged)
	pm("SectorsDiscarded", "sectors", s.SectorsDiscarded)
	pm("DiscardTime", "ms", s.DiscardTime)
	pm("FlushesCompleted", "requests", s.FlushesCompleted)
	pm("FlushTime", "ms", s.FlushTime)
}

var printer = message.NewPrinter(language.English)

func pm(name, units string, value int64) {
	ps(name, units, printer.Sprintf("%d", value))
}

func ps(name, units, value string) {
	log.Printf("| %-17s | %-8s | %11s |\n", name, units, value)
}

func (s *Stats) Diff(s2 *Stats) *Stats {
	var diff = new(Stats)
	diff.ReadsCompleted = s2.ReadsCompleted - s.ReadsCompleted
	diff.ReadsMerged = s2.ReadsMerged - s.ReadsMerged
	diff.SectorsRead = s2.SectorsRead - s.SectorsRead
	diff.ReadTime = s2.ReadTime - s.ReadTime
	diff.WritesCompleted = s2.WritesCompleted - s.WritesCompleted
	diff.WritesMerged = s2.WritesMerged - s.WritesMerged
	diff.SectorsWritten = s2.SectorsWritten - s.SectorsWritten
	diff.WriteTime = s2.WriteTime - s.WriteTime
	diff.IOInProgress = s2.IOInProgress
	diff.IOTime = s2.IOTime - s.IOTime
	diff.WeightedIOTime = s2.WeightedIOTime - s.WeightedIOTime
	diff.DiscardsCompleted = s2.DiscardsCompleted - s.DiscardsCompleted
	diff.DiscardsMerged = s2.DiscardsMerged - s.DiscardsMerged
	diff.SectorsDiscarded = s2.SectorsDiscarded - s.SectorsDiscarded
	diff.DiscardTime = s2.DiscardTime - s.DiscardTime
	diff.FlushesCompleted = s2.FlushesCompleted - s.FlushesCompleted
	diff.FlushTime = s2.FlushTime - s.FlushTime
	return diff
}

func GetStats(connect, block string) *Stats {
	var commandString = fmt.Sprintf("%s cat /sys/block/%s/stat", connect, block)
	out, err := exec.Command("bash", "-c", commandString).Output()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Fields(string(out))
	var stats = new(Stats)
	stats.ReadsCompleted, _ = strconv.ParseInt(parts[0], 10, 64)
	stats.ReadsMerged, _ = strconv.ParseInt(parts[1], 10, 64)
	stats.SectorsRead, _ = strconv.ParseInt(parts[2], 10, 64)
	stats.ReadTime, _ = strconv.ParseInt(parts[3], 10, 64)
	stats.WritesCompleted, _ = strconv.ParseInt(parts[4], 10, 64)
	stats.WritesMerged, _ = strconv.ParseInt(parts[5], 10, 64)
	stats.SectorsWritten, _ = strconv.ParseInt(parts[6], 10, 64)
	stats.WriteTime, _ = strconv.ParseInt(parts[7], 10, 64)
	stats.IOInProgress, _ = strconv.ParseInt(parts[8], 10, 64)
	stats.IOTime, _ = strconv.ParseInt(parts[9], 10, 64)
	stats.WeightedIOTime, _ = strconv.ParseInt(parts[10], 10, 64)
	stats.DiscardsCompleted, _ = strconv.ParseInt(parts[11], 10, 64)
	stats.DiscardsMerged, _ = strconv.ParseInt(parts[12], 10, 64)
	stats.SectorsDiscarded, _ = strconv.ParseInt(parts[13], 10, 64)
	stats.DiscardTime, _ = strconv.ParseInt(parts[14], 10, 64)
	stats.FlushesCompleted, _ = strconv.ParseInt(parts[15], 10, 64)
	stats.FlushTime, _ = strconv.ParseInt(parts[16], 10, 64)
	return stats
}
