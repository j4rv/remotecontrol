package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/micmonay/keybd_event"

	"github.com/gorilla/websocket"
)

func main() {
	printIPs()
	initIndex()
	initActionHandlers()
	httpStart()
}

var upgrader = websocket.Upgrader{} // use default options

// printIPs Prints the device's IPs. For convenience.
func printIPs() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error while retrieving the network interfaces:", err)
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println("Error while retrieving an interface address:", err)
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				switch v.IP.String() {
				case "::1":
				case "127.0.0.1":
					continue
				default:
					log.Println("IP Network:", v.IP)
				}
			case *net.IPAddr:
				log.Println("IP Address:", v.IP)
			}
		}
	}
}

// httpStart executes the ListenAndServe command
func httpStart() {
	log.Println("Starting http server...")
	err := http.ListenAndServe(":80", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.RemoteAddr, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// initIndex handles the root page
func initIndex() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*err := tmpl.Execute(w, nil)
		if err != nil {
			log.Print("Error while executing the template:", err)
		}*/
		w.Write([]byte(indexTmpl))
	})
}

// initActionHandlers handles all available actions that can be called remotely
func initActionHandlers() {

	///////////
	// SOUND //
	///////////
	http.HandleFunc("/volumeUp", func(w http.ResponseWriter, r *http.Request) {
		logIfError(volumeUp())
	})
	http.HandleFunc("/volumeDown", func(w http.ResponseWriter, r *http.Request) {
		logIfError(volumeDown())
	})
	http.HandleFunc("/silence", func(w http.ResponseWriter, r *http.Request) {
		logIfError(silence())
	})
	http.HandleFunc("/nextSong", func(w http.ResponseWriter, r *http.Request) {
		logIfError(nextSong())
	})
	http.HandleFunc("/prevSong", func(w http.ResponseWriter, r *http.Request) {
		logIfError(prevSong())
	})
	http.HandleFunc("/pauseSong", func(w http.ResponseWriter, r *http.Request) {
		logIfError(pauseSong())
	})

	///////////////
	// SHUTDOWNS //
	///////////////
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		mins, err := strconv.Atoi(r.URL.Query()["mins"][0])
		if err != nil {
			log.Println(err)
			return
		}
		logIfError(shutdownInSecs(mins * 60))
	})
	http.HandleFunc("/abortShutdown", func(w http.ResponseWriter, r *http.Request) {
		logIfError(abortShutdown())
	})

	///////////
	// MOUSE //
	///////////
	http.HandleFunc("/mouseMoveWebSocket", handleMouseMoveWebSocket)
	http.HandleFunc("/leftClick", func(w http.ResponseWriter, r *http.Request) {
		logIfError(mouseClick("left"))
	})
	http.HandleFunc("/middleClick", func(w http.ResponseWriter, r *http.Request) {
		logIfError(mouseClick("center"))
	})
	http.HandleFunc("/rightClick", func(w http.ResponseWriter, r *http.Request) {
		logIfError(mouseClick("right"))
	})

	//////////////
	// KEYBOARD //
	//////////////
	http.HandleFunc("/keyEsc", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_ESC))
	})
	http.HandleFunc("/keyEnter", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_ENTER))
	})
	http.HandleFunc("/keyUp", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_UP))
	})
	http.HandleFunc("/keyDown", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_DOWN))
	})
	http.HandleFunc("/keyLeft", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_LEFT))
	})
	http.HandleFunc("/keyRight", func(w http.ResponseWriter, r *http.Request) {
		logIfError(keypress(keybd_event.VK_RIGHT))
	})

	////////////
	// OTHERS //
	////////////
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query()["msg"]
		log.Println(msg)
	})

}

func handleMouseMoveWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("could not upgrade:", err)
		return
	}
	defer c.Close()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		splitMsg := strings.Split(string(msg), ":")
		x, err := strconv.Atoi(splitMsg[0])
		logIfError(err)
		y, err := strconv.Atoi(splitMsg[1])
		logIfError(err)

		if err != nil {
			return
		}

		logIfError(moveMouse(x, y))
	}
}

func logIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
