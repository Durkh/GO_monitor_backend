package main

import (
	"stats/CPU"
	"stats/Memory"
)

type Sendable struct {
	CpuStats CPU.CPU       `json:"cpu_stats"`
	MemStat  Memory.Memory `json:"mem_stat"`
}

func main() {

	cpuInfo	:= make(chan CPU.CPU)
	memInfo	:= make(chan Memory.Memory)

	go CPU.GetCPUStats(cpuInfo)
	go Memory.GetMemStats(memInfo)

	sendablePackage := Sendable{
		CpuStats: <-cpuInfo,
		MemStat:  <-memInfo,
	}

	_ = sendablePackage

}
