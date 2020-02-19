package main

const indexTmpl = `
<!doctype html>

<html lang="en">
<head>
	<meta charset="utf-8">
	<title>PC Control</title>
	<style>
		* {
			text-align: center;
		}

		body {
			background: #112024;
		}

		button {
			font-size: 2em;
			margin: 0.5em;
			padding: 0.5em;
			border-radius: 6px;
			background: #1e7eb0;
			color: white;
			border: none;
			box-shadow: 0 4px 12px #00000087;
			touch-action: none; /* To disable Zoom by double clicking in touch devices */
		}

		button:hover {
			background: #3aaed8;
		}

		button:active {
			box-shadow: 0 2px 4px #00000087;
		}

		@media only screen
		and (-webkit-min-device-pixel-ratio: 2) {
			body {
				font-size: large;
			}
		}

		@media only screen
		and (-webkit-min-device-pixel-ratio: 3) {
			body {
				font-size: x-large;
			}
		}

		@media only screen
		and (-webkit-min-device-pixel-ratio: 4) {
			body {
				font-size: xx-large;
			}
		}
	</style>
</head>

<body>
	<div>
		<button onclick="Do('volumeUp')">Volume up</button>
		<button onclick="Do('volumeDown')">Volume down</button>
		<button onclick="Do('silence')">Silence</button>
	</div>
	<div>
		<button onclick="Do('shutdown1m')">Shutdown in 1 minute</button>
		<button onclick="Do('shutdown30m')">Shutdown in 30 minutes</button>
		<button onclick="Do('shutdown60m')">Shutdown in 60 minutes</button>
		<button onclick="Do('shutdown120m')">Shutdown in 120 minutes</button>
		<button onclick="Do('abortShutdown')">Abort shutdown</button>
	</div>
	<script>
		function Do(action) {
			fetch(action)
		}
	</script>
</body>
</html>
`
