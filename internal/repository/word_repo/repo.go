package word_repo

import (
	"adrianizen/library.id/internal/config"
	"adrianizen/library.id/internal/pkg/time_pkg"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type WordRepo struct {
}

var wordFileMutex sync.Mutex

func (w *WordRepo) openDB() ([]Word, error) {
	var databaseFile string = config.RootDirectory + "/database/words.json"

	var dataWords []Word
	data, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return dataWords, err
	}
	err = json.Unmarshal(data, &dataWords)
	if err != nil {
		return dataWords, err
	}

	return dataWords, nil
}

func (w *WordRepo) SyncGet(words []string) ([]uint, error) {
	var databaseFile string = config.RootDirectory + "/database/words.json"
	var r []uint

	wordFileMutex.Lock()
	defer wordFileMutex.Unlock()

	dataWords, err := w.openDB()

	var unsyncedWords []string
	var unsyncedWordIDs []int

	wordsLen := len(words)
	wordsDbLen := len(dataWords)

	returnIDs := make([]uint, wordsLen)
	for idx, w := range words {
		exist := false
		for _, dWord := range dataWords {
			if dWord.Word == w {
				exist = true
				returnIDs[idx] = dWord.ID
			}
		}
		if exist == false {
			unsyncedWords = append(unsyncedWords, w)
			unsyncedWordIDs = append(unsyncedWordIDs, idx)
		}
	}

	t := time.Now()
	tString := t.Format(time_pkg.TIME_DB)
	currentID := wordsDbLen
	for idx, uWord := range unsyncedWords {
		unsyncedWordID := unsyncedWordIDs[idx]
		currentID = currentID + 1
		dataWords = append(dataWords, Word{
			ID:        uint(currentID),
			Word:      uWord,
			CreatedAt: tString,
		})

		returnIDs[unsyncedWordID] = uint(currentID)
	}

	// save
	saveJson, err := json.Marshal(dataWords)
	f, err := os.OpenFile(databaseFile, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return r, err
	}

	_, err = f.Write(saveJson)
	if err != nil {
		return r, err
	}

	defer f.Close()

	return returnIDs, nil
}

func (w *WordRepo) GetByIDs(ids []uint) ([]Word, error) {
	var databaseFile string = config.RootDirectory + "/database/words.json"

	var dataWords []Word
	var returnWords []Word
	data, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return dataWords, err
	}
	err = json.Unmarshal(data, &dataWords)
	if err != nil {
		return dataWords, err
	}

	for _, id := range ids {
		for _, w := range dataWords {
			if w.ID == id {
				returnWords = append(returnWords, w)
			}
		}
	}

	return returnWords, nil
}
