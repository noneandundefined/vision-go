package vemail

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
				color: #fff;
			}
		</style>
	</head>
	<body>
		<section
			style="
				max-width: 600px;
				margin: 0 auto;
				text-align: center;
				background-color: #000024;
				padding: 40px;
				border-radius: 3px;
			"
		>
			<div style="display: flex; align-items: center">
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
					<p style="font-size: 13px; color: #eee; margin-left: -14px;">Server monitoring</p>
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
				<div style="text-align: center; padding: 0 20px">
					<p>CPU</p>
					<p>%.2f%%%% / 100%%%%</p>
				</div>

				<div style="text-align: center; padding: 0 20px">
					<p>Memmory</p>
					<p>%.2f%%%% / 100%%%%</p>
				</div>

				<div style="text-align: center; padding: 0 20px">
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
				<div style="text-align: center; padding: 0 20px">
					<p>Request count</p>
					<p>%d</p>
				</div>

				<div style="text-align: center; padding: 0 20px">
					<p>Response time</p>
					<p>%.2f</p>
				</div>

				<div style="text-align: center; padding: 0 20px">
					<p>Error count</p>
					<p>%d</p>
				</div>
			</div>

			<h4 style="text-align: left">Database statistic</h4>
			<h4 style="text-align: left">Notifications</h4>
		</section>
	</body>
</html>
    `, visionStats.System.CPUUsage, visionStats.System.MemoryUsage, visionStats.System.NetworkRecv,
		visionStats.Requests.Total, visionStats.Requests.AvgLatencyMs, visionStats.Requests.Errors)

	return htmlContent
}
