package cfgdur

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDurationSeconds(t *testing.T) {
	var d Duration
	if err := json.Unmarshal([]byte("5"), &d); err != nil {
		t.Fatal(err)
	}
	if d.Duration() != 5*time.Second {
		t.Fatalf("got %v", d.Duration())
	}
}
