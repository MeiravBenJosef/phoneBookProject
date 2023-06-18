package utills

import (
	"github.com/MeiravBenJosef/phoneBookProject/models"
	"github.com/MeiravBenJosef/phoneBookProject/database"


	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson" 


	"errors"
	"strings"
	"log"
	"os"
	)

var validate = validator.New()

//loggers
var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var warnLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
var errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)


//msgForTag is a a helper function to map all kinds of struct tags to it's error message.
func msgForTag(tag string) string {
    switch tag {
    case "required":
        return "Required"
    case "excludesall":
        return "Shouldn't contain the following chars: !@#?*/$&<>"
    }
    return ""
}

//ValidateRequestBody will parse the request body to a contact struct, and will validate according to struct tags
//In case of error in the struct validations  based on tags, the function will return all errors struct's validation test
func ValidateRequestBody(contact interface{}, c *fiber.Ctx) string{
	if err := c.BodyParser(contact); err != nil {
		errorLogger.Println("Failed to parge request context into contact")
        return err.Error()
    }

    //validate required request's fields
    validationErr := validate.Struct(contact)
	if validationErr != nil {
		var ve validator.ValidationErrors
        if errors.As(validationErr, &ve) {
            out := make([]string, len(ve))
            for i, fe := range ve {
                out[i] = fe.Field() + ": " + msgForTag(fe.Tag())
            }
			errorLogger.Println("Failed with validation tags errors" + strings.Join(out, ", "))
			return strings.Join(out, ", ")
        }
		errorLogger.Println("Failed with validation errors")
		return validationErr.Error()
    }
	return "Success"
    }


//CreateContactUpdatePayload will create the currect payload for update mongo db request
//the function will build the correct payload based on which properties recieved on the request
//this function assumes at least one of the two: phone/address requires update
func CreateContactUpdatePayload(contact models.Contact, name string) bson.M{
	if contact.Phone == "notProvided"{
		//in this case, phone wasn't recieved in request, so update only address
		warnLogger.Println("Phone update isn't required for " + name + " , updating only contact's address")
		return bson.M{database.ContactAddressKey: contact.Address}
	}
	if contact.Address == "notProvided"{
		//in this case, address wasn't recieved in request, so update only phone
		warnLogger.Println("Adress update isn't required for " + name +" , updating only contact's address")
		return bson.M{database.ContactPhoneKey: contact.Phone}
	}
	logger.Println("Editing contact's phone and address for " + name)
	return bson.M{database.ContactPhoneKey: contact.Phone, database.ContactAddressKey: contact.Address}
}

