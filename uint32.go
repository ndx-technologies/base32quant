package base32quant

// UInt32 is a base32 encoded uint32 into 7 bytes.
type UInt32 struct{ v uint32 }

func NewUInt32(v uint32) UInt32 { return UInt32{v} }

func (s UInt32) UInt32() uint32 { return s.v }

func (s UInt32) MarshalText() ([]byte, error) {
	b := make([]byte, 7)
	Encode(b, s.v)
	return b, nil
}

func (s *UInt32) UnmarshalText(b []byte) (err error) {
	if len(b) != 7 {
		return ErrInvalidLength
	}
	s.v, err = Decode[uint32](b)
	return err
}

func NewUInt32FromString(s string) (v UInt32, err error) {
	err = (&v).UnmarshalText([]byte(s))
	return v, err
}

func (s UInt32) String() string {
	b := make([]byte, 7)
	Encode(b, s.v)
	return string(b)
}
