package main

import (
	// "context"
	// "fmt"
	// "log"
	// "os"

	"github.com/blizardinka/BattleHelperWebApp/internal"
	"github.com/blizardinka/BattleHelperWebApp/database"

	"github.com/FuzzyStatic/blizzard/v2"
	// "github.com/FuzzyStatic/blizzard/v2/hsgd"
)

var (
	clientID		string = internal.ClientID
	clientSecret	string = internal.ClientSecret
	blizz			*blizzard.Client
)	

// func init() {
// 	// clientID = os.Getenv("8a416153c3354c6eb59cd0aed8d11829")
// 	clientID = internal.ClientID
// 	if clientID == "" {
// 		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
// 	}

// 	// clientSecret = os.Getenv(configs.ClientID)
// 	clientSecret = internal.ClientSecret
// 	if clientSecret == "" {
// 		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
// 	}
// }

func main() {
	// blizz = blizzard.NewClient(clientID, clientSecret, blizzard.EU, blizzard.EnUS)

	// err := blizz.AccessTokenRequest(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// dat, _, err := blizz.HSBattlegroundsCardsSearch(
	// 	context.Background(),
	// 	"", "", "", "", "",
	// 	[]int{}, []int{}, []int{}, 0, 0,
	// 	[]hsgd.Tier{hsgd.TierHero},
	// 	"", "", "",)
	// if err != nil {
	// fmt.Println(err)
	// }

	// fmt.Printf("%+v\n", dat)

	database.TestConection()
}
