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
			background: #1c3e48;
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

		.flexSmall {
			flex-grow: 1;
		}

		.flexBig {
			flex-grow: 2;
		}

		.specialKeysContainer {
			margin: 0.5em;
			display: grid;
			grid-template-columns: auto auto auto;
			grid-gap: 0.5em;
			max-width: 800px;
			margin-left: auto;
			margin-right: auto;
		}

		.specialKey {
			margin: 0;
			padding: 0;
			padding-top: 100%;
			position: relative;
		}

		.specialKeyText {
			position: absolute;
			top: 0;
			left: 0;
			bottom: 0;
			right: 0;
			padding-top: 36%; /*hacky vertical center*/
		}

		.mouseButtonContainer {
			display: flex;
			max-width: 800px;
			margin-left: auto;
			margin-right: auto;
		}

		.mouseButton {
			background: #1e7eb0;
			height: 4em;
			margin: 0.2em;
			border-radius: 6px;
			box-shadow: 0 4px 12px #00000087;
			touch-action: none; /* To disable Zoom by double clicking in touch devices */
		}

		.mouseButton:hover {
			background: #3aaed8;
		}

		.mouseButton:active {
			box-shadow: 0 2px 4px #00000087;
		}

		#touchpadCanvas {
			background: #111;
			border-radius: 1em;
			width: 100%;
			max-width: 800px;
			box-shadow: inset 0 4px 2px 4px #00000087;
		}
		
	</style>
</head>

<body>
	<h2>Sound and music</h2>
	<div>
		<button onclick="Do('volumeUp')">🔊</button>
		<button onclick="Do('volumeDown')">🔉</button>
		<button onclick="Do('silence')">🔇</button>
	</div>
	<div>
		<button onclick="Do('prevSong')">⏮️</button>
		<button onclick="Do('pauseSong')">⏯️</button>
		<button onclick="Do('nextSong')">⏭️</button>
	</div>

	<h2>Special keys</h2>
	<div class="specialKeysContainer">
		<button class="specialKey" onclick="Do('keyEsc')">
			<div class="specialKeyText">Esc</div> 
		</button>
		<button class="specialKey" onclick="Do('keyUp')">
			<div class="specialKeyText">⬆️</div> 
		</button>
		<button class="specialKey" onclick="Do('keyEnter')">
			<div class="specialKeyText">Enter</div> 
		</button>

		<button class="specialKey" onclick="Do('keyLeft')">
			<div class="specialKeyText">⬅️</div> 
		</button>
		<button class="specialKey" onclick="Do('keyDown')">
			<div class="specialKeyText">⬇️</div> 
		</button>
		<button class="specialKey" onclick="Do('keyRight')">
			<div class="specialKeyText">➡️</div> 
		</button>
	</div>

	<h2>Mouse trackpad</h2>
	<div>
		<div class="mouseButtonContainer">
			<div class="mouseButton flexBig" onclick="Do('leftClick')"></div>
			<div class="mouseButton flexSmall" onclick="Do('middleClick')"></div>
			<div class="mouseButton flexBig" onclick="Do('rightClick')"></div>
		</div>
		<canvas id="touchpadCanvas" width="800" height="600">
			Your browser does not support the HTML canvas tag.
		</canvas>
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

		function websocketUri(relative) {
			var loc = window.location, uri
			if (loc.protocol === "https:") {
				uri = "wss:"
			} else {
				uri = "ws:"
			}
			uri += "//" + loc.host
			uri += loc.pathname + relative
			return uri
		}

		let mouseMoveSocket
		let lastPoint
		let movement

		var canvas = document.getElementById("touchpadCanvas")
		canvas.addEventListener("touchstart", handleStart, false)
		canvas.addEventListener("touchmove", handleMove, false)
		canvas.addEventListener("touchend", handleEnd, false)
		//canvas.addEventListener("touchcancel", handleCancel, false)

		function touchToPoint(touch) {
			return {
				x: touch.pageX - touch.target.offsetLeft,
				y: touch.pageY - touch.target.offsetTop,
			}
		}
		  
		function handleStart(evt) {
			evt.preventDefault()
			mouseMoveSocket = new WebSocket(websocketUri("mouseMove"))
			var touch = evt.changedTouches[0]
			movement = {x: 0, y: 0}
			lastPoint = touchToPoint(touch)
		}

		function handleEnd(evt) {
			evt.preventDefault()
			mouseMoveSocket.close()
			var touch = evt.changedTouches[0]
			var currPoint = touchToPoint(touch)

			if(currPoint.x == lastPoint.x && currPoint.y == lastPoint.y && !movement.x && !movement.y) {
				Do("leftClick")
			}
		}

		function handleMove(evt) {
			evt.preventDefault()
			if (!lastPoint) {
				return
			}

			var touch = evt.changedTouches[0]
			var currPoint = touchToPoint(touch)

			if (currPoint.x < 0 || currPoint.x > canvas.width) {
				if (currPoint.y < 0 || currPoint.y > canvas.height) {
					// we left the canvas
					handleLeave()
					return
				}
			}

			movement = {
				x: currPoint.x - lastPoint.x,
				y: currPoint.y - lastPoint.y,
			}

			// parseInt() because Chrome uses decimals
			mouseMoveSocket.send(parseInt(movement.x) + ":" + parseInt(movement.y))
			lastPoint = currPoint
		}

		function handleLeave() {
			window.alert(lastPoint.x + ":" + lastPoint.y)
			lastPoint = null
		}

	</script>
</body>
</html>
`
