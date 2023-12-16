package common

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func RuntimeProf() (close func()) {
	var cpuprofile = flag.Bool("cpuprofile", false, "write cpu profile to `file`")
	flag.Parse()

	if !*cpuprofile {
		return func() {}
	}

	fmt.Println("starting CPU profile")
	f, err := os.Create("default.pgo")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	close = func() {
		pprof.StopCPUProfile()
		f.Close()
	}
	return close
}
