package testutils

import (
	"encoding/json"
	"flesh/app/controllers"
	"flesh/app/models"
	sjs "github.com/bitly/go-simplejson"
	"github.com/robfig/revel"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"
)

var (
	cachedData sjs.Json
	allWords   []string

	TwoDaysBack, _    = time.ParseDuration("-48h")
	TwoDaysForward, _ = time.ParseDuration("48h")
	OneDayBack, _     = time.ParseDuration("-24h")
	OneDayForward, _  = time.ParseDuration("24h")
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
		for i := 0; i < 2; i++ {
			InsertTestOrganization()
		}
		revel.INFO.Print("Inserting random Users")
		for i := 0; i < 200; i++ {
			InsertTestUser()
		}
		revel.INFO.Print("Inserting random Members")
		for i := 0; i < 100; i++ {
			InsertTestMember()
		}
		revel.INFO.Print("Inserting random Games")
		for i := 0; i < 2; i++ {
			InsertTestGame()
		}
		revel.INFO.Print("Inserting random Players")
		for i := 0; i < 50; i++ {
			InsertTestPlayer()
		}
		revel.INFO.Print("Inserting random OZ Candidates")
		for i := 0; i < 5; i++ {
			InsertTestOzPool()
		}
		revel.INFO.Print("Inserting random OZs")
		for i := 0; i < 5; i++ {
			InsertTestOz()
		}
		revel.INFO.Print("Confirming random OZs")
		for i := 0; i < 1; i++ {
			ConfirmRandomOz()
		}
		revel.INFO.Print("Simulating tags by OZs")
		for i := 0; i < 5; i++ {
			TagByRandomOzs()
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
	return GenerateString(2, " ")
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
	first_name := GenerateWord().(string)
	last_name := GenerateWord().(string)
	sep := "."
	screen_name_long := first_name + sep + last_name
	screen_name := screen_name_long[:int(math.Min(20, float64(len(screen_name_long))))]
	email := screen_name + "@gmail.com"
	user, statusCode, err := models.NewUser(email, first_name, last_name, screen_name, "", "password")
	if err != nil {
		revel.WARN.Print(statusCode, err)
	}
	return user
}

func InsertTestOrganization() *models.Organization {
	name := GenerateName().(string)
	slug := strings.Replace(name, " ", "_", -1)
	org := &models.Organization{0, name + " university", slug, GenerateWord().(string), "US/Pacific", "A testing organization", models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(org)
	if err != nil {
		revel.WARN.Print(err)
	}
	return org
}

func InsertTestGame() *models.Game {
	org := SelectTestOrganization()
	now := time.Now()
	twoDaysAgo := now.Add(TwoDaysBack)
	twoDaysHence := now.Add(TwoDaysForward)
	oneDayAgo := now.Add(OneDayBack)
	oneDayHence := now.Add(OneDayForward)
	name := GenerateName().(string)
	slug := strings.Replace(name, " ", "_", -1)
	game := &models.Game{0,
		name,
		slug,
		org.Id,
		"US/Pacific",
		&twoDaysAgo,
		&oneDayHence,
		&oneDayAgo,
		&twoDaysHence,
		"A testing game",
		models.TimeTrackedModel{},
	}
	err := controllers.Dbm.Insert(game)
	if err != nil {
		revel.WARN.Print(err)
	}
	return game
}

func InsertTestPlayer() *models.Player {
	user := SelectTestUser()
	game := SelectTestGame()
	player := &models.Player{0, user.Id, game.Id, nil, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(player)
	if err != nil {
		revel.WARN.Print(err)
	}
	human_code := models.HumanCode{player.Id, "", models.TimeTrackedModel{}}
	human_code.GenerateCode()
	err = controllers.Dbm.Insert(&human_code)
	if err != nil {
		revel.WARN.Print(err)
	}
	err = models.CreateJoinedGameEvent(player)
	if err != nil {
		revel.WARN.Print(err)
	}
	return player
}

func InsertTestMember() *models.Member {
	user := SelectTestUser()
	organization := SelectTestOrganization()
	member := &models.Member{0, user.Id, organization.Id, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(member)
	if err != nil {
		revel.WARN.Print(err)
	}
	return member
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
	oz := &models.Oz{ozPool.Id, false, models.TimeTrackedModel{}}
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

func SelectTestConfirmedOz() *models.Oz {
	query := `
    SELECT *
    FROM "oz"
    WHERE confirmed = TRUE
    ORDER BY random()
    LIMIT 1
    `
	ozs, _ := controllers.Dbm.Select(models.Oz{}, query)
	oz := ozs[0].(*models.Oz)
	return oz
}

func SelectTestHuman(and string, args ...interface{}) *models.Player {
	query := `
    SELECT player.id Id, player.user_id User_id, player.game_id Game_id, player.created Created, player.updated Updated
    FROM "player"
    LEFT OUTER JOIN tag
    	ON player.id = taggee_id
    LEFT OUTER JOIN oz
    	on player.id = oz.id
    WHERE taggee_id IS NULL
    AND (oz.id IS NULL
    	 OR oz.confirmed = FALSE)
	` + and + `
    ORDER BY random()
    LIMIT 1
    `
	players, _ := controllers.Dbm.Select(models.Player{}, query, args...)
	player := players[0].(*models.Player)
	return player
}

func ConfirmRandomOz() {
	query := `
	SELECT oz.*
	FROM player p
	INNER JOIN  "oz"
		ON p.id = oz.id
	WHERE oz.confirmed = FALSE
	ORDER BY random()
	LIMIT 1
    `
	ozs, _ := controllers.Dbm.Select(models.Oz{}, query)
	oz := ozs[0].(*models.Oz)
	oz.Confirm()
}

func TagByRandomOzs() {
	oz := SelectTestConfirmedOz()
	oz_player, _ := models.PlayerFromId(oz.Id)
	human := SelectTestHuman("AND player.game_id = $1", oz_player.Game_id)
	if human == nil {
		return
	}
	now := time.Now()
	game, _ := models.GameFromId(human.Game_id)
	_, _, err := models.NewTag(game, oz_player, human, &now)
	if err != nil {
		revel.WARN.Print(err)
	}
}
