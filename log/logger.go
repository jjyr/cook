package log

type Logger interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
}
