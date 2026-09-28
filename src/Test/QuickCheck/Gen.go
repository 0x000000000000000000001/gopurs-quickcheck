package Test_QuickCheck_Gen

import "math"

// Match the signed Int32 view of a JavaScript Float32Array.
func Float32ToInt32(n float64) int64 {
	return int64(int32(math.Float32bits(float32(n))))
}
