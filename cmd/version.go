package cmd

import (
	"strconv"
	"time"
)

var (
	version    = "0.2.3"
	commitHash = "0fdda28e3c72d1128eb03c33b443112633c2f8aa"
	modifiedAt = "1597828236"
)

func lastModified() time.Time {
	return unixStringToTime(modifiedAt)
}

func unixStringToTime(unixStr string) time.Time {
	i, _ := strconv.ParseInt(unixStr, 10, 64)
	return time.Unix(i, 0).UTC()
}
