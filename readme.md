## PC Remote Control

Extremely simple implementation or a "Remote controller" for my convenience,
but Open Source in case it's useful for someone else.

To install it (Requires Go):

``go get github.com/j4rv/remotecontrol/bin/remotecontrol``

To execute it: ``remotecontrol`` (duh). An HTTP server will start running (uses port 80).

You will get the list of your PC's IPs in the console.
Then, as long as you are in the same network,
you should be able to ""control"" your PC from a mobile device,
or a laptop, or a fridge if it has a browser.

#### Warning

Obvious warning, but **only use this if your PC is in a trusted, private network**.

![Example screenshot](https://raw.githubusercontent.com/j4rv/remotecontrol/master/images/screenshot.jpg)

#### TODOs

- [x] Volume control
- [x] Timed shutdowns
- [ ] Prettier interface
- [ ] Improve security
- [x] Mouse control (simple trackpad)
- [x] Keyboard control (navigation)
- [ ] Keyboard control (text input)