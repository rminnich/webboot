package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/u-root/webboot/pkg/menu"
)

var (
	v         = flag.Bool("verbose", false, "Verbose output")
	verbose   = func(string, ...interface{}) {}
	dir       = flag.String("dir", "", "Path of cached directory")
	network   = flag.Bool("network", true, "If network is false we will not set up network")
	dryRun    = flag.Bool("dryrun", false, "If dry_run is true we won't boot the iso.")
	logBuffer bytes.Buffer
)

func handleError(err error) {
	if err == menu.ExitRequest {
		menu.Close()
		os.Exit(0)
	} else if err == menu.BackRequest {
		return
	}

	errorText := err.Error() + "\n" + logBuffer.String() + wifiStdout.String() + wifiStderr.String()
	menu.DisplayResult(strings.Split(errorText, "\n"), ui.PollEvents())

	logBuffer.Reset()
	wifiStdout.Reset()
	wifiStderr.Reset()
}

func main() {

	flag.Parse()
	if *v {
		verbose = log.Printf
	}

	if err := menu.Init(); err != nil {
		log.Fatalf(err.Error())
	}

	// Buffer the log output, else it might overlap with the menu
	log.SetOutput(&logBuffer)

	for {
		if err := setupNetwork(ui.PollEvents()); err != nil {
			handleError(err)
		}
	}

}
