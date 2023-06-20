package models

type Contact struct {
	FirstName string `json:"name,omitempty" validate:"required,excludesall=!@#?*/$&<>,min=1,max=15" bson:"first_name, omitempty"`
	LastName string `json:"lastName,omitempty" validate:"required,excludesall=!@#?*/$&<>,min=1,max=20" bson:"last_name, omitempty"`
	Phone string `json:"phone,omitempty" validate:"excludesall=!@#?*/$&<>,max=30" bson:"phone_number, omitempty"`
	Address string `json:"address,omitempty" validate:"excludesall=!@#?*/$&<>,max=30" bson:"address, omitempty"`
	FullName string `json:"-" bson:"full_name, omitempty"`
}

type SearchTerm struct {
	SearchTerm string `json:"searchTerm,omitempty" validate:"required,excludesall=!@#?*/$&<>" bson:"searchTerm, omitempty"`
}