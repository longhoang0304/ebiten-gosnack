package main

import (
	"go-snack/gosnack"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	file, err := os.Create("profile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	gosnack.SetupWindow()
	gosnack.RunGame()
}
