package docker

import (
	"encoding/json"
	"fmt"
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
