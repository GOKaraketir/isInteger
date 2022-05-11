package test

import (
	"github.com/GOKaraketir/isInteger"
	"testing"
)

func TestInteger(t *testing.T) {
	if isInteger.IsInteger(int(0)) &&
		isInteger.IsInteger(uint(0)) &&
		isInteger.IsInteger(int8(0)) &&
		isInteger.IsInteger(int16(0)) &&
		isInteger.IsInteger(int32(0)) &&
		isInteger.IsInteger(int64(0)) &&
		isInteger.IsInteger(uint8(0)) &&
		isInteger.IsInteger(uint16(0)) &&
		isInteger.IsInteger(uint32(0)) &&
		isInteger.IsInteger(uint64(0)) {
	} else {
		t.Fatal("Test Failed")
	}
}

func TestNotInteger(t *testing.T) {
	if isInteger.IsInteger(struct{}{}) || isInteger.IsInteger(1.0) {
		t.Fatal("Test Failed")
	}
}
