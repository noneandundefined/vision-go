package vision

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/noneandundefined/vision-go/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Vision struct {
	stats *types.MonitoringStats
}

// The NewVision function is a constructor for creating
// a new instance of the Vision structure.
// It initializes the stats field with default values,
// including creating an empty LastErrors
// slice to store the latest errors.
func NewVision() *Vision {
	return &Vision{
		stats: &types.MonitoringStats{
			RequestCount:   0,
			ErrorCount:     0,
			TotalLatency:   0,
			DBQueryCount:   0,
			DBErrorCount:   0,
			DBTotalLatency: 0,
			LastErrors:     make([]types.ErrorLog, 0),
		},
	}
}

// VisionRequest request tracking
// ------------------------------
// The VisionRequest function takes the duration of the request
// and updates the statistics of requests to the server.
// It increases the counter of the total number
// of requests (RequestCount) and adds the transmitted
// duration to the total delay time of all requests (TotalLatency),
// ensuring that these values are securely updated using a
// mutex lock (Lock/Unlock) to avoid conflicts
// when accessing statistics from different execution threads.
func (v *Vision) VisionRequest(duration time.Duration) {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.RequestCount++
	v.stats.TotalLatency += duration
}

// VisionError error tracking
// --------------------------
// The VisionError function accepts an err error object and
// performs the following actions: locks the mutex
// for secure access to the fields of the stats structure,
// increments the error count counter, creates an error log
// containing the current time (Timestamp) and
// error message (Error), checks the length of the array
// of recent errors (LastErrors) and, if it contains
// more than 10 elements, it deletes the oldest element,
// and then adds a new error to the array.
func (v *Vision) VisionError(err error) {
	v.stats.Lock()
	defer v.stats.Unlock()
	v.stats.ErrorCount++

	err_log := types.ErrorLog{
		Timestamp: time.Now(),
		Error:     err.Error(),
	}

	// Limit the number of recent errors to 10
	if len(v.stats.LastErrors) >= 10 {
		v.stats.LastErrors = v.stats.LastErrors[1:]
	}

	v.stats.LastErrors = append(v.stats.LastErrors, err_log)
}

// VisionDBQuery tracking a database request
// -----------------------------------------
// The VisionDBQuery function takes the duration
// of a database query and updates the corresponding
// statistics by incrementing the database query
// counter (DBQueryCount) and adding the
// transmitted duration to the total amount
// of database query delays (DBTotalLatency)
func (v *Vision) VisionDBQuery(duration time.Duration) {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.DBQueryCount++
	v.stats.DBTotalLatency += duration
}

// VisionDBError tracking a DB error
// ---------------------------------
// The VisionDBError function increments
// the database error counter (DBErrorCount)
// when an error occurs.
func (v *Vision) VisionDBError() {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.DBErrorCount++
}

// The getCPUUsage function tries to get
// the percentage of CPU usage for the last second.
func getCPUUsage() (float64, error) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, fmt.Errorf("failed to get CPU usage: %w", err)
	}
	if len(cpuUsage) > 0 {
		return cpuUsage[0], nil
	}
	return 0, fmt.Errorf("no CPU usage data available")
}

// The getMemoryUsage function gets
// information about virtual memory usage.
func getMemoryUsage() (float64, error) {
	memoryUsage, err := mem.VirtualMemory()
	if err != nil {
		return 0, fmt.Errorf("failed to get memory usage: %w", err)
	}
	return memoryUsage.UsedPercent, nil
}

// The getNetworkStats function collects network
// interface statistics. If the network data cannot
// be retrieved, an error is returned. If the
// interface is found, the amount of data received i
// n megabytes is calculated and the result is returned.
func getNetworkStats() (float64, error) {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return 0, fmt.Errorf("failed to get network stats: %w", err)
	}

	if len(netStats) > 0 {
		return float64(netStats[0].BytesRecv) / 1024 / 1024, nil
	}
	return 0, fmt.Errorf("no network interface found")
}

// GetVisionStats getting statistics
// ---------------------------------
// The GetVisionStats function collects
// all the data into one structure
// and outputs it as a JSON structure.
func (v *Vision) GetVisionStats() types.MonitoringResponse {
	v.stats.Lock()
	defer v.stats.Unlock()

	// Get system metrics
	cpuUsage, err := getCPUUsage()
	if err != nil {
		log.Printf("Error getting CPU usage: %v", err)
		cpuUsage = 0 // Set a default value or handle the error as needed
	}

	memoryUsage, err := getMemoryUsage()
	if err != nil {
		log.Printf("Error getting memory usage: %v", err)
		memoryUsage = 0 // Set a default value or handle the error as needed
	}

	networkRecv, err := getNetworkStats()
	if err != nil {
		networkRecv = 0 // Set default
	}

	return types.MonitoringResponse{
		Requests: struct {
			Total        int64   `json:"request_count"`
			Errors       int64   `json:"request_error_count"`
			SuccessRate  float64 `json:"request_success_count"`
			AvgLatencyMs float64 `json:"request_avg_latency_ms"`
		}{
			Total:  v.stats.RequestCount,
			Errors: v.stats.ErrorCount,
			SuccessRate: func(stats *types.MonitoringStats) float64 {
				if stats.RequestCount == 0 {
					return 0
				}
				successCount := stats.RequestCount - stats.ErrorCount
				successRate := float64(successCount) / float64(stats.RequestCount) * 100
				return math.Round(successRate*100) / 100
			}(v.stats),
			AvgLatencyMs: func(stats *types.MonitoringStats) float64 {
				if stats.RequestCount == 0 {
					return 0
				}
				avgLatency := float64(stats.TotalLatency.Milliseconds()) / float64(stats.RequestCount)
				return math.Round(avgLatency*100) / 100
			}(v.stats),
		},
		Database: struct {
			TotalQueries int64   `json:"database_queries"`
			Errors       int64   `json:"database_error"`
			AvgLatencyMs float64 `json:"database_avg_latency_ms"`
		}{
			TotalQueries: v.stats.DBQueryCount,
			Errors:       v.stats.DBErrorCount,
			AvgLatencyMs: func(stats *types.MonitoringStats) float64 {
				if stats.DBQueryCount == 0 {
					return 0
				}
				return float64(v.stats.DBTotalLatency.Milliseconds()) / float64(v.stats.DBQueryCount)
			}(v.stats),
		},
		System: struct {
			CPUUsage    float64 `json:"cpu_usage"`
			MemoryUsage float64 `json:"memory_usage"`
			NetworkRecv float64 `json:"network_recv"`
		}{
			CPUUsage:    cpuUsage,
			MemoryUsage: memoryUsage,
			NetworkRecv: networkRecv,
		},
		LastErrors: v.stats.LastErrors,
	}
}
