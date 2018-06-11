package benchmark

import (
	"testing"
	"time"
)

func BenchmarkTime(b *testing.B) {
	v := time.Now().Second();

	for true  {
		if (time.Now().Second() -v ) > 3 {
			break
		}
 	}
}
