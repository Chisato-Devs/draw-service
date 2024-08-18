package structs

type Stats struct {
	MemoryUsage uint64  `json:"memory_usage"`
	NumCPU      int     `json:"num_cpu"`
	Goroutines  int     `json:"goroutines"`
	CPUUsage    float64 `json:"cpu_usage"`
}
