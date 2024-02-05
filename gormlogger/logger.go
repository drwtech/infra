package gormlogger

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	gormlogger "gorm.io/gorm/logger"
)

type gormLoggerImpl struct {
	ctx           context.Context
	log           *log.Helper
	logLvl        gormlogger.LogLevel
	slowThreshold time.Duration
}

func NewDefault(log log.Logger) gormlogger.Interface {
	return NewGormLogger(context.Background(), log, 5*time.Second)
}

// NewGormLogger return gormlogger.Interface
func NewGormLogger(ctx context.Context, logger log.Logger, slowThreshold time.Duration) gormlogger.Interface {
	return &gormLoggerImpl{
		ctx:           ctx,
		log:           log.NewHelper(logger),
		logLvl:        gormlogger.Info,
		slowThreshold: slowThreshold,
	}
}

func (l *gormLoggerImpl) LogMode(lvl gormlogger.LogLevel) gormlogger.Interface {
	l.logLvl = lvl
	return l
}

func (l *gormLoggerImpl) Info(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Infof(format, args...)
}

func (l *gormLoggerImpl) Warn(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Warnf(format, args...)
}

func (l *gormLoggerImpl) Error(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Errorf(format, args...)
}

func (l *gormLoggerImpl) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logLvl > gormlogger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && l.logLvl >= gormlogger.Error:
			sql, rows := fc()
			l.log.WithContext(ctx).Errorw("gorm trace error", "trace_type", "error",
				"elapsed", elapsed, "rows", rows, "sql", sql)
		case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLvl >= gormlogger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
			// rows == -1 means no rows
			l.log.WithContext(ctx).Warnw("gorm trace warn", "trace_type", "warn",
				"slow_log", slowLog, "elapsed", elapsed, "rows", rows, "sql", sql)
		case l.logLvl == gormlogger.Info:
			sql, rows := fc()
			l.log.WithContext(ctx).Infow("gorm trace info", "trace_type", "info",
				"file", utils.FileWithLineNum(), "elapsed", elapsed, "rows", rows, "sql", sql)
		}
	}
}
