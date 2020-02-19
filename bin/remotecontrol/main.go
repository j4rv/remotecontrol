package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
)

func main() {
	printIPs()
	initIndex()
	initActionHandlers()
	httpStart()
}

// printIPs Prints the device's IPs. For convenience.
func printIPs() {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				log.Println("IP Network:", v.IP)
			case *net.IPAddr:
				log.Println("IP Address:", v.IP)
			}
		}
	}
}

// httpStart executes the ListenAndServe command
func httpStart() {
	log.Println("Starting http server...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// initIndex handles the root page
func initIndex() {
	tmpl := template.Must(template.New("index").Parse(indexTmpl))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Print("Error while executing the template: ", err)
		}
	})
}

// initActionHandlers handles all available actions that can be called remotely
func initActionHandlers() {
	http.HandleFunc("/volumeUp", func(w http.ResponseWriter, r *http.Request) {
		logIfError(volumeUp())
	})
	http.HandleFunc("/volumeDown", func(w http.ResponseWriter, r *http.Request) {
		logIfError(volumeDown())
	})
	http.HandleFunc("/silence", func(w http.ResponseWriter, r *http.Request) {
		logIfError(silence())
	})

	// Keep it simple, stupid. No generic way to call "shutdown(x)" from outside.
	http.HandleFunc("/shutdown1m", func(w http.ResponseWriter, r *http.Request) {
		logIfError(shutdownInSecs(1 * 60))
	})
	http.HandleFunc("/shutdown30m", func(w http.ResponseWriter, r *http.Request) {
		logIfError(shutdownInSecs(30 * 60))
	})
	http.HandleFunc("/shutdown60m", func(w http.ResponseWriter, r *http.Request) {
		logIfError(shutdownInSecs(60 * 60))
	})
	http.HandleFunc("/shutdown120m", func(w http.ResponseWriter, r *http.Request) {
		logIfError(shutdownInSecs(120 * 60))
	})
	http.HandleFunc("/abortShutdown", func(w http.ResponseWriter, r *http.Request) {
		logIfError(abortShutdown())
	})
}

func logIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
