package tests

import (
	"encoding/json"
	"flesh/app/controllers"
	"flesh/app/models"
	sjs "github.com/bitly/go-simplejson"
	"github.com/robfig/revel"
	"io/ioutil"
	"math/rand"
	"strings"
)

const (
	JSON_CONTENT string = "application/json"
)

var (
	cachedData *sjs.Json
)

type FleshTest struct {
	revel.TestSuite
}

func (t *FleshTest) Before() {
	TestInit()
}

func (t *FleshTest) After() {
	TestClean()
}

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
	return GenerateString(0, "-").(string) + "@" + GenerateString(1, "").(string) + ".com"
}

func GenerateStructArray(keyToGenerator map[string]func() interface{}, numEntries int) []map[string]interface{} {
	if numEntries < 0 {
		numEntries = rand.Intn(5) + 1
	}
	userStructure := make([]map[string]interface{}, numEntries)
	for i := 0; i < len(userStructure); i++ {
		userStructure[i] = make(map[string]interface{})
		for key, valFunc := range keyToGenerator {
			userStructure[i][key] = valFunc()
		}
	}

	return userStructure
}

func GenerateJsonBytes(underKey string, keyToGenerator map[string]func() interface{}, numEntries int) []byte {
	userStructure := GenerateStructArray(keyToGenerator, numEntries)
	underJsonKey := EmbedMapUnderKey(underKey, userStructure)
	jsonBytes, err := json.Marshal(underJsonKey)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func GenerateJson(underKey string, keyToGenerator map[string]func() interface{}, numEntries int) string {
	return string(GenerateJsonBytes(underKey, keyToGenerator, numEntries))
}

func EmbedMapUnderKey(underKey string, mp []map[string]interface{}) map[string][]map[string]interface{} {
	underJsonKey := make(map[string][]map[string]interface{})
	underJsonKey[underKey] = mp
	return underJsonKey
}

func ConvertMappedStructArrayToBytes(mappedStructArray map[string][]map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(mappedStructArray)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func ConvertMappedStructArrayToString(mappedStructArray map[string][]map[string]interface{}) string {
	return string(ConvertMappedStructArrayToBytes(mappedStructArray))
}

func InsertTestUser() *models.User {
	user := &models.User{0, GenerateEmail().(string), GenerateName().(string), GenerateName().(string), GenerateSlug().(string), "", "", nil, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(user)
	if err != nil {
		revel.WARN.Print(err)
	}
	return user
}

func InsertTestOrganization() *models.Organization {
	org := &models.Organization{0, GenerateName().(string), GenerateSlug().(string), GenerateWord().(string), "US/Pacific", models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(org)
	if err != nil {
		revel.WARN.Print(err)
	}
	return org
}

func InsertTestGame() *models.Game {
	org := SelectTestOrganization()
	// make sure you have an organization available!
	game := &models.Game{0, GenerateName().(string), GenerateSlug().(string), org.Id, "US/Pacific", nil, nil, nil, nil, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(game)
	if err != nil {
		revel.WARN.Print(err)
	}
	return game
}

func SelectTestUser() *models.User {
	query := `
    SELECT *
    FROM "user"
    LIMIT 1
    `
	users, _ := controllers.Dbm.Select(models.User{}, query)
	user := users[0].(*models.User)
	return user
}

func SelectTestGame() *models.Game {
	query := `
    SELECT *
    FROM "game"
    LIMIT 1
    `
	games, _ := controllers.Dbm.Select(models.Game{}, query)
	game := games[0].(*models.Game)
	return game
}

func SelectTestOrganization() *models.Organization {
	query := `
    SELECT *
    FROM "organization"
    LIMIT 1
    `
	organizations, _ := controllers.Dbm.Select(models.Organization{}, query)
	organization := organizations[0].(*models.Organization)
	return organization
}
