package keypair_repo_test

import (
	"adrianizen/library.id/internal/pkg/keybuilder_pkg"
	"adrianizen/library.id/internal/repository/keypair_repo"
	"adrianizen/library.id/internal/repository/word_repo"
	"fmt"
	"strings"
	"testing"
)

func TestUpsert(t *testing.T) {
	testString := "Jakarta_Bandung_Port"
	wordRepo := word_repo.WordRepo{}

	splitString := strings.Split(testString, "_")
	returnIDs, err := wordRepo.SyncGet(splitString)
	keyBuilder := keybuilder_pkg.BuildKeyString(returnIDs)

	keypairRepo := keypair_repo.KeypairRepo{}

	resultID, err := keypairRepo.Upsert(keyBuilder, "100000")

	fmt.Println(keyBuilder)
	fmt.Println(resultID)
	fmt.Println(err)
}
