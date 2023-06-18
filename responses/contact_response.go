package responses

import "github.com/gofiber/fiber/v2"

type ContactResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}

//error mapping
const (
	ContactAlreadyExists string = "Contact already exists!"
	ContactNotFound string = "Contact wasn't found!"
	NoUpdatedFields string = "No Contact's fields to update was provided"
	DeletedContact string = "Contact successfully deleted!"
	PafinationParamsConvertFailed string = "Error appeared when trying to convert pagination params"
)