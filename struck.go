package struck

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const (
	KeyTimestamp = "Timestamp"
	KeyHost      = "Host"
)

type Struck struct {
	ConstantLabels map[string]string
	Labels         []string
}

func (s *Struck) Log(labels ...interface{}) {
	if len(labels) != len(s.Labels) {
		panic("Number of labels does not match.")
	}
	m := map[string]interface{}{}
	for i, label := range labels {
		m[s.Labels[i]] = label
	}
	for k, v := range s.ConstantLabels {
		m[k] = v
	}
	m[KeyTimestamp] = time.Now().Unix()
	m[KeyHost], _ = os.Hostname()
	err := json.NewEncoder(os.Stdout).Encode(m)
	if err != nil {
		log.Fatalf("Failed to log message", err)
	}

}
