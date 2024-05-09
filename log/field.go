package log

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func With(fields map[string]interface{}) *logrus.Entry {
	return WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return defaultLogger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return defaultLogger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return defaultLogger.WithError(err)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return defaultLogger.WithContext(ctx)
}

func WithTime(t time.Time) *logrus.Entry {
	return defaultLogger.WithTime(t)
}
