package demo

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type formatterFunc func(entry *log.Entry) ([]byte, error)

func (fn formatterFunc) Format(entry *log.Entry) ([]byte, error) {
	return fn(entry)
}

type levelIcon log.Level

func (l levelIcon) String() string {
	switch log.Level(l) {
	case log.InfoLevel:
		return "🛈"
	case log.WarnLevel:
		return "⚠"
	case log.ErrorLevel:
		fallthrough
	case log.FatalLevel:
		fallthrough
	case log.PanicLevel:
		return "💀"
	}
	return " "
}

var demoFormatter formatterFunc = func(entry *log.Entry) ([]byte, error) {
	if entry.Level > log.FatalLevel && !strings.HasPrefix(entry.Caller.Func.Name(), "github.com/Zenika/marcel/standalone/demo") {
		return nil, nil
	}

	s := fmt.Sprintf("%s  %s", levelIcon(entry.Level), entry.Message)
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}

	return []byte(s), nil
}
