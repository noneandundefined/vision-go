package vemail

import (
	"fmt"
	"time"

	"github.com/noneandundefined/vision-go"
)

func LoadEmailTemplate(stats *vision.Vision) string {
	visionStats := stats.GetVisionStats()

	/// Notifications
	notifications := ""
	/// CPUs
	if visionStats.System.CPUUsage >= 45 && visionStats.System.CPUUsage < 75 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #e8b000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [CPU]</p>
			<p style="color: #fff; padding-left: 20px">
				⚠️ Warning! CPU exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.System.CPUUsage)
	} else if visionStats.System.CPUUsage >= 75 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #cf0000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [CPU]</p>
			<p style="color: #fff; padding-left: 20px">
				🚨 Instability! CPU exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.System.CPUUsage)
	}

	/// Memory
	if visionStats.System.MemoryUsage >= 55 && visionStats.System.MemoryUsage < 75 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #e8b000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Memory]</p>
			<p style="color: #fff; padding-left: 20px">
				⚠️ Warning! Memory exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.System.MemoryUsage)
	} else if visionStats.System.MemoryUsage >= 75 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #cf0000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Memory]</p>
			<p style="color: #fff; padding-left: 20px">
				🚨 Instability! Memory exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.System.MemoryUsage)
	}

	/// Response server time
	if visionStats.Requests.AvgLatencyMs >= 10000 && visionStats.Requests.AvgLatencyMs < 13000 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #e8b000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Server latency (ms)]</p>
			<p style="color: #fff; padding-left: 20px">
				⚠️ Warning! Server latency exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.Requests.AvgLatencyMs)
	} else if visionStats.Requests.AvgLatencyMs >= 13000 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #cf0000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Server latency (ms)]</p>
			<p style="color: #fff; padding-left: 20px">
				🚨 Instability! Server latency exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.Requests.AvgLatencyMs)
	}

	/// Databases
	if visionStats.Database.AvgLatencyMs >= 10000 && visionStats.Database.AvgLatencyMs < 13000 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #e8b000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Database latency (ms)]</p>
			<p style="color: #fff; padding-left: 20px">
				⚠️ Warning! Database latency exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.Database.AvgLatencyMs)
	} else if visionStats.Database.AvgLatencyMs >= 13000 {
		notifications += fmt.Sprintf(`
		<div
			style="
				margin: 10px 0;
				background-color: #cf0000;
				display: flex;
				padding: 12px;
				border-radius: 6px;
			"
		>
			<p style="color: #fff">[%s] [Database latency (ms)]</p>
			<p style="color: #fff; padding-left: 20px">
				🚨 Instability! Database latency exceeds %.2f%%
			</p>
		</div>
		`, time.Now().Format("02.01.2006 15:04:05"), visionStats.Database.AvgLatencyMs)
	}

	/// HTML email
	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="preconnect" href="https://fonts.googleapis.com" />
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
		<link
			href="https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap"
			rel="stylesheet"
		/>
		<style>
			body,
			* {
				font-family: 'Inter', sans-serif;
			}

			p {
				margin: 0;
				padding: 0;
			}
		</style>
	</head>
	<body>
		<section style="padding: 20px; border-radius: 3px">
			<div
				style="display: flex; align-items: center; margin-bottom: 40px"
			>
				<img
					width="50"
					height="50"
					alt=""
					src="https://github.com/Artymiik/vision/blob/main/public/logo-vision-none.png?raw=true"
				/>
				<div style="margin-left: 1rem">
					<p>
						<a
							style="text-decoration: underline"
							href="https://github.com/Artymiik/vision"
							>Data Sources</a
						>
						/ Vision
					</p>
					<p style="font-size: 13px; color: #666; margin-left: -14px">
						Server monitoring
					</p>
				</div>
			</div>

			<h4 style="text-align: left">System statistic</h4>
			<div
				style="
					display: flex;
					align-items: center;
					text-align: center;
					justify-content: space-around;
				"
			>
				<div style="text-align: center; padding: 0 40px">
					<p>CPU</p>
					<p>%.2f%% / 100%%</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Memory</p>
					<p>%.2f%% / 100%%</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Network</p>
					<p>%.2f MB / 100MB</p>
				</div>
			</div>

			<h4 style="text-align: left">Server statistic</h4>
			<div
				style="
					display: flex;
					align-items: center;
					text-align: center;
					justify-content: space-around;
				"
			>
				<div style="text-align: center; padding: 0 40px">
					<p>Total request</p>
					<p>%d</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Response time</p>
					<p>%.2f</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Total errors</p>
					<p>%d</p>
				</div>
			</div>

			<h4 style="text-align: left">Database statistic</h4>
			<div
				style="
					display: flex;
					align-items: center;
					text-align: center;
					justify-content: space-around;
				"
			>
				<div style="text-align: center; padding: 0 40px">
					<p>Total queries</p>
					<p>%d</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Total errors</p>
					<p>%d</p>
				</div>

				<div style="text-align: center; padding: 0 40px">
					<p>Response time</p>
					<p>%.2f</p>
				</div>
			</div>
			<h4 style="text-align: left">Notifications</h4>
			<div
				style="
					align-items: center;
					text-align: left;
					justify-content: space-around;
				"
			>
				%s
			</div>
		</section>
	</body>
</html>
    `, visionStats.System.CPUUsage, visionStats.System.MemoryUsage, (visionStats.System.NetworkRecv / 100),
		visionStats.Requests.Total, visionStats.Requests.AvgLatencyMs, visionStats.Requests.Errors,
		visionStats.Database.TotalQueries, visionStats.Database.Errors, visionStats.Database.AvgLatencyMs,
		notifications)

	return htmlContent
}
