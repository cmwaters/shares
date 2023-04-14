package shares

import (
	"fmt"
)

// A share represents a byte slice belonging to a set of data (usually represented
// as one or more byte arrays) from the same namespace.
type Share []byte

// Version extracts the version of the scheme that the share has been
// encoded in. As per the specification, the first byte must always be
// the version number. 0 is reserved for invalid shares.
func Version(share Share) uint8

// Reader interface for reading data from a set of shares.
type Reader interface {
	// Read takes a set of shares belonging to the same namespace and returns
	// the original data, the namespace and an error if the shares are invalid.
	Read([]Share) ([][]byte, []byte, error)
}

// Writer interface for writing data to a set of shares.
type Writer interface {
	// Write takes a namespace and a set of data and returns a set of shares
	Write(namespace []byte, data [][]byte) ([]Share, error)

	// Pad allows an existing set of shares to be extended with additional
	// padding Shares. These must have the same namespace as the original shares.
	Pad(namespace []byte, len int) []Share
}

// Schema is the interface that must be implemented by all share schemas.
type Schema interface {
	Version() uint8
	Reader
	Writer
	Counter() Counter
}

// Counter is an interface for calculating the amount of shares a set of data
// will require without needing to actually encode the data.
type Counter interface {
	Add(dataLen int, remainingCapacity int) bool
	Shares() int
}

type ErrUnsupportedVersion struct {
	Version uint8
}

func (e *ErrUnsupportedVersion) Error() string {
	return fmt.Sprintf("unsupported version: %d", e.Version)
}
