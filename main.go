package main

import (
	"fmt"
	"log"

	"github.com/Vico1993/vdocker/docker"
)

func main() {
	dockerList, err := docker.Ps()
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, d := range dockerList {
		fmt.Println(d)
	}
}
