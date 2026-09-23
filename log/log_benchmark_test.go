package log_test

import (
	"testing"

	"github.com/yaoapp/kun/log"
)

func BenchmarkLogTraceSimpleDisabled(b *testing.B) {
	log.SetLevel(log.InfoLevel) // Trace 关闭

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Trace("high frequency trace message without formatting")
	}
}

func BenchmarkLogTraceWithIsTraceCheck(b *testing.B) {
	log.SetLevel(log.InfoLevel)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if log.IsTrace() {
			log.Trace("expensive %v", map[string]int{"a": 1})
		}
	}
}

func BenchmarkLogTraceDisabled(b *testing.B) {
	log.SetLevel(log.InfoLevel)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Trace("high frequency metric %d %s", i, "data")
	}
}
