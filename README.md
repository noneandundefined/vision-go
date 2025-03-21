<p align="center">
  <a href="" target="_blank" rel="noopener noreferrer">
    <img width="150" src="https://github.com/noneandundefined/vision-ui/blob/main/public/logo-vision-none.png" alt="Vision logo">
  </a>
  <a href="" target="_blank" rel="noopener noreferrer">
    <img width="65" src="https://github.com/noneandundefined/vision-go/blob/main/public/go.png" alt="JS logo">
  </a>
</p>
</br>
<p align="center">
  <a href="https://github.com/Artymiik/vision/actions/workflows/compiler-client.yml"><img src="https://github.com/Artymiik/vision/actions/workflows/compiler-client.yml/badge.svg" alt="Compiler and Build"></a>
</p>
<p align="center">
  <a href="https://www.npmjs.com/package/@artemiik/vision-ui"><img src="https://img.shields.io/npm/v/%40artemiik%2Fvision-ui" alt="Compiler and Build"></a>
  <a href="https://www.npmjs.com/package/@artemiik/vision-ui"><img src="https://img.shields.io/npm/dm/%40artemiik%2Fvision-ui" alt="Compiler and Build"></a>
</p>

## Introduction

Vision allows any user, whether it's your development team, to visually look at the server in production mode. Be aware of all errors, server response time, CPU and RAM load. Look at the Vision!

## Versions

| Vision Version | Release Date | Notes                                                           |
| -------------- | ------------ | --------------------------------------------------------------- |
| 1.0.0          | 2025-02-16   | [tag v1.0.0](https://github.com/Artymiik/vision-ui/tree/v1.0.0) |

## Documentation

## Go

Configuring GO for server monitoring

### Content

-   Initial configuration
-   Models
-   Methods of obtaining monitoring
-   Statistics
-   Linking monitoring to an endpoint

## Setup

> Below are the basic functions to get the basic monitoring components for your server, select the necessary ones and use them

### Initial configuration

```sh
go get github.com/noneandundefined/vision-go
```

```go
package main

import "github.com/noneandundefined/vision-go"

type Handler struct {
	monitor *vision.Vision
}

// The NewHandler function is a constructor for creating
// a new instance of the Vision structure.
// It initializes the stats field with default values,
// including creating an empty LastErrors
// slice to store the latest errors.
func NewHandler(monitor *vision.Vision) *Handler {
	return &Handler{
		monitor: monitor,
	}
}
```

### Models

1. A model for monitoring the application:

```go
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
```

2. The Server error model

```go
type ErrorLog struct {
	Timestamp time.Time `json:"timestamp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Error     string    `json:"error"`
}
```

3. A model for the vision monitoring response

```go
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
```

### Methods of obtaining monitoring

The VisionRequest function takes the duration of the request and updates the statistics of requests to the server. It increases the counter of the total number of requests (RequestCount) and adds the transmitted duration to the total delay time of all requests (TotalLatency), ensuring that these values are securely updated using a mutex lock (Lock/Unlock) to avoid conflicts when accessing statistics from different execution threads.

```go
func (v *Vision) VisionRequest(duration time.Duration) {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.RequestCount++
	v.stats.TotalLatency += duration
}
```

The VisionError function accepts an err error object and performs the following actions: locks the mutex for secure access to the fields of the stats structure, increments the error count counter, creates an error log containing the current time (Timestamp) and error message (Error), checks the length of the array of recent errors (LastErrors) and, if it contains more than 10 elements, it deletes the oldest element, and then adds a new error to the array.

```go
func (v *Vision) VisionError(err error) {
	v.stats.Lock()
	defer v.stats.Unlock()
	v.stats.ErrorCount++

	err_log := ErrorLog{
		Timestamp: time.Now(),
		Error:     err.Error(),
	}

	// Limit the number of recent errors to 10
	if len(v.stats.LastErrors) >= 10 {
		v.stats.LastErrors = v.stats.LastErrors[1:]
	}

	v.stats.LastErrors = append(v.stats.LastErrors, err_log)
}
```

The VisionDBQuery function takes the duration of a database query and updates the corresponding statistics by incrementing the database query counter (DBQueryCount) and adding the transmitted duration to the total amount of database query delays (DBTotalLatency)

```go
func (v *Vision) VisionDBQuery(duration time.Duration) {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.DBQueryCount++
	v.stats.DBTotalLatency += duration
}
```

The VisionDBError function increments the database error counter (DBErrorCount) when an error occurs.

```go
func (v *Vision) VisionDBError() {
	v.stats.Lock()
	defer v.stats.Unlock()

	v.stats.DBErrorCount++
}

