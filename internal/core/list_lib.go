package core

import (
	"adrianizen/library.id/internal/repository/keypair_repo"
	"adrianizen/library.id/internal/repository/word_repo"
)

func GetList() ([]keypair_repo.Keypair, error) {
	keypairRepo := keypair_repo.KeypairRepo{}
	dataKeypair, err := keypairRepo.GetList()

	return dataKeypair, err
}

func GetWordsLib(wordIDs []uint) ([]word_repo.Word, error) {
	wordRepo := word_repo.WordRepo{}
	dataWords, err := wordRepo.GetByIDs(wordIDs)

	return dataWords, err
}
