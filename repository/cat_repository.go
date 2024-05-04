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
	GetCats(data dto.CatGet) ([]dto.ResponseCat, error)
}

func (c *catRepository) GetCats(data dto.CatGet) ([]dto.ResponseCat, error) {

	var cats []dto.ResponseCat

	// Build SQL query
	query := "SELECT * FROM cats WHERE 1=1"
	var args []interface{}

	// Add conditions based on the provided parameters
	if data.ID != "" {
		query += " AND id = ?"
		args = append(args, data.ID)
	}
	if data.Race != "" {
		query += " AND race = ?"
		args = append(args, data.Race)
	}
	if data.Sex != "" {
		query += " AND sex = ?"
		args = append(args, data.Sex)
	}
	// Add other conditions...

	// Execute raw SQL query
	result := c.db.Raw(query, args...).Scan(&cats)
	if result.Error != nil {
		return nil, result.Error
	}

	return cats, nil
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
