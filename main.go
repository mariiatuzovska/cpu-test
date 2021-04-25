package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var (
	cpu    = flag.Bool("cpu", false, "Write CPU profile")
	memory = flag.Bool("mem", false, "Write memory profile")
)

func main() {

	flag.Parse()

	if *cpu {
		f, err := os.Create("./cpu-profile.pb.gz")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	go leakyFunction()
	time.Sleep(500 * time.Millisecond)

	if *memory {
		f, err := os.Create("./mem-profile.pb.gz")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	log.Println("Done.")
}

func leakyFunction() {
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pprof time")
	}
}
