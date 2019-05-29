package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/Vico1993/vdocker/docker"
)

func main() {
	out, err := exec.Command("docker", "ps", "--format", "['{{.Image}}','{{.Names}}','{{.Status}}']").Output()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(bytes.NewReader(out))
	for i := 1; ; i++ {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		str := string(line[:])
		// first line exeption
		if strings.Contains(str, "CONTAINER ID") {
			continue
		}

		var d docker.Dockerps
		if err := json.Unmarshal([]byte(strings.ReplaceAll(str, "'", "\"")), &d); err != nil {
			log.Fatal(err)
		}
	}
}
