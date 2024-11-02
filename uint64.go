package base32quant

// UInt64 is a base32 encoded uint64 into 13 bytes.
type UInt64 struct{ v uint64 }

func NewUInt64(v uint64) UInt64 { return UInt64{v} }

func (s UInt64) UInt64() uint64 { return s.v }

func (s UInt64) MarshalText() ([]byte, error) {
	b := make([]byte, 13)
	Encode(b, s.v)
	return b, nil
}

func (s *UInt64) UnmarshalText(b []byte) (err error) {
	if len(b) != 13 {
		return ErrInvalidLength
	}
	s.v, err = Decode[uint64](b)
	return err
}

func NewUInt64FromString(s string) (v UInt64, err error) {
	err = (&v).UnmarshalText([]byte(s))
	return v, err
}

func (s UInt64) String() string {
	b := make([]byte, 13)
	Encode(b, s.v)
	return string(b)
}
