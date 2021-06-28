package stream

// Producer interface
type Producer interface {
	ToProtoBytes(recordValue []byte, sbj string) []byte
	Write(msg []byte, topic string) error
	Close()
}

// SchemaRegistry interface
type SchemaRegistry interface {
	GetSchemaID(subj string) (int, error)
	getSubjPath(sbj string) string
}
