package chbs

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type wlErr string

func (w wlErr) Error() string {
	return string(w)
}

var WordList []string

func loadWordList(fName string) (wlArr []string, length int, err error) {
	wlFile, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, -1, err
	}
	wlArr = strings.Split(string(wlFile), ";")
	if len(wlArr) < 1 {
		err = wlErr("Length of WordList array is less than 1")
	}
	return
}

func firstUpper(s string) string {
	if len(s) > 1 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return strings.ToUpper(s)
}

func NewCHBS(num int, sep string) (chbs string) {
	for i := 0; i < num-1; i++ {
		chbs += firstUpper(WordList[rand.Intn(len(WordList))+1]) + sep
	}
	chbs += firstUpper(WordList[rand.Intn(len(WordList))+1])
	return
}

func init() {
	var err error
	WordList, _, err = loadWordList("wordlist.txt")
	if err != nil {
		log.Fatal("Loading WordList failed: " + err.Error())
	}
	rand.Seed(time.Now().UnixNano())
	_ = WordList
}
