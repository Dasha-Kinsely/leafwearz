package logging

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func getRotateLogs(filename string, storageLength int) io.Writer {
	writer, err := rotatelogs.New(
		filename+".%Y-%m-%d_%H:%M:%S",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(storageLength*24)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return writer
}