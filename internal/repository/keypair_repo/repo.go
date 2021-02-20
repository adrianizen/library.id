package keypair_repo

import (
	"adrianizen/library.id/internal/config"
	"adrianizen/library.id/internal/pkg/time_pkg"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type KeypairRepo struct {
}

var databaseFile string = config.RootDirectory + "../../../../database/keypair.json"
var logFile string = config.RootDirectory + "../../../../database/keypair_log.json"

var keypairFileMutex sync.Mutex

func (w *KeypairRepo) openDB() ([]Keypair, error) {
	var dataKeypair []Keypair
	data, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return dataKeypair, err
	}
	err = json.Unmarshal(data, &dataKeypair)
	if err != nil {
		return dataKeypair, err
	}

	return dataKeypair, nil
}

func (w *KeypairRepo) openLogDB() ([]KeypairLog, error) {
	var dataKeypairLogs []KeypairLog
	data, err := ioutil.ReadFile(logFile)
	if err != nil {
		return dataKeypairLogs, err
	}
	err = json.Unmarshal(data, &dataKeypairLogs)
	if err != nil {
		return dataKeypairLogs, err
	}

	return dataKeypairLogs, nil
}

func (w *KeypairRepo) addLog(keyPairID uint, value string) error {
	dataLogs, err := w.openLogDB()
	if err != nil {
		return err
	}

	logsDbLen := len(dataLogs)

	t := time.Now()
	tString := t.Format(time_pkg.TIME_DB)

	insertedID := logsDbLen + 1
	dataLogs = append(dataLogs, KeypairLog{
		ID:        uint(insertedID),
		KeypairID: keyPairID,
		Value:     value,
		CreatedAt: tString,
		UpdatedAt: tString,
	})

	saveJson, err := json.Marshal(dataLogs)
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(saveJson)
	if err != nil {
		return err
	}

	return nil
}

func (w *KeypairRepo) Upsert(key string, value string) (uint, error) {
	keypairFileMutex.Lock()
	defer keypairFileMutex.Unlock()
	dataKeypair, err := w.openDB()

	dbLen := len(dataKeypair)
	var isIdxExist bool
	var idxExist, insertedID int

	for idx, k := range dataKeypair {
		if k.Key == key {
			isIdxExist = true
			idxExist = idx
			break
		}
	}

	t := time.Now()
	tString := t.Format(time_pkg.TIME_DB)
	if isIdxExist {
		currentData := dataKeypair[idxExist]
		dataKeypair[idxExist] = Keypair{
			ID:        currentData.ID,
			Key:       key,
			Value:     value,
			CreatedAt: currentData.CreatedAt,
			UpdatedAt: tString,
		}

		err = w.addLog(currentData.ID, value)
		if err != nil {
			return 0, err
		}
	} else {
		insertedID := dbLen + 1
		dataKeypair = append(dataKeypair, Keypair{
			ID:        uint(insertedID),
			Key:       key,
			Value:     value,
			CreatedAt: tString,
			UpdatedAt: tString,
		})
		err = w.addLog(uint(insertedID), value)
		if err != nil {
			return 0, err
		}
	}

	// save
	saveJson, err := json.Marshal(dataKeypair)
	f, err := os.OpenFile(databaseFile, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return 0, err
	}
	_, err = f.Write(saveJson)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	if idxExist != 0 {
		return uint(idxExist), nil
	}

	return uint(insertedID), nil
}
