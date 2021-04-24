package CPU

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"stats/Error"
)

type CPU struct {
	CoreUsage []float64          `json:"core_usage,omitempty"`
	CoreTemp  map[string]float64 `json:"core_temp,omitempty"`
	SensorsID []string           `json:"sensors_id,omitempty"`
	Clock     []float64          `json:"clock,omitempty"`
}

func GetCPUStats(cpuStats chan<- CPU) {

	var (
		buffer CPU
		err    error
	)

	buffer.CoreUsage, err = cpu.Percent(1, true)
	Error.PSCPUError(err)

	buffer.CoreTemp = make(map[string]float64)
	sensors, err := host.SensorsTemperatures()
	Error.PSCPUError(err)

	for _, sensor := range sensors {
		if sensor.Temperature != 0 {
			buffer.CoreTemp[sensor.SensorKey] = sensor.Temperature
			buffer.SensorsID = append(buffer.SensorsID, sensor.SensorKey)
		}
	}

	informations, err := cpu.Info()
	Error.PSCPUError(err)
	for _, info := range informations {
		buffer.Clock = append(buffer.Clock, info.Mhz)
	}

	cpuStats <- buffer
}
