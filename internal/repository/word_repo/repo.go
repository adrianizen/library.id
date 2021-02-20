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

var databaseFile string = config.RootDirectory + "../../../../database/words.json"
var wordFileMutex sync.Mutex

func (w *WordRepo) SyncGet(words []string) ([]uint, error) {
	var r []uint
	var dataWords []Word

	wordFileMutex.Lock()
	defer wordFileMutex.Unlock()

	data, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return r, err
	}
	json.Unmarshal(data, dataWords)

	var unsyncedWords []string
	var unsyncedWordIDs []int

	wordsLen := len(words)
	lastID := len(dataWords)

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
	currentID := lastID
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
