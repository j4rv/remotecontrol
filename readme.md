## PC Remote Control

Extremely simple implementation or a "Remote controller" for my convenience,
but Open Source in case it's useful for someone else.

To install it (Requires Go):

``go get https://github.com/j4rv/remotecontrol/bin/remotecontrol``

To execute it: ``remotecontrol`` (duh). An HTTP server will start running (uses port 80).

You will get the list of your PC's IPs in the console.
Then, as long as you are in the same network,
you should be able to ""control"" your PC from a mobile device,
or a laptop, or a fridge if it has a browser.

![Example screenshot](https://github.com/j4rv/remotecontrol/images/screenshot.png)