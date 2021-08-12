package msgpack

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/vmihailenco/msgpack"
)

func TestEncodeInt8(t *testing.T) {
	var b []byte

	b = AppendInt8(b, 20)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	n, err := dec.DecodeInt8()
	if err != nil {
		t.Fatal(err)
	}

	if n != 20 {
		t.Fatalf("Differs: %d<>%d", n, 20)
	}
}

func TestEncodeInt64(t *testing.T) {
	var b []byte

	b = AppendInt64(b, -20)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	n, err := dec.DecodeInt64()
	if err != nil {
		t.Fatal(err)
	}

	if n != -20 {
		t.Fatalf("Differs: %d<>%d", n, -20)
	}
}

func TestEncodeFloat32(t *testing.T) {
	var b []byte

	b = AppendFloat32(b, 20.1234)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	n, err := dec.DecodeFloat32()
	if err != nil {
		t.Fatal(err)
	}

	if n != 20.1234 {
		t.Fatalf("Differs: %f<>%f", n, 20.1234)
	}
}

func TestEncodeFloat64(t *testing.T) {
	var b []byte

	b = AppendFloat64(b, 20.1234)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	n, err := dec.DecodeFloat64()
	if err != nil {
		t.Fatal(err)
	}

	if n != 20.1234 {
		t.Fatalf("Differs: %f<>%f", n, 20.1234)
	}
}

func TestEncodeSmallString(t *testing.T) {
	var b []byte

	b = AppendString(b, "a")

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	s, err := dec.DecodeString()
	if err != nil {
		t.Fatal(err)
	}

	if s != "a" {
		t.Fatalf("Differs: %s<>a", s)
	}
}

func TestEncodeMidString(t *testing.T) {
	var b []byte

	a := strings.Repeat("a", 16553)

	b = AppendString(b, a)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	s, err := dec.DecodeString()
	if err != nil {
		t.Fatal(err)
	}

	if s != a {
		t.Fatalf("Differs: %s<>a", s)
	}
}

func TestEncodeLongString(t *testing.T) {
	var b []byte

	a := strings.Repeat("a", 32220)

	b = AppendString(b, a)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	s, err := dec.DecodeString()
	if err != nil {
		t.Fatal(err)
	}

	if s != a {
		t.Fatalf("Differs: %s<>a", s)
	}
}

func TestEncodeTime(t *testing.T) {
	ts := time.Now()

	b := AppendTime(nil, ts)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	x, err := dec.DecodeTime()
	if err != nil {
		t.Fatal(err)
	}

	if ts.UnixNano() != x.UnixNano() {
		t.Fatalf("Differs: %d<>%d", ts.UnixNano(), x.UnixNano())
	}
}

func TestEncodeTimeSeconds(t *testing.T) {
	ts := time.Unix(
		time.Now().Unix(), 0)

	b := AppendTime(nil, ts)

	bf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(bf)

	x, err := dec.DecodeTime()
	if err != nil {
		t.Fatal(err)
	}

	if ts.UnixNano() != x.UnixNano() {
		t.Fatalf("Differs: %d<>%d", ts.UnixNano(), x.UnixNano())
	}
}
