package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pijalu/go.hands.two/frinsultdata/model"
)

// init will automigrate our DB
func init() {
	err := withDB(func(db *gorm.DB) error {
		db.AutoMigrate(&model.Frinsult{})
		return nil
	})
	if err != nil {
		panic(err)
	}

	i, err := GetFrinsultByID(1)
	if err != nil {
		panic(err)
	}

	if i.ID != 1 {
		initialLoad()
	}
}

// initialLoad loads some insults
func initialLoad() {
	log.Printf("Initial load")

	err := withDB(func(db *gorm.DB) error {
		for _, t := range []string{
			"con",
			"connard",
			"enculé",
			"raclure",
			"crétin des Alpages",
			"déchet de matrice",
		} {
			db.Create(&model.Frinsult{Text: t})
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

}
