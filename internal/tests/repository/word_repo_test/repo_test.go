package word_repo_test

import (
	"adrianizen/library.id/internal/repository/word_repo"
	"fmt"
	"strings"
	"testing"
)

func TestSyncGet(t *testing.T) {
	testString := "Jakarta_Bandung_Port"
	wordRepo := word_repo.WordRepo{}

	splitString := strings.Split(testString, "_")
	returnIDs, err := wordRepo.SyncGet(splitString)

	fmt.Println(returnIDs)
	fmt.Println(err)
}
