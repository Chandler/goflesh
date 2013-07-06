package testutils

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

var (
	cachedData sjs.Json
	allWords   []string
)

func init() {
	jsonBytes, err := ioutil.ReadFile("tests/test.json")
	if err != nil {
		panic(err)
	}

	cd, err := sjs.NewJson(jsonBytes)
	cachedData = *cd
	if err != nil {
		panic(err)
	}

	allWords, err = cachedData.GetPath("words", "all").StringArray()
	if err != nil {
		panic(err)
	}
}

func GenerateTestData() {
	isDev := revel.Config.BoolDefault("mode.dev", false)
	if isDev {
		revel.INFO.Print("Inserting random Organizations")
		for i := 0; i < 20; i++ {
			InsertTestOrganization()
		}
		revel.INFO.Print("Inserting random Users")
		for i := 0; i < 400; i++ {
			InsertTestUser()
		}
		revel.INFO.Print("Inserting random Games")
		for i := 0; i < 40; i++ {
			InsertTestGame()
		}
		revel.INFO.Print("Inserting random Players")
		for i := 0; i < 800; i++ {
			InsertTestPlayer()
		}
		revel.INFO.Print("Inserting random OZ Candidates")
		for i := 0; i < 160; i++ {
			InsertTestOzPool()
		}
		revel.INFO.Print("Inserting random OZs")
		for i := 0; i < 80; i++ {
			InsertTestOz()
		}
	}
}

func GenerateRandomWordArray(numWords int) []string {
	if numWords == 0 {
		numWords = rand.Intn(5) + 1
	}

	words := make([]string, numWords)

	lenNouns := len(allWords)
	for i := 0; i < numWords; i++ {
		index := rand.Intn(lenNouns)
		words[i] = allWords[index]
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
	game := &models.Game{0, GenerateName().(string), GenerateSlug().(string), org.Id, "US/Pacific", nil, nil, nil, nil, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(game)
	if err != nil {
		revel.WARN.Print(err)
	}
	return game
}

func InsertTestPlayer() *models.Player {
	user := SelectTestUser()
	game := SelectTestGame()
	player := &models.Player{0, user.Id, game.Id, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(player)
	if err != nil {
		revel.WARN.Print(err)
	}
	return player
}

func InsertTestOzPool() *models.OzPool {
	player := SelectTestPlayer()
	ozPool := &models.OzPool{player.Id, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(ozPool)
	if err != nil {
		revel.WARN.Print(err)
	}
	return ozPool
}

func InsertTestOz() *models.Oz {
	ozPool := SelectTestOzPool()
	oz := &models.Oz{ozPool.Id, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(oz)
	if err != nil {
		revel.WARN.Print(err)
	}
	return oz
}

func SelectTestUser() *models.User {
	query := `
    SELECT *
    FROM "user"
    ORDER BY random()
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
    ORDER BY random()
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
    ORDER BY random()
    LIMIT 1
    `
	organizations, _ := controllers.Dbm.Select(models.Organization{}, query)
	organization := organizations[0].(*models.Organization)
	return organization
}

func SelectTestPlayer() *models.Player {
	query := `
    SELECT *
    FROM "player"
    ORDER BY random()
    LIMIT 1
    `
	players, _ := controllers.Dbm.Select(models.Player{}, query)
	player := players[0].(*models.Player)
	return player
}

func SelectTestOzPool() *models.OzPool {
	query := `
    SELECT *
    FROM "oz_pool"
    ORDER BY random()
    LIMIT 1
    `
	ozPools, _ := controllers.Dbm.Select(models.OzPool{}, query)
	ozPool := ozPools[0].(*models.OzPool)
	return ozPool
}

func SelectTestOz() *models.Oz {
	query := `
    SELECT *
    FROM "oz"
    ORDER BY random()
    LIMIT 1
    `
	ozs, _ := controllers.Dbm.Select(models.Oz{}, query)
	oz := ozs[0].(*models.Oz)
	return oz
}
