package msgpack

import (
	"time"
	"unsafe"
)

func AppendInt8(dst []byte, n int8) []byte {
	return append(dst, 0xd0, byte(n))
}

func AppendInt16(dst []byte, n int16) []byte {
	return append(dst, 0xd1, byte(n>>8), byte(n))
}

func AppendInt32(dst []byte, n int32) []byte {
	return append(dst, 0xd2,
		byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func AppendInt64(dst []byte, n int64) []byte {
	return append(dst, 0xd3,
		byte(n>>56), byte(n>>48),
		byte(n>>40), byte(n>>32),
		byte(n>>24), byte(n>>16),
		byte(n>>8), byte(n))
}

func AppendUint8(dst []byte, n int8) []byte {
	return append(dst, 0xcc, byte(n))
}

func AppendUint16(dst []byte, n int16) []byte {
	return append(dst, 0xcd, byte(n>>8), byte(n))
}

func AppendUint32(dst []byte, n int32) []byte {
	return append(dst, 0xce,
		byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func AppendUint64(dst []byte, n int64) []byte {
	return append(dst, 0xcf,
		byte(n>>56), byte(n>>48),
		byte(n>>40), byte(n>>32),
		byte(n>>24), byte(n>>16),
		byte(n>>8), byte(n))
}

func AppendFloat32(dst []byte, fl float32) []byte {
	ptr := *(*[4]byte)(unsafe.Pointer(&fl))
	return append(dst, 0xca,
		ptr[3], ptr[2], ptr[1], ptr[0])

}

// AppendFloat64 ...
func AppendFloat64(dst []byte, fl float64) []byte {
	ptr := *(*[8]byte)(unsafe.Pointer(&fl))
	return append(dst, 0xcb,
		ptr[7], ptr[6], ptr[5], ptr[4],
		ptr[3], ptr[2], ptr[1], ptr[0])

}

// AppendString ...
func AppendString(dst []byte, s string) []byte {
	dst = appendLen(dst, len(s))
	return append(dst, s...)
}

func appendLen(dst []byte, size int) []byte {
	switch {
	case size < 32:
		dst = append(dst, byte(5<<5|size))
	case size < 256:
		dst = append(dst, 0xd9, byte(size))
	case size < 65536:
		dst = append(dst, 0xda, byte(size>>8), byte(size))
	default:
		dst = append(dst, 0xdb,
			byte(size>>24), byte(size>>16),
			byte(size>>8), byte(size))
	}

	return dst
}

func AppendBytes(dst, s []byte) []byte {
	dst = appendByteLen(dst, len(s))
	return append(dst, s...)
}

func appendByteLen(dst []byte, size int) []byte {
	switch {
	case size < 256:
		dst = append(dst, 0xc4, byte(size))
	case size < 65536:
		dst = append(dst, 0xc5, byte(size>>8), byte(size))
	default:
		dst = append(dst, 0xc6,
			byte(size>>24), byte(size>>16),
			byte(size>>8), byte(size))
	}

	return dst
}

func AppendArrayLen(dst []byte, size int) []byte {
	switch {
	case size < 16:
		dst = append(dst, byte(9<<4|size))
	case size < 65536:
		dst = append(dst, 0xdc, byte(size>>8), byte(size))
	default:
		dst = append(dst, 0xdd,
			byte(size>>24), byte(size>>16),
			byte(size>>8), byte(size))
	}

	return dst
}

func AppendExt(dst []byte, kind byte, b []byte) []byte {
	dst = appendExtLen(dst, len(b))
	dst = append(dst, kind)
	return append(dst, b...)
}

func appendExtLen(dst []byte, size int) []byte {
	switch {
	case size == 1:
		dst = append(dst, 0xd4)
	case size == 2:
		dst = append(dst, 0xd5)
	case size == 4:
		dst = append(dst, 0xd6)
	case size == 8:
		dst = append(dst, 0xd7)
	case size == 16:
		dst = append(dst, 0xd8)
	case size < 256:
		dst = append(dst, 0xc7, byte(size))
	case size < 65536:
		dst = append(dst, 0xc8, byte(size>>8), byte(size))
	default:
		dst = append(dst, 0xc9,
			byte(size>>24), byte(size>>16),
			byte(size>>8), byte(size))
	}

	return dst
}

// AppendTime appends the given time to dst.
func AppendTime(dst []byte, ts time.Time) []byte {
	secs := ts.Unix()
	nanos := ts.Nanosecond()

	if nanos == 0 {
		dst = append(dst, 0xd6, 1,
			byte(secs>>24), byte(secs>>16),
			byte(secs>>8), byte(secs))
	} else {
		n := uint64(nanos)<<34 | uint64(secs)

		dst = append(dst, 0xd7, 1,
			byte(n>>56), byte(n>>48),
			byte(n>>40), byte(n>>32),
			byte(n>>24), byte(n>>16),
			byte(n>>8), byte(n))
	}

	return dst
}
