package user

type (
	User struct {
		Username    string  `form:"username" validate:"required,max=50,no_specials"`
		Password    string  `form:"password" validate:"required,min=8,max=72,contains_digit,contains_uppercase"`
		Description *string `form:"description"`
		Firstname   *string `form:"firstname" validate:"omitempty,max=25"`
		Lastname    *string `form:"lastname" validate:"omitempty,max=25"`
		Photo       *[]byte
	}
	UpdateUser struct {
		Username    *string `json:"username" validate:"max=50,no_specials"`
		Firstname   *string `json:"firstname" validate:"omitempty,max=25"`
		Lastname    *string `json:"lastname" validate:"omitempty,max=25"`
		Description *string `json:"description" validate:"max=512"`
		Photo       *[]byte `json:"photo"`
		Password    *string `json:"password" validate:"omitempty,min=8,max=72,contains_digit,contains_uppercase"`
	}

	GetFavourites struct {
		Count  *int    `json:"count"`
		Offset *int    `json:"offset"`
		Sort   *string `json:"sort"`
	}
)
