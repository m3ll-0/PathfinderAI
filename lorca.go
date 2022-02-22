package main

import (
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

// Go types that are bound to the UI must be thread-safe, because each binding
// is executed in its own goroutine. In this simple case we may use atomic

func setupLorca(){
	ui.SetBounds(lorca.Bounds{
		WindowState: lorca.WindowStateMaximized,
	})

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	ui.Bind("loadNextBoard", loadNextBoard)
	ui.Bind("doPathFinding", doPathFinding)
	ui.Bind("stopSolvers", stopSolvers)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
