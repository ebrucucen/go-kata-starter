package trialBigOWithElse

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"flag"
	"log"
	"pprof"
	"runtime"
)


var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

// This function will return a string array of length 10 with numbers from 1 to 10
func GetListUpto10WithElse() [10]string {
	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

	startTime := time.Now()
	returnVal := [10]string{}
	for i := 0; i < 10; i++ {
		if (i+1)%3 == 0 {
			returnVal[i] = "fizz"

		} else {
			returnVal[i] = strconv.Itoa(i + 1)
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n Elapsed time WithElse %d NanoSeconds\n", elapsedTime.Nanoseconds())

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close()
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }

	return returnVal
}

