package dto

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
)
