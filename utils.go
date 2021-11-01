package traq_ws_bot

import "time"

func min(d1, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d1
	}
	return d2
}
