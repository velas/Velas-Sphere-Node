package entropy

import (
	"bytes"
	"math"
)

func Shannon(data []byte) (entropy float64) {
	if data == nil {
		return 0
	}

	if len(data) < 1 {
		return 0
	}

	for i := 0; i < 256; i++ {
		px := float64(bytes.Count(data, []byte{byte(i)})) / float64(len(data))

		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}

	return entropy
}
