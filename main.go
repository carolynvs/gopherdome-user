package main

import (
	"fmt"

	"github.com/carolynvs/gopherdome-dump"
	"github.com/carolynvs/gopherdome-sdk"
)

func main() {
	fmt.Println("Welcome to the Gopher Thunderdome!")

	leGopher := gophersdk.GetGopher("Gopherina")
	gopherdump.Dump(leGopher)
}
