// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package types

import (
	"sync"
	"time"
)

// A model for monitoring the application
type MonitoringStats struct {
	sync.Mutex
	RequestCount   int64
	ErrorCount     int64
	TotalLatency   time.Duration
	DBQueryCount   int64
	DBErrorCount   int64
	DBTotalLatency time.Duration
	LastErrors     []ErrorLog
}

// The Server error model
type ErrorLog struct {
	Timestamp time.Time `json:"timestamp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Error     string    `json:"error"`
}

// A model for the vision monitoring response
type MonitoringResponse struct {
	Requests struct {
		Total        int64   `json:"request_count"`
		Errors       int64   `json:"request_error_count"`
		SuccessRate  float64 `json:"request_success_count"`
		AvgLatencyMs float64 `json:"request_avg_latency_ms"`
	} `json:"requests"`
	Database struct {
		TotalQueries int64   `json:"database_queries"`
		Errors       int64   `json:"database_error"`
		AvgLatencyMs float64 `json:"database_avg_latency_ms"`
	} `json:"database"`
	System struct {
		CPUUsage    float64 `json:"cpu_usage"`
		MemoryUsage float64 `json:"memory_usage"`
		NetworkRecv float64 `json:"network_recv"`
	} `json:"system"`
	LastErrors []ErrorLog `json:"last_errors,omitempty"`
}
