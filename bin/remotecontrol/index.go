package main

const indexTmpl = `
<!doctype html>

<html lang="en">
<head>
	<meta charset="utf-8">
	<title>PC Control</title>
	<script src="https://kit.fontawesome.com/cb96d94f42.js" crossorigin="anonymous"></script>
	<style>
		* {
			text-align: center;
			font-family: arial,sans-serif;
		}

		body {
			background: #40454d;
			background: linear-gradient(90deg, #3d4e51 0%, #40454d 5%, #40454d 95%, #2b2a2f 100%);
			font-size: 0;
		}

		h2 {
			font-size: 2rem;
			color: white;
			margin-bottom: 0;
		}

		button {
			font-size: 3rem;
			margin: 0.6rem;
			padding: 1rem;
			border-radius: 32px;
			min-width: 10rem;
			min-height: 6rem;
			background: #6d7d84;
			color: white;
			border: none;
			box-shadow: 0 4px 12px #00000087, inset 2px 4px 0px 2px #6d9799, inset -2px -4px 0px 2px #3d4e51;
			touch-action: none; /* To disable Zoom by double clicking in touch devices */
		}

		button:active {
			box-shadow: 0 2px 4px #00000087, inset 2px 4px 0px 2px #3d4e51, inset -2px -4px 0px 2px #6d9799;
		}	

		button:hover {
			background: #576f71;
		}

		.flexSmall {
			flex-grow: 1;
		}

		.flexBig {
			flex-grow: 3;
		}

		.specialKey {
			min-height: 8rem;
		}

		.mouseButtonContainer {
			display: flex;
			max-width: 80%;
			margin-left: auto;
			margin-right: auto;
		}

		.mouseButton {
			height: 6rem;
			margin: 1rem 0 0 0;
			border-radius: 6px;
		}

		#touchpadCanvas {
			background: #111;
			border-radius: 6px 6px 32px 32px;
			width: 100%;
			max-width: 80%;
			box-shadow: inset 0 4px 2px 4px #00000087;
		}

	</style>
</head>

<body>
	<h2>Sound and music</h2>
	<div>
		<button onclick="Do('volumeUp')"><i class="fas fa-volume-up"></i></button>
		<button onclick="Do('volumeDown')"><i class="fas fa-volume-down"></i></button>
		<button onclick="Do('silence')"><i class="fas fa-volume-mute"></i></button>
	</div>
	<div>
		<button onclick="Do('prevSong')"><i class="fas fa-sm fa-fast-backward"></i></button>
		<button onclick="Do('pauseSong')"><i class="fas fa-xs fa-play"></i></button>
		<button onclick="Do('nextSong')"><i class="fas fa-sm fa-fast-forward"></i></button>
	</div>

	<h2>Special keys</h2>
	<div>
		<button class="specialKey" onclick="Do('keyEsc')">
			Esc
		</button>
		<button class="specialKey" onclick="Do('keyUp')">
			<i class="fas fa-arrow-up"></i>
		</button>
		<button class="specialKey" onclick="Do('keyEnter')">
			Enter
		</button>
	</div>

	<div>
		<button class="specialKey" onclick="Do('keyLeft')">
			<i class="fas fa-arrow-left"></i>
		</button>
		<button class="specialKey" onclick="Do('keyDown')">
			<i class="fas fa-arrow-down"></i>
		</button>
		<button class="specialKey" onclick="Do('keyRight')">
			<i class="fas fa-arrow-right"></i>
		</button>
	</div>

	<h2>Mouse trackpad</h2>
	<div>
		<div class="mouseButtonContainer">
			<button class="mouseButton flexBig" onclick="Do('leftClick')"></button>
			<button class="mouseButton flexSmall" onclick="Do('middleClick')"></button>
			<button class="mouseButton flexBig" onclick="Do('rightClick')"></button>
		</div>
		<canvas id="touchpadCanvas" width="800" height="600">
			Your browser does not support the HTML canvas tag.
		</canvas>
	</div>

	<h2>Delayed shutdowns</h2>
	<div>
		<button onclick="Do('shutdown?mins=1')">1m</button>
		<button onclick="Do('shutdown?mins=15')">15m</button>
		<button onclick="Do('shutdown?mins=30')">30m</button>
	</div>
	<div>
		<button onclick="Do('shutdown?mins=60')">1h</button>
		<button onclick="Do('shutdown?mins=120')">2h</button>
		<button onclick="Do('shutdown?mins=240')">4h</button>
	</div>
	<div>
		<button onclick="Do('abortShutdown')">Abort</button>
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

		const mouseMoveSocket = new WebSocket(websocketUri("mouseMoveWebSocket"))
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
			var touch = evt.changedTouches[0]
			movement = {x: 0, y: 0}
			lastPoint = touchToPoint(touch)
		}

		function handleEnd(evt) {
			evt.preventDefault()
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
