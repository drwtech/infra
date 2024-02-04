package dcron

import "github.com/go-kratos/kratos/v2/log"

type DLog struct {
	Log *log.Helper
}

func (l *DLog) Infof(format string, args ...any) {
	l.Log.Infof(format, args)
}

func (l *DLog) Warnf(format string, args ...any) {
	l.Log.Warnf(format, args)
}

func (l *DLog) Errorf(format string, args ...any) {
	l.Log.Errorf(format, args)
}

func (l *DLog) Printf(format string, args ...any) {
	l.Log.Debugf(format, args)
}
