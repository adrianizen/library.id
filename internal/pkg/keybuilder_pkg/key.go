package keybuilder_pkg

import (
	"strconv"
	"strings"
)

func BuildKeyString(wordIDs []uint) string {
	var str []string
	for _, id := range wordIDs {
		convString := strconv.FormatUint(uint64(id), 10)
		str = append(str, convString)
	}

	return strings.Join(str, "_")
}
