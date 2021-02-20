package core

import "adrianizen/library.id/internal/repository/keypair_repo"

func GetLogList(keypairID uint) ([]keypair_repo.KeypairLog, error) {
	keypairRepo := keypair_repo.KeypairRepo{}
	dataKeypairLog, err := keypairRepo.GetLogList(keypairID)

	return dataKeypairLog, err
}
