package usecase

import (
	"7Zero4/model/dto"
	"7Zero4/repository"
	"7Zero4/utils"
)

type CatUseCase interface {
	CreateCat(data dto.RequestCreateCat) error
}

type catUsecase struct {
	catRepo       repository.CatRepository
	raceWhitelist map[string]struct{}
	sexWhitelist  map[string]struct{}
}

// func (c *catUsecase) GetCat(data dto.CatGet)

func (c *catUsecase) validateCatRequest(data dto.RequestCreateCat) bool {

	if data.Name == "" || len(data.Name) > 30 {
		return false
	}

	_, raceOk := c.raceWhitelist[data.Race]
	if !raceOk {
		return false
	}

	_, SexOk := c.sexWhitelist[data.Sex]
	if !SexOk {
		return false
	}

	if data.Age < 1 || data.Age > 120082 {
		return false
	}

	if data.Name == "" || len(data.Name) > 200 {
		return false
	}

	if len(data.Image) == 0 {
		return false
	}

	return true
}

func (c *catUsecase) CreateCat(data dto.RequestCreateCat) error {
	validation := c.validateCatRequest(data)
	if !validation {
		return utils.ReqBodyNotValidError()
	}

	err := c.catRepo.InsertCat(data)
	if err != nil {
		return err
	}

	return nil
}

func NewCatUsecase(catRepo repository.CatRepository) CatUseCase {
	usecase := new(catUsecase)
	usecase.catRepo = catRepo
	usecase.sexWhitelist = map[string]struct{}{
		"male":   {},
		"female": {},
	}
	usecase.raceWhitelist = map[string]struct{}{
		"Persian":           {},
		"Maine Coon":        {},
		"Siamese":           {},
		"Ragdoll":           {},
		"Bengal":            {},
		"Sphynx":            {},
		"British Shorthair": {},
		"Abyssinian":        {},
		"Scottish Fold":     {},
		"Birman":            {},
	}

	return usecase
}
