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
		for i := 0; i < 20; i++ {
			InsertTestOrganization()
		}
		revel.INFO.Print("Inserting random Users")
		for i := 0; i < 400; i++ {
			InsertTestUser()
		}
		revel.INFO.Print("Inserting random Members")
		for i := 0; i < 40; i++ {
			InsertTestMember()
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
		for i := 0; i < 100; i++ {
			InsertTestOz()
		}
		revel.INFO.Print("Confirming random OZs")
		for i := 0; i < 80; i++ {
			ConfirmRandomOz()
		}
		revel.INFO.Print("Simulating tags by OZs")
		for i := 0; i < 100; i++ {
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
	now := time.Now()
	user := &models.User{0, email, first_name, last_name, screen_name, "", "", nil, models.TimeTrackedModel{&now, &now}}
	err := controllers.Dbm.Insert(user)
	if err != nil {
		revel.WARN.Print(err)
	}
	return user
}

func InsertTestOrganization() *models.Organization {
	name := GenerateName().(string)
	slug := strings.Replace(name, " ", "_", -1)
	org := &models.Organization{0, name + " university", slug, GenerateWord().(string), "US/Pacific", models.TimeTrackedModel{}}
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
	player := &models.Player{0, user.Id, game.Id, models.TimeTrackedModel{}}
	err := controllers.Dbm.Insert(player)
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

func SelectTestHuman() *models.Player {
	query := `
    SELECT player.id, player.user_id, player.game_id, player.created, player.updated
    FROM "player"
    LEFT OUTER JOIN tag
    	ON player.id = taggee_id
    LEFT OUTER JOIN oz
    	on player.id = oz.id
    WHERE taggee_id IS NULL
    AND (oz.id IS NULL
    	 OR oz.confirmed = FALSE)
    ORDER BY random()
    LIMIT 1
    `
	players, _ := controllers.Dbm.Select(models.Player{}, query)
	player := players[0].(*models.Player)
	return player
}

func ConfirmRandomOz() {
	query := `
	UPDATE "oz"
	SET
		confirmed = TRUE
	WHERE id IN (
		SELECT id
		FROM "oz"
		WHERE confirmed = FALSE
		ORDER BY random()
		LIMIT 1
	)
    `
	controllers.Dbm.Exec(query)
}

func TagByRandomOzs() {
	oz := SelectTestConfirmedOz()
	oz_player, _ := models.PlayerFromId(oz.Id)
	human := SelectTestHuman()
	now := time.Now()
	game, _ := models.GameFromId(human.Game_id)
	models.NewTag(game, oz_player, human, &now)
}
