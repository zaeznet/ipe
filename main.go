// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/dimiro1/ipe/ipe"
)

// Main function, initialze the system
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var filename = flag.String("config", "config.json", "Config file location")
	flag.Parse()

	printBanner()

	ipe.Start(*filename)
}

// Print a beautifull banner
func printBanner() {
	fmt.Print("\033[36m")
	fmt.Print(`
██╗██████╗ ███████╗
██║██╔══██╗██╔════╝
██║██████╔╝█████╗  
██║██╔═══╝ ██╔══╝  
██║██║     ███████╗
╚═╝╚═╝     ╚══════╝`)
	fmt.Println("\033[0m")
	fmt.Println("\033[32mWelcome to Ipê - Yet another Pusher server clone\033[0m")
	fmt.Println("\033[33mBy: Claudemiro Alves Feitosa Neto <dimiro1@gmail.com>\033[0m")
}
