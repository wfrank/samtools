package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/wfrank/samtools/pkg/isam"
)

func main() {

	hosts := flag.String("hosts", os.Getenv("ISAM_HOSTS"), "Comma separated ISAM hosts to monitor")
	user := flag.String("user", os.Getenv("ISAM_USER"), "Username with monitor privilege")
	pass := flag.String("pass", os.Getenv("ISAM_PASS"), "Password of the User")

	flag.Parse()

	if *hosts == "" || *user == "" || *pass == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, host := range strings.Split(*hosts, ",") {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			server := isam.NewClient(host, *user, *pass)
			stats, err := server.PollSystemStats()
			if err != nil {
				log.Printf("error polling system stats from %s: %v", host, err)
				return
			}
			fmt.Printf("%s: %+v\n", host, *stats)
		}(host)
	}
	wg.Wait()
}
