package exception_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/exception"
)

func TestExceptionNewAndTrim(t *testing.T) {
	// Standard exception
	ex1 := exception.New("user not found", 404)
	assert.Equal(t, 404, ex1.Code)
	assert.Equal(t, "user not found", ex1.Message)

	// Formatted message
	ex2 := exception.New("failed to process %s", 500, "payment")
	assert.Equal(t, 500, ex2.Code)
	assert.Equal(t, "failed to process payment", ex2.Message)

	// Embedded Exception pattern
	ex3 := exception.New("Exception|401: unauthorized token", 500)
	assert.Equal(t, 401, ex3.Code)
	assert.Equal(t, "unauthorized token", ex3.Message)

	// Trim Exception|...
	msg1 := exception.Trim(errors.New("Exception|403: forbidden access"))
	assert.Equal(t, "forbidden access", msg1)

	// Trim Error: ...
	msg2 := exception.Trim(errors.New("Error: timeout reached"))
	assert.Equal(t, "timeout reached", msg2)

	// Trim normal error
	msg3 := exception.Trim(errors.New("io EOF"))
	assert.Equal(t, "io EOF", msg3)

	// Trim nil error
	assert.Equal(t, "", exception.Trim(nil))
}

func TestExceptionThrowCatch(t *testing.T) {
	defer func() {
		err := exception.Catch(recover())
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Exception|500: fatal database error")
	}()

	exception.New("fatal database error", 500).Throw()
}

func BenchmarkExceptionNewStandard(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = exception.New("simple static error", 400)
	}
}

func BenchmarkExceptionNewEmbedded(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = exception.New("Exception|404: resource not found", 500)
	}
}

func BenchmarkExceptionTrim(b *testing.B) {
	err := fmt.Errorf("Exception|404: resource not found")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = exception.Trim(err)
	}
}
