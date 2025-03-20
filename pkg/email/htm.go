package email

import (
	"fmt"

	"github.com/noneandundefined/vision-go"
)

func LoadEmailTemplate(stats *vision.Vision) string {
	// Получаем статистику один раз для избежания повторных вызовов
	visionStats := stats.GetVisionStats()

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="preconnect" href="https://fonts.googleapis.com" />
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
		<link
			href="https://fonts.googleapis.com/css2?family=DM+Mono:ital,wght@0,300;0,400;0,500;1,300;1,400;1,500&display=swap"
			rel="stylesheet"
		/>
		<style>
			body,
			* {
				font-family: 'DM Mono', monospace;
			}
		</style>
	</head>
	<body>
		<section
			style="
				min-width: 500px;
				margin: 0 auto;
				align-items: center;
				background-color: #eee;
				padding: 10;
				border-radius: 3px;
			"
		>
			<h4>System statistic</h4>
			<div
				style="
					display: flex;
					align-items: center;
					justify-content: space-around;
				"
			>
				<div style="text-align: center">
					<p>CPU</p>
					<p>%.2f%%%% / 100%%%%</p>
				</div>

				<div style="text-align: center">
					<p>Memmory</p>
					<p>%.2f%%%% / 100%%%%</p>
				</div>

				<div style="text-align: center">
					<p>Network</p>
					<p>%.2f MB / 100MB</p>
				</div>
			</div>

			<h4>Server statistic</h4>
			<div
				style="
					display: flex;
					align-items: center;
					justify-content: space-around;
				"
			>
				<div style="text-align: center">
					<p>Request count</p>
					<p>%d</p>
				</div>

				<div style="text-align: center">
					<p>Response time</p>
					<p>%.2f</p>
				</div>

				<div style="text-align: center">
					<p>Error count</p>
					<p>%d</p>
				</div>
			</div>

			<h4>Database statistic</h4>
			<h4>Notifications</h4>
		</section>
	</body>
</html>
    `, visionStats.System.CPUUsage, visionStats.System.MemoryUsage, visionStats.System.NetworkRecv,
		visionStats.Requests.Total, visionStats.Requests.AvgLatencyMs, visionStats.Requests.Errors)

	return htmlContent
}
