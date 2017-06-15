package main

import (
	"./vakthund"
	"flag"
	"log"
)

var branchVar string
var scriptVar string
var devVar bool

func main() {

	flag.StringVar(&branchVar, "branch", "master", "The branch that should be watched and deployed to")
	flag.StringVar(&scriptVar, "script", "/scripts/test.sh", "Path to deploy script")

	flag.BoolVar(&devVar,"dev",false,"Toggle development mode")

	flag.Parse()

	log.Println("Vakthund by @dskeppstedt,", "version 1")
	log.Println("Guards", branchVar, ",voof!")
	log.Println("Deployment script:", scriptVar)
	log.Println("Dev mode:",devVar)
	vakthund.Start(branchVar, scriptVar,devVar)

}
