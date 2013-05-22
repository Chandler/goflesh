package tests

import (
	"encoding/json"
	sjs "github.com/bitly/go-simplejson"
	"io/ioutil"
	"math/rand"
	"strings"
)

var cachedData *sjs.Json

func GetTestData() *sjs.Json {
	if cachedData != nil {
		return cachedData
	}

	jsonBytes, err := ioutil.ReadFile("tests/test.json")
	if err != nil {
		panic(err)
	}

	cachedData, err = sjs.NewJson(jsonBytes)
	if err != nil {
		panic(err)
	}

	return cachedData
}

func GenerateRandomWordArray(numWords int) []string {
	if numWords == 0 {
		numWords = rand.Intn(5) + 1
	}

	words := make([]string, numWords)

	nouns, err := (*GetTestData()).GetPath("words", "nouns").StringArray()
	if err != nil {
		panic(err)
	}
	lenNouns := len(nouns)
	for i := 0; i < numWords; i++ {
		index := rand.Intn(lenNouns)
		words[i] = nouns[index]
	}

	return words
}

func GenerateString(numWords int, sep string) string {
	words := GenerateRandomWordArray(numWords)
	return strings.Join(words, sep)
}

func GenerateWord() string {
	return GenerateRandomWordArray(1)[0]
}

func GenerateName() string {
	return GenerateString(0, " ")
}

func GenerateSlug() string {
	return GenerateString(2, "_")
}

func GenerateEmail() string {
	return GenerateString(0, "-") + "@" + GenerateString(1, "")
}

func GenerateJson(keyToGenerator map[string]func() string, numEntries int) string {
	if numEntries < 0 {
		numEntries = rand.Intn(5)
	}
	userStructure := make([]map[string]string, rand.Intn(5)+1)
	for i := 0; i < len(userStructure); i++ {
		userStructure[i] = make(map[string]string)
		for key, valFunc := range keyToGenerator {
			userStructure[i][key] = valFunc()
		}
	}

	jsonBytes, err := json.Marshal(userStructure)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
