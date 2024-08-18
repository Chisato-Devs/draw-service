package v1

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"net/http"
	"runtime"
)

func memoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func cpuPercent() (float64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) != 0 {
		context.CreateErrorResponse(w, "There shouldn't be anything in the link parameters!", http.StatusBadRequest)
		return
	}

	cpuUsage, _ := cpuPercent()
	stats := structs.Stats{
		MemoryUsage: memoryUsage(),
		NumCPU:      runtime.NumCPU(),
		Goroutines:  runtime.NumGoroutine(),
		CPUUsage:    cpuUsage,
	}

	jsonBytes, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		return
	}
}
