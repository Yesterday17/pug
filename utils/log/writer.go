package log

type logWriter struct {
	output func(string)
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	l.output(string(p))
	return len(p), nil
}
