package service

import (
	"adrianizen/library.id/internal/config"
	"adrianizen/library.id/internal/core"
	"adrianizen/library.id/internal/repository/word_repo"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// var tplList = template.Must(template.ParseFiles(config.RootDirectory + "/html/list.html"))

func UpsertUIHandler(w http.ResponseWriter, r *http.Request) {
	var tplUpsert = template.Must(template.ParseFiles(config.RootDirectory + "/html/upsert.html"))
	tplUpsert.Execute(w, nil)
}

func UpsertHandler(w http.ResponseWriter, r *http.Request) {
	var tplUpsert = template.Must(template.ParseFiles(config.RootDirectory + "/html/upsert.html"))
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	keypair := r.FormValue("keypair")
	value := r.FormValue("value")

	id, err := core.UpsertData(keypair, value)

	fmt.Println(id)
	fmt.Println(err)
	tplUpsert.Execute(w, nil)
}

type KeypairTransformer struct {
	Key       string
	Value     string
	CreatedAt string
}

func ListUIHandler(w http.ResponseWriter, r *http.Request) {
	var tplKeypairList = template.Must(template.ParseFiles(config.RootDirectory + "/html/list.html"))
	keyPairList, err := core.GetList()
	if err != nil {
		fmt.Println(err)
		return
	}

	var ids []uint
	for _, k := range keyPairList {
		keypair := k.Key
		splitKeypair := strings.Split(keypair, "_")

		for _, s := range splitKeypair {
			u64, _ := strconv.ParseUint(s, 10, 32)
			sUint := uint(u64)
			ids = append(ids, sUint)
		}
	}

	wordList, err := core.GetWordsLib(ids)
	wordListMap := make(map[string]word_repo.Word)
	for _, w := range wordList {
		wordIDString := strconv.FormatUint(uint64(w.ID), 10)
		wordListMap[wordIDString] = w
	}

	var transformer []KeypairTransformer
	for _, k := range keyPairList {
		keypair := k.Key
		splitKeypair := strings.Split(keypair, "_")
		wordArrTransformed := []string{}
		for _, s := range splitKeypair {
			w := wordListMap[s]
			wordArrTransformed = append(wordArrTransformed, w.Word)
		}

		wordTransformed := strings.Join(wordArrTransformed, " - ")
		transformer = append(transformer, KeypairTransformer{
			Key:       wordTransformed,
			Value:     k.Value,
			CreatedAt: k.CreatedAt,
		})
	}
	tplKeypairList.Execute(w, transformer)
}
