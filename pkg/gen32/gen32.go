// Package gen32 prvides necessary functions and variables for genCRC32 hashing
package gen32

var (
	// CRC32Polynomials CRC32 polynomials currently known to produce non-collision hashes after preprocessing
	CRC32Polynomials = []uint32{
		0x8741C726,
		0x87496166,
		0x8E2371EF,
		0x8EE5368F,
		0x8EFD4BCD,
		0x945D045D,
		0x9D9947FD,
		0xEB31D82E,
	}
)

// PreprocessBytes transforms DNA sequence into intermediate byte slice for appropriate CRC32 non-collision hashing
func PreprocessBytes(seq []byte) []byte {
	ret := make([]byte, len(seq)/2+len(seq)%2)
	for i, b := range seq {
		ret[i/2] <<= 4
		ret[i/2] |= b & 0b1110
	}
	return ret
}
