package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	parallel := flag.Int("parallel", 1, "sets the limit of maximum parallel goroutines")
	timeout := flag.Duration("timeout", time.Second, "time limit for connection")
	flag.Parse()

	args := flag.Args()

	fmt.Println("Parallel is:", *parallel)
	fmt.Println("Timeout is:", *timeout)
	fmt.Println("Args are:", args)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("Flag %s was setted with %v\n", f.Name, f.Value)
	})
}
