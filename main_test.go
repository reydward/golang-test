package main

import (
	"flag"
	"log"
	"testing"
)

var envFile = flag.String("envFile", "", "Name of environment file")

func TestGreet(t *testing.T) {
	log.Print("=========================> envFile: " + *envFile)
}
