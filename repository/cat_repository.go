package repository

import (
	"7Zero4/model/dto"
	"strings"

	"gorm.io/gorm"
)

type catRepository struct {
	db *gorm.DB
}

type CatRepository interface {
	InsertCat(data dto.RequestCreateCat) error
}

func (c *catRepository) InsertCat(data dto.RequestCreateCat) error {
	result := c.db.Exec("INSERT INTO cat(cat_name,cat_race,cat_sex,cat_age,description,image) VALUES(?,?,?,?,?,?)", data.Name, data.Race, data.Sex, data.Age, data.Desc, strings.Join(data.Image, "||"))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewCatRepository(db *gorm.DB) CatRepository {
	repo := new(catRepository)
	repo.db = db
	return repo
}