```

The getCPUUsage function tries to get the percentage of CPU usage for the last second.

```go
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
```

The getMemoryUsage function gets information about virtual memory usage.

```go
func getMemoryUsage() (float64, error) {
	memoryUsage, err := mem.VirtualMemory()
	if err != nil {
		return 0, fmt.Errorf("failed to get memory usage: %w", err)
	}
	return memoryUsage.UsedPercent, nil
}
```

The getNetworkStats function collects network interface statistics. If the network data cannot be retrieved, an error is returned. If the interface is found, the amount of data received in megabytes is calculated and the result is returned.

```go
func getNetworkStats() (float64, error) {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return 0, fmt.Errorf("failed to get network stats: %w", err)
	}

	if len(netStats) > 0 {
		bytesPerSec := float64(netStats[0].BytesRecv) / time.Since(netStats[0].TimeStamp).Seconds()
		usagePercentage := bytesPerSec / float64(maxBandwidth) * 100
		return usagePercentage, nil
	}
	return 0, fmt.Errorf("no network interface found")
}
```

### Statistics

The GetVisionStats function collects all the data into one structure and outputs it as a JSON structure.

```go
func (v *Vision) GetVisionStats() MonitoringResponse {
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

	return MonitoringResponse{
		Requests: struct {
			Total        int64   `json:"request_count"`
			Errors       int64   `json:"request_error_count"`
			SuccessRate  float64 `json:"request_success_count"`
			AvgLatencyMs float64 `json:"request_avg_latency_ms"`
		}{
			Total:  v.stats.RequestCount,
			Errors: v.stats.ErrorCount,
			SuccessRate: func(stats *MonitoringStats) float64 {
				if stats.RequestCount == 0 {
					return 0
				}
				successCount := stats.RequestCount - stats.ErrorCount
				successRate := float64(successCount) / float64(stats.RequestCount) * 100
				return math.Round(successRate*100) / 100
			}(v.stats),
			AvgLatencyMs: func(stats *MonitoringStats) float64 {
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
			AvgLatencyMs: float64(v.stats.DBTotalLatency.Milliseconds()) / float64(v.stats.DBQueryCount),
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
```

```go
package api

import "github.com/noneandundefined/vision-go"

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))))

	// Vision monitoring
	monitoring := vision.NewVision()
	router.HandleFunc("/admin/vision", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/vision/index.html")
	}).Methods("GET")
}
```

This one index.html it may be an outdated version, see the current version [index.html -> dev-helpers](https://github.com/noneandundefined/vision-ui/blob/main/dev-helpers/index.html)

```html
<!-- HTML for dev server -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<link
			rel="icon"
			type="image/svg+xml"
			href="https://github.com/noneandundefined/vision/blob/main/public/logo-vision-none.png?raw=true"
		/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta
			name="description"
			content="Keep up to date with what is happening on your server using - Vision"
		/>

		<!-- Stylization and actions -->
		<script
			type="module"
			src="https://unpkg.com/@artemiik/vision-ui@1.0.2/dist/vision.bundle.js"
			defer
		></script>
		<link
			rel="stylesheet"
			href="https://unpkg.com/@artemiik/vision-ui@1.0.2/dist/vision.css"
		/>

		<title>Look at the Vision!</title>

		<!-- Defining a meta tag for the monitoring URL -->
		<!-- Thanks to this meta tag, Vision will take server monitoring -->
		<meta
			name="monitoring-url"
			content="http://localhost:8001/micro/user/admin/vision/stats"
		/>
	</head>
	<body>
		<!--
			This HTML file is a template.

			Vision allows any user, whether it's your development team,
			to visually look at the server in production mode.

			Be aware of all errors, server response time, CPU and RAM load.
			Look at the Vision!
		-->
		<div id="vision"></div>
	</body>
</html>
```

Additional documentation on using Vision UI -> [Look at the Vision!](https://github.com/noneandundefined/vision-ui/blob/main/docs/languages/go.md)

## Использование Selectel Object Storage для хранения конфиденциальных данных

Вы можете использовать Selectel Object Storage для безопасного хранения и получения конфиденциальных данных, включая файлы `.env`.

### Подготовка

1. Создайте аккаунт в [Selectel](https://selectel.ru/services/cloud/storage/)
2. Создайте контейнер (бакет) для хранения данных
3. Получите ключи доступа (Access Key и Secret Key)
4. Установите необходимые зависимости:

```bash
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/credentials
go get github.com/aws/aws-sdk-go-v2/service/s3
```

### Загрузка .env файла в Selectel

```go
package main

import (
	"log"

	"github.com/noneandundefined/vision-go/pkg/cloud"
)

func main() {
	// Создаем клиент Selectel
	client, err := cloud.NewSelectelClient(cloud.SelectelConfig{
		AccessKey: "ваш_access_key",
		SecretKey: "ваш_secret_key",
		Endpoint:  "https://s3.selcdn.ru", // Эндпоинт Selectel
		Bucket:    "имя_вашего_бакета",
	})

	if err != nil {
		log.Fatalf("Ошибка создания клиента: %v", err)
	}

	// Загружаем .env файл в Selectel
	err = cloud.UploadEnvFile(client, "имя_вашего_бакета", "env/config.env", ".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки: %v", err)
	}

	log.Println("Файл .env успешно загружен в Selectel Object Storage")
}
```

### Загрузка .env файла из Selectel

```go
package main

import (
	"log"

	"github.com/noneandundefined/vision-go/pkg/cloud"
)

func main() {
	// Загружаем .env файл из Selectel и применяем переменные окружения
	err := cloud.LoadEnvFromSelectel(
		"ваш_access_key",
		"ваш_secret_key",
		"https://s3.selcdn.ru",
		"имя_вашего_бакета",
		"env/config.env",
	)

	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	log.Println("Конфигурация успешно загружена из Selectel Object Storage")

	// Теперь можно использовать переменные окружения
	// ...
}
```

### Безопасность

1. Храните ключи доступа к Selectel в безопасном месте
2. Используйте HTTPS для соединения с Selectel
3. Настройте права доступа к бакету, чтобы ограничить доступ
4. Рассмотрите возможность шифрования файлов перед загрузкой для дополнительной безопасности
