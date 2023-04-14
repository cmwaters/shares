package shares

import "errors"

// Read provides a generic function. It interprets the version of the schema from
// the first share and if it is supported, it will call the appropriate Read function
func Read(s []Share) ([][]byte, []byte, error) {
	if len(s) == 0 {
		return nil, nil, ErrEmptyShares
	}
	switch Version(s[0]) {
	case Compact().Version():
		return Compact().Read(s)
	case Sparse().Version():
		return Sparse().Read(s)
	default:
		return nil, nil, &ErrUnsupportedVersion{Version: Version(s[0])}
	}
}

// Write provides a generic function. It interprets the version of the schema from
// the first share and if it is supported, it will call the appropriate Write function
func Write(namespace []byte, data [][]byte, version uint8) ([]Share, error) {
	switch version {
	case Compact().Version():
		return Compact().Write(namespace, data)
	case Sparse().Version():
		return Sparse().Write(namespace, data)
	default:
		return nil, &ErrUnsupportedVersion{Version: version}
	}
}

// Pad provides a generic function. It interprets the version of the schema from
// the first share and if it is supported, it will call the appropriate Pad function3d
func Pad(namespace []byte, len int, version uint8) []Share {
	switch version {
	case Compact().Version():
		return Compact().Pad(namespace, len)
	case Sparse().Version():
		return Sparse().Pad(namespace, len)
	default:
		return nil
	}
}

var ErrEmptyShares = errors.New("empty set of shares")
