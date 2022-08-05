package logging

type KafkaLogWriter struct {}

func (cursor *KafkaLogWriter) Write(data []byte) (code int, err error) {
	return 0, nil
}