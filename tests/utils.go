package tests

import (
	"encoding/json"
	sjs "github.com/bitly/go-simplejson"
	"github.com/robfig/revel"
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

func GenerateString(numWords int, sep string) interface{} {
	words := GenerateRandomWordArray(numWords)
	return strings.Join(words, sep)
}

func GenerateWord() interface{} {
	return GenerateRandomWordArray(1)[0]
}

func GenerateName() interface{} {
	return GenerateString(0, " ")
}

func GenerateSlug() interface{} {
	return GenerateString(2, "_")
}

func GenerateEmail() interface{} {
	return GenerateString(0, "-").(string) + "@" + GenerateString(1, "").(string)
}

func GenerateStructArray(keyToGenerator map[string]func() interface{}, numEntries int) []map[string]interface{} {
	if numEntries < 0 {
		numEntries = rand.Intn(5)
	}
	userStructure := make([]map[string]interface{}, rand.Intn(5)+1)
	for i := 0; i < len(userStructure); i++ {
		userStructure[i] = make(map[string]interface{})
		for key, valFunc := range keyToGenerator {
			userStructure[i][key] = valFunc()
		}
	}

	return userStructure
}

func GenerateJson(underKey string, keyToGenerator map[string]func() interface{}, numEntries int) string {
	userStructure := GenerateStructArray(keyToGenerator, numEntries)
	underJsonKey := make(map[string][]map[string]interface{})
	underJsonKey[underKey] = userStructure
	jsonBytes, err := json.Marshal(underJsonKey)
	if err != nil {
		panic(err)
	}

	revel.WARN.Print(string(jsonBytes))

	return string(jsonBytes)
}
