package main
import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"github.com/joho/godotenv"
)

type Song struct {
	ID					int		`json:"title"`
	Title				string	`json:"title"`
	ArtistID			int		`json:"artistId"`
	RomTitle			string	`json:"romTitle"`
	Key					int		`json:"key"`
	Mode				int		`json:"mode"`
	Tempo				int		`json:"tempo"`
	DurationMs			int		`json:"durationMs"`
	TimeSignature		string	`json:"timeSignature"`
	LanguageID			int		`json:"languageId"`
	Initialism			string	`json:"initialism"`
	SpotifyLink			string	`json:"spotifyLink"`
	YouTubeLink			string	`json:"youtubeLink"`
	OtherLink			string	`json:"otherLink"`
	Energy				float64	`json:"energy"`
	Danceability		float64	`json:"danceability"`
	Valence				float64	`json:"valence"`
	Acousticness		float64	`json:"acousticness"`
	Instrumentalness	float64	`json:"instrumentalness"`
	DateReleased		string	`json:"dateReleased"`
	Verified			bool	`json:"verified"`
}

type Genre struct {
	ID		int		`json:"id"`
	Name	string `json:"name"'`
}

type Language struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}

type Mood struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}

type Tag struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}

type Musician struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	RomName	string	`json:"romName"`
	EnName 	string	`json:"enName"`
}

func setup (db *gorm.DB) {
	err := db.AutoMigrate(&Song{})

	if err != nil {
		fmt.Println("Auto migration failed")
		fmt.Println(err)
	}

	err = db.AutoMigrate(&Musician{})

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env loading failed")
	}
	DSN := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to database")
	}

	sqlDb, sqlErr := db.DB()

	if sqlErr != nil {
		fmt.Println("something is wrong with db.DB()")
	}

	defer sqlDb.Close()

	setup(db)

	fmt.Println("database successfully connected")

	e := echo.New()

	e.GET("/", func (c echo.Context) error {
		var song Song
		var songs []Song
		db.First(&song)
		db.Limit(10).Find(&songs)
		fmt.Println(song)
		fmt.Printf("%+v\n\n", songs)

		return c.JSON(http.StatusOK, songs)
	})

	e.GET("/songs", func (c echo.Context) error {
		//
		return c.JSON(http.StatusOK, "working")
	})

	e.POST("/songs", func (c echo.Context) error {
		return c.JSON(http.StatusOK, "add song route")
	})

	e.PUT("/songs/:id", func (c echo.Context) error {
		return c.JSON(http.StatusOK, "edit song route")
	})



	e.DELETE("/songs/:id", func (c echo.Context) error {
		return c.JSON(http.StatusOK, "delete song route")
	})

	e.POST("/songs/spotify", func (c echo.Context) error {
		return c.JSON(http.StatusOK, "add song with Spotify link")
	})

	e.POST("/songs/csv", func (c echo.Context) error {
		return c.JSON(http.StatusOK, "Csv imported")
	})



	e.Logger.Fatal(e.Start(":1000"))
}