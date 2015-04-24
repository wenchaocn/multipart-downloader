package main

import (
	"flag"
	"log"
	"os"
	"time"
	md "github.com/alvatar/multipart-downloader"
)


var (
	nConns   = flag.Uint("n", 1, "Number of concurrent connections")
	sha256   = flag.String("S", "", "File containing SHA-256 hash, or a SHA-256 string")
	useEtag  = flag.Bool("E", false, "Verify using Etag as MD5")
	timeout  = flag.Uint("t", 5000, "Timeout for all connections in milliseconds")
	verbose  = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()
	log.SetPrefix("godl: ")
	if len(flag.Args()) == 0 {
		log.Fatal("No URLs provided")
		os.Exit(1)
	}

	// Initialize download
	dldr := md.NewMultiDownloader(flag.Args(), *nConns, time.Duration(*timeout) * time.Millisecond)
	md.SetVerbose(*verbose)

	// Gather info from all sources
	err := dldr.GatherInfo()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// ...
}
