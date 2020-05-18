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
			font-family: arial,sans-serif;
		}

		body {
			background: #112024;
			font-size: 2em;
		}

		h2 {
			color: white;
			margin-bottom: 0;
		}

		button {
			font-size: 1.5em;
			margin: 0.5em;
			padding: 0.5em;
			border-radius: 6px;
			min-width: 3em;
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
		
	</style>
</head>

<body>
	<h2>Volume</h2>
	<div>
		<button onclick="Do('volumeUp')">🔊</button>
		<button onclick="Do('volumeDown')">🔉</button>
		<button onclick="Do('silence')">🔇</button>
	</div>

	<h2>Tracks</h2>
	<div>
		<button onclick="Do('prevSong')">⏮️</button>
		<button onclick="Do('pauseSong')">⏯️</button>
		<button onclick="Do('nextSong')">⏭️</button>
	</div>

	<h2>Delayed shutdowns</h2>
	<div>
		<button onclick="Do('shutdown1m')">Shutdown 1m</button>
		<button onclick="Do('shutdown30m')">Shutdown 30m</button>
		<button onclick="Do('shutdown60m')">Shutdown 1h</button>
		<button onclick="Do('shutdown120m')">Shutdown 2h</button>
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
