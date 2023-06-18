package handlers

import (
"context"
"net/http"
"time"
"strconv"
"log"
"os"
"github.com/MeiravBenJosef/phoneBookProject/configs"
"github.com/MeiravBenJosef/phoneBookProject/models"
"github.com/MeiravBenJosef/phoneBookProject/database"
"github.com/MeiravBenJosef/phoneBookProject/responses"
"github.com/MeiravBenJosef/phoneBookProject/utills"

"go.mongodb.org/mongo-driver/mongo"
"github.com/gofiber/fiber/v2"
"go.mongodb.org/mongo-driver/bson" 

)

//connect to the contacts collection on mongo db client
//var dbClient *mongo.Client = configs.DB.Db
var contactCollection *mongo.Collection = configs.GetCollection(configs.DB, "contacts")

//loggers
var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

//CreateContact will create a new contact based on it's full name.
//The function validates that the contact doesn't exist, and if so, creates it.
func CreateContact(c *fiber.Ctx) error {
	//set a timeout of 10 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var contact models.Contact
    defer cancel()
    //validate the request body
	logger.Println("Validating the request body properties for create contact")
	requestValidationRes:=utills.ValidateRequestBody(&contact, c)
	if requestValidationRes != "Success"{
		return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": requestValidationRes}})
	}

	//validate contact with the same name doesn't exist already
	var fullName = utills.BuildFullName(contact.FirstName, contact.LastName)
	logger.Println("Validating "+ fullName + " contact doesn't already exists in the phone book.")
	results, err := contactCollection.Find(ctx, bson.M{database.ContactFullNameKey: fullName})
	hasResults := results.Next(ctx)
	if hasResults{
		errorLogger.Println(1, "Found "+ fullName + " in the phone book, BadRequest.")
		return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": responses.ContactAlreadyExists}})
	}

	//in case contact doesn't exists, create it.
	logger.Println(fullName + " doesn't exist in phone book. Creating it.")

    newContact := models.Contact{
        FirstName:     contact.FirstName,
        LastName: contact.LastName,
        Phone:    contact.Phone,
		Address:    contact.Address,
		FullName: fullName,
	}

    result, err := contactCollection.InsertOne(ctx, newContact)
    if err != nil {
		errorLogger.Println("An error occurred when trying to insert "+ fullName + " to collection")
        return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	logger.Println("Successfully created new contact with the name "+ fullName + " phone: "+ contact.Phone + " and address " + contact.Address)
    return c.Status(http.StatusCreated).JSON(responses.ContactResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

//GetAContact searches for a contact by it's full name 
//maybe add here specific not found error
func GetAContact(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var contact models.Contact
	defer cancel()

	 //validate the request body
	 logger.Println("Validating the request body properties for search contact.")
	 requestValidationRes:=utills.ValidateRequestBody(&contact, c)
	 if requestValidationRes != "Success"{
		 return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": requestValidationRes}})
	 }

	 //create the full name 
	fullName := utills.BuildFullName(contact.FirstName, contact.LastName)
	logger.Println("Searching " + fullName + " in phone book")
    var foundContact models.Contact
    err := contactCollection.FindOne(ctx, bson.M{database.ContactFullNameKey: fullName}).Decode(&foundContact)
    if err != nil {
		errorLogger.Println("Error when searching " + fullName + " in collection")
        return c.Status(http.StatusNotFound).JSON(responses.ContactResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": responses.ContactNotFound}})
    }

	logger.Println("Successfully found " + fullName + " in phone book with phone: "+ foundContact.Phone + " and address: " + foundContact.Address)
    return c.Status(http.StatusOK).JSON(responses.ContactResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": foundContact}})
}


//EditAContact will edit one of the following contact's properties: phone, address.
func EditAContact(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    contact:= models.Contact{
		Phone: "notProvided",
		Address: "notProvided",
	}
    defer cancel()

     //validate the request body
	logger.Println("Validating the request body properties for edit contact")
	requestValidationRes:=utills.ValidateRequestBody(&contact, c)
	if requestValidationRes != "Success"{
		return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": requestValidationRes}})
	}

	fullName := utills.BuildFullName(contact.FirstName, contact.LastName)

	if contact.Phone == "notProvided" && contact.Address == "notProvided"{
		errorLogger.Println("No edit fields were provided, not editing " + fullName)
		return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": responses.NoUpdatedFields}})
	}

	//create the update payload based on the body request fields
    update := utills.CreateContactUpdatePayload(contact, fullName)

	//update contact in database
    result, err := contactCollection.UpdateOne(ctx, bson.M{database.ContactFullNameKey: fullName}, bson.M{"$set": update})

    if err != nil {
		errorLogger.Println("Error when executing update to colletion for " + fullName)
        return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
    //get updated user details
    var updatedContact models.Contact
    if result.MatchedCount == 1 {
		//in case we updated the contact successfully, return it as result
        err := contactCollection.FindOne(ctx, bson.M{database.ContactFullNameKey: fullName}).Decode(&updatedContact)

        if err != nil {
			// error in fetching the updated contact
			errorLogger.Println("Error when fetching the updated contact " + fullName)
            return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
    }else if result.MatchedCount == 0 {
		//this case respresent the case contact to update wasn't found in database
		errorLogger.Println("Contact" + fullName + " wasn't found, can't update it.")
		return c.Status(http.StatusNotFound).JSON(responses.ContactResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": responses.ContactNotFound}})
	}
	
	logger.Println("Succssfully edited " + fullName)
    return c.Status(http.StatusOK).JSON(responses.ContactResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedContact}})
}

//DeleteAContact will delete the contact by it's full name
func DeleteAContact(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var contact models.Contact
    defer cancel()
	
	 //validate the request body
	 logger.Println("Validating the request body properties for delete contact")
	 requestValidationRes:=utills.ValidateRequestBody(&contact, c)
	 if requestValidationRes != "Success"{
		//validation failed
		 return c.Status(http.StatusBadRequest).JSON(responses.ContactResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": requestValidationRes}})
	 }

	 //build full name
	fullName := utills.BuildFullName(contact.FirstName, contact.LastName)
	
	logger.Println("Deleting " + fullName + " from contacs collection")
    result, err := contactCollection.DeleteOne(ctx, bson.M{database.ContactFullNameKey: fullName})
    if err != nil {
		errorLogger.Println("Error when Deleting " + fullName + " from contacs collection")
        return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    if result.DeletedCount < 1 {
		errorLogger.Println("Contact " + fullName + " wasn't found in collection!")
        return c.Status(http.StatusNotFound).JSON(
            responses.ContactResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": responses.ContactNotFound}},
        )
    }

	logger.Println("Contact " + fullName + " successfully deleted!")
    return c.Status(http.StatusOK).JSON(
        responses.ContactResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": responses.DeletedContact}},
    )
}


//GetContacts will return all contacts, with a pagination feature. 
//Page of pagination feature is provided by the user
func GetContacts(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert page and limit to integers
	page, pageErr:= strconv.Atoi(c.Params("page"))
	limit, limitErr:= strconv.Atoi(os.Getenv("PAGINATIONLIMIT"))
	if limitErr != nil || pageErr != nil{
		errorLogger.Println("Error when converting pagination params to integers")
		return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": responses.PafinationParamsConvertFailed}})
	}
	//create a pagination object with page and limit
	mongoPaginate := utills.NewMongoPaginate(limit, page)

    var contacts []models.Contact

	logger.Println("Executing get contacts request to the database contacs collection")
    results, err := contactCollection.Find(ctx, bson.M{}, mongoPaginate.GetPaginatedOpts())

    if err != nil {
		errorLogger.Println("Error when executing get contacts request to the database contacs collection")
        return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    //reading from the database
    defer results.Close(ctx)
    for results.Next(ctx) {
		var singleUser models.Contact
		if err = results.Decode(&singleUser); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.ContactResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
  
		contacts = append(contacts, singleUser)
	 }
  
	 logger.Println("Successfully returned contacts!")
	 return c.Status(http.StatusOK).JSON(
        responses.ContactResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": contacts}},
    )
}