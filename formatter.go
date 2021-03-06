package logrus

import "time"

const defaultTimestampFormat = time.RFC3339

const (
	errorKey   = "error"
	messageKey = "msg"
	timeKey    = "time"
	levelKey   = "level"
)

// The formatter interface is used to implement a custom formatter. It takes an
// `Entry`. It exposes all the fields, including the default ones:
//
// * `entry.Data["msg"]`. The Message passed from Info, Warn, Error ..
// * `entry.Data["time"]`. The timestamp.
// * `entry.Data["level"]. The level the entry was logged at.
//
// Any additional fields added with `WithField` or `WithFields` are also in
// `entry.Data`. Format is expected to return an array of bytes which are then
// logged to `Logger.out`.
type Formatter interface {
	Format(*Entry) ([]byte, error)
}

// This is to not silently overwrite `time`, `msg` and `level` fields when
// dumping it. If this code wasn't there doing:
//
//  logrus.WithField("level", 1).Info("hello")
//
// Would just silently drop the user provided level. Instead with this code
// it'll logged as:
//
//  {"level": "info", "fields.level": 1, "msg": "hello", "time": "..."}
//
// It's not exported because it's still using Data in an opinionated way. It's to
// avoid code duplication between the two default formatters.
func prefixFieldClashes(data Fields) {
	if t, ok := data[timeKey]; ok {
		data["fields.time"] = t
	}

	if m, ok := data[messageKey]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data[levelKey]; ok {
		data["fields.level"] = l
	}
}
