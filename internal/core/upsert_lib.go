package core

import (
	"adrianizen/library.id/internal/pkg/keybuilder_pkg"
	"adrianizen/library.id/internal/repository/keypair_repo"
	"adrianizen/library.id/internal/repository/word_repo"
	"strings"
)

func UpsertData(keypair string, price string) (uint, error) {
	var resultID uint
	splitString := strings.Split(keypair, "_")

	wordRepo := word_repo.WordRepo{}
	returnIDs, err := wordRepo.SyncGet(splitString)
	if err != nil {
		return 0, err
	}
	keyBuilder := keybuilder_pkg.BuildKeyString(returnIDs)
	keypairRepo := keypair_repo.KeypairRepo{}

	resultID, err = keypairRepo.Upsert(keyBuilder, price)
	if err != nil {
		return 0, err
	}

	return resultID, nil
}
