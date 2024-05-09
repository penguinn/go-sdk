package log

import (
	"sort"

	"github.com/sirupsen/logrus"
)

var JSONFormat = &logrus.JSONFormatter{
	CallerPrettyfier: CallerPrettyfier,
	TimestampFormat:  "2006-01-02 15:04:05",
}

var TextFormat = &logrus.TextFormatter{
	CallerPrettyfier: CallerPrettyfier,
	TimestampFormat:  "2006-01-02 15:04:05",
	DisableColors:    true,
	SortingFunc:      SortFunc,
}

func SortFunc(keys []string) {
	if len(keys) < 4 {
		return
	}
	sortedKeys := []string{}
	for _, key := range keys {
		if key != "time" && key != "level" && key != "file" && key != "msg" {
			sortedKeys = append(sortedKeys, key)
		}
	}
	sort.Strings(sortedKeys)
	keys[0] = "time"
	keys[1] = "level"
	keys[2] = "file"
	keys[3] = "msg"
	for i := 0; i < len(sortedKeys); i++ {
		keys[i+4] = sortedKeys[i]
	}
}
