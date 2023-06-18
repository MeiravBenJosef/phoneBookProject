package models

type Contact struct {
	FirstName string `json:"name,omitempty" validate:"required,excludesall=!@#?*/$&<>" bson:"first_name, omitempty"`
	LastName string `json:"lastName,omitempty" validate:"required,excludesall=!@#?*/$&<>" bson:"last_name, omitempty"`
	Phone string `json:"phone,omitempty" validate:"excludesall=!@#?*/$&<>" bson:"phone_number, omitempty"`
	Address string `json:"address,omitempty" validate:"excludesall=!@#?*/$&<>" bson:"address, omitempty"`
	FullName string `json:"-" bson:"full_name, omitempty"`
}

type SearchTerm struct {
	SearchTerm string `json:"searchTerm,omitempty" validate:"required,excludesall=!@#?*/$&<>" bson:"searchTerm, omitempty"`
}