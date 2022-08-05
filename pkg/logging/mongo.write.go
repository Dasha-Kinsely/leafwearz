package logging

type MongoLogWriter struct {}

// TODO: implement this

func (cursor *MongoLogWriter) Write(data []byte) (code int, err error) {
	if err = cursor.InsertLogToMongo(data); err != nil {
		return 0, err
	}
	return
}

func (cursor *MongoLogWriter) InsertLogToMongo(data []byte) error {
	return nil
}