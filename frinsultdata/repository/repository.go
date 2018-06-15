package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pijalu/go.hands.two/env"
	"github.com/pijalu/go.hands.two/frinsultdata/model"

	// Insert all gorm db driver
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// withDB execute a function with a gorm db
func withDB(f func(*gorm.DB) error) error {
	dbType := env.GetEnvWithDefault("DB_TYPE", "sqlite3")
	dbParam := env.GetEnvWithDefault("DB_PARAM", "frinsult.db")

	db, err := gorm.Open(dbType, dbParam)
	if err != nil {
		return err
	}
	defer db.Close()
	return f(db)
}

// GetFrinsultByID returns a frinsult by id
func GetFrinsultByID(ID uint) (*model.Frinsult, error) {
	log.Printf("Getting insult for id %d", ID)
	var frinsult model.Frinsult
	err := withDB(func(db *gorm.DB) error {
		db.First(&frinsult, ID)
		return nil
	})
	return &frinsult, err
}

// DeleteFrinsultByID deletes a frinsult
func DeleteFrinsultByID(ID uint) error {
	log.Printf("Deleting insult for id %d", ID)
	frinsult := model.Frinsult{Model: gorm.Model{ID: ID}}

	return withDB(func(db *gorm.DB) error {
		db.Delete(&frinsult)
		return nil
	})
}

// UpdateFrinsult updates a frinsult
func UpdateFrinsult(f *model.Frinsult) error {
	log.Printf("Updating insult for id %d", f.ID)

	return withDB(func(db *gorm.DB) error {
		db.Save(f)
		return nil
	})
}

// InsertFrinsult insert a frinsult
func InsertFrinsult(f *model.Frinsult) (*model.Frinsult, error) {
	log.Printf("Updating insult for id %d", f.ID)

	err := withDB(func(db *gorm.DB) error {
		db.Create(f)
		return nil
	})

	return f, err
}

// GetFrinsults returns a list of frinsult
func GetFrinsults() ([]model.Frinsult, error) {
	log.Printf("Listing insults")
	f := make([]model.Frinsult, 0, 200)

	return f, withDB(func(db *gorm.DB) error {
		db.Find(&f)
		return nil
	})

}

// VoteForFrinsult update the score by the vote value
func VoteForFrinsult(ID uint, vote int) error {
	log.Printf("Voting for insult id %d - value: %d", ID, vote)

	return withDB(func(db *gorm.DB) error {
		db.Model(&model.Frinsult{Model: gorm.Model{ID: ID}}).Update(
			"score",
			gorm.Expr("score + ?", vote))
		return nil
	})
}
