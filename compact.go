package shares

const CompactVersion uint8 = 1

type compactSchema struct {
	
}

func Compact() Schema {
	return &compactSchema{}
}

func V1() Schema {
	return Compact()
}

func (s *compactSchema) Version() uint8 { return CompactVersion }

func (s *compactSchema) Read([]Share) ([][]byte, []byte, error) {
	panic("not implemented")
}

func (s *compactSchema) Write(namespace []byte, data [][]byte) ([]Share, error) {
	panic("not implemented")
}

func (s *compactSchema) Pad(namespace []byte, len int) []Share {
	panic("not implemented")
}

func (s *compactSchema) Counter() Counter {
	panic("not implemented")
}
