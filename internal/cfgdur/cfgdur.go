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
	// 配置里写成纯数字（如 5）时按「秒」解释：
	// time.Duration 以纳秒为单位，直接赋值会把 5 当成 5 纳秒，
	// 导致请求几乎立刻超时。
	*d = Duration(time.Duration(n) * time.Second)
	return nil
}

func (d Duration) Duration() time.Duration { return time.Duration(d) }
