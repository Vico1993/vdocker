package docker

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Dockerps is the struct to handle docker ps output
type Dockerps struct {
	Image  string
	Names  string
	Status string
}

// UnmarshalJSON convert or JSON array into the struct
func (d *Dockerps) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&d.Image, &d.Names, &d.Status}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in DockerPS: %d != %d", g, e)
	}
	return nil
}

// Ps return the docker ps
func Ps() ([]Dockerps, error) {
	var res []Dockerps

	out, err := exec.Command("docker", "ps", "--format", "['{{.Image}}','{{.Names}}','{{.Status}}']").Output()
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(bytes.NewReader(out))
	for i := 1; ; i++ {
		line, _, err := r.ReadLine()
		if err != nil && err == io.EOF {
			return res, nil
		}

		str := string(line[:])
		// first line exeption
		if strings.Contains(str, "CONTAINER ID") {
			continue
		}

		var d Dockerps
		if e := json.Unmarshal([]byte(strings.ReplaceAll(str, "'", "\"")), &d); e != nil {
			return nil, err
		}

		res = append(res, d)
	}
}
