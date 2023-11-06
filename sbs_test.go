package sbs

import (
	"bytes"
	"testing"
)

func TestByteSliceToString(t *testing.T) {
	if r := ByteSliceToString([]byte{'t', 'e', 's', 't', 'i', 'n', 'g'}); r != "testing" {
		t.FailNow()
	}
}

func TestByteSliceToStringAlt(t *testing.T) {
	if r := ByteSliceToStringAlt([]byte{'t', 'e', 's', 't', 'i', 'n', 'g'}); r != "testing" {
		t.FailNow()
	}
}

func TestStringToByteSlice(t *testing.T) {
	if r := StringToByteSlice("testing"); !bytes.Equal(r, []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}) {
		t.FailNow()
	}
}

func TestStringToByteSliceAlt(t *testing.T) {
	if r := StringToByteSliceAlt("testing"); !bytes.Equal(r, []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}) {
		t.FailNow()
	}
}

func BenchmarkByteSliceToString(b *testing.B) {
	arr := []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ByteSliceToString(arr)
	}
}

func BenchmarkByteSliceToStringAlt(b *testing.B) {
	arr := []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ByteSliceToStringAlt(arr)
	}
}

func BenchmarkByteSliceToStringStandard(b *testing.B) {
	arr := []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = string(arr)
	}
}

func BenchmarkStringToByteSlice(b *testing.B) {
	str := "testing"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringToByteSlice(str)
	}
}

func BenchmarkStringToByteSliceAlt(b *testing.B) {
	str := "testing"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringToByteSliceAlt(str)
	}
}

func BenchmarkStringToByteSliceStandard(b *testing.B) {
	str := "testing"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = []byte(str)
	}
}
