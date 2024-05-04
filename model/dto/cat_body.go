package dto

import "time"

type (
	RequestCreateCat struct {
		Name  string   `json:"name"`
		Race  string   `json:"race"`
		Sex   string   `json:"sex"`
		Age   int      `json:"ageInMonth"`
		Desc  string   `json:"description"`
		Image []string `json:"imageUrls"`
	}

	SuccessCreateCat struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	}

	CatGet struct {
		ID         string `form:"id"`
		Limit      int    `form:"limit,default=5"`
		Offset     int    `form:"offset,default=0"`
		Race       string `form:"race"`
		Sex        string `form:"sex"`
		HasMatched bool   `form:"hasMatched"`
		AgeInMonth string `form:"ageInMonth"`
		Owned      bool   `form:"owned"`
		Search     string `form:"search"`
	}

	// ResponseCat represents the response structure for a cat
	ResponseCat struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Race        string    `json:"race"`
		Sex         string    `json:"sex"`
		AgeInMonth  int       `json:"ageInMonth"`
		ImageUrls   []string  `json:"imageUrls"`
		Description string    `json:"description"`
		HasMatched  bool      `json:"hasMatched"`
		CreatedAt   time.Time `json:"createdAt"`
	}
)
