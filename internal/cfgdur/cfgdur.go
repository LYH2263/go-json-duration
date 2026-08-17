package cfgdur

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var n int64
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	*d = Duration(time.Duration(n) * time.Second)
	return nil
}

func (d Duration) Duration() time.Duration { return time.Duration(d) }
