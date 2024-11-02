package base32quant

import "errors"

var (
	ErrInvalidLength    = errors.New("invalid length")
	ErrInvalidCharacter = errors.New("invalid character")
)

var digits = [...]byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', '2', '3', '4', '5', '6', '7',
}

func Encode[T uint32 | uint64](b []byte, n T) error {
	for i := range len(b) {
		b[len(b)-i-1], n = digits[n&0x1f], n>>5
	}
	return nil
}

func Decode[T uint32 | uint64](b []byte) (n T, err error) {
	for i := 0; i < len(b); i++ {
		n <<= 5
		switch {
		case b[i] >= 'a' && b[i] <= 'z':
			n |= T(b[i] - 'a')
		case b[i] >= 'A' && b[i] <= 'Z':
			n |= T(b[i] - 'A')
		case b[i] >= '2' && b[i] <= '7':
			n |= T(b[i] - '2' + 26)
		default:
			return 0, ErrInvalidCharacter
		}
	}
	return n, nil
}
