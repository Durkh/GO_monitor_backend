package Memory

import (
	"github.com/shirou/gopsutil/v3/mem"
	"stats/Error"
)

type Memory struct {
	TotalMemory uint64  `json:"total_memory,omitempty"`
	Used        uint64  `json:"used,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty"`
}

func GetMemStats(memStats chan<- Memory) {

	virtMem, err := mem.VirtualMemory()
	Error.PSMemError(err)

	buffer := Memory{
		TotalMemory: virtMem.Total,
		Used:        virtMem.Used,
		UsedPercent: virtMem.UsedPercent,
	}

	memStats <- buffer
}
