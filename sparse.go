package shares

const SparseVersion uint8 = 2

type sparseSchema struct {
}

func Sparse() Schema {
	return &sparseSchema{}
}

func V2() Schema {
	return Sparse()
}

func (s *sparseSchema) Version() uint8 { return SparseVersion }

func (s *sparseSchema) Read([]Share) ([][]byte, []byte, error) {
	panic("not implemented")
}

func (s *sparseSchema) Write(namespace []byte, data [][]byte) ([]Share, error) {
	panic("not implemented")
}

func (s *sparseSchema) Pad(namespace []byte, len int) []Share {
	panic("not implemented")
}

func (s *sparseSchema) Counter() Counter {
	panic("not implemented")
}
