package main

import (
	"net/http"
	"testing"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"io"
	
	
	"github.com/stretchr/testify/assert"
	"github.com/gofiber/fiber/v2"
	"github.com/MeiravBenJosef/phoneBookProject/models"
	"github.com/MeiravBenJosef/phoneBookProject/responses"

)

var app *fiber.App


//Test to describes different test's cases properties
type Test struct {
	description string
	expectedCode  int
	expectedBody  responses.ContactResponse
	checkBody bool
}

//names of the test contacts 
const (
	validNewContactFirstName string = "mia"
	validNewContactLastName string = "poleg"
	NoPhoneNewContactFirstName string = "miana"
	NoPhoneNewContactLastName string = "poleg"
	NotFoundContactFirstName string = "Notfoundfirst"
	NotFoundContactLastName string = "NotfoundLast"
	NotValidFirstName string = "Not Valid <?"

)

//contact request inputs, used for create test
var contactTestInputs = []models.Contact{
	{
		FirstName: validNewContactFirstName,
		LastName: validNewContactLastName,
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: validNewContactFirstName,
		LastName: validNewContactLastName,
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: "",
		LastName: "Goldfarb",
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: "mri",
		LastName: "",
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: NoPhoneNewContactFirstName,
		LastName: NoPhoneNewContactLastName,
	},
	{
		FirstName: NotValidFirstName,
		LastName: NoPhoneNewContactLastName,
	},
	
}

//conract request inputs, used for all other operations
var searchContactTestInputs = []models.Contact{
	{
		FirstName: validNewContactFirstName,
		LastName: validNewContactLastName,
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: NotFoundContactFirstName,
		LastName: NotFoundContactLastName,
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: "",
		LastName: "Goldfarb",
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: "mri",
		LastName: "",
		Phone: "0542022898",
		Address: "marcus 13, jerusalem",
	},
	{
		FirstName: NoPhoneNewContactFirstName,
		LastName: NoPhoneNewContactLastName,
	},
	
	
}



//main function to setup the app and run all other test
func TestMain(m *testing.M) {
	// Setup the app as it is done in the main function
	app = fiber.New()
	setupRoutes(app)

	code:= m.Run()

	os.Exit(code)
}


//createResultContactResponse will return ContactResponse that contains contact details result
func createResultContactResponse(inputs []models.Contact, inputidx int, status int, msg string) responses.ContactResponse{
	return responses.ContactResponse{
		Status: status,
		Message: msg,
		Data: &fiber.Map{"data": inputs[inputidx]},
	}
}

//createResultContactResponse will return ContactResponse that contains error
func createErrorContactResponse(datamsg string, status int, msg string) responses.ContactResponse{
	return responses.ContactResponse{
		Status: status,
		Message: msg,
		Data: &fiber.Map{"data": datamsg},
	}
}

//helperRunner will execute all test cases of all tests, and run validations on the result
func helperRunner(endpoint string, requesttype string, tests []Test, inputs []models.Contact, t *testing.T){

	for i, test := range tests {
		// Create a new http request with the route
		// from the test case
		jsonBytes, err :=json.Marshal(inputs[i])
		if err != nil{
			log.Println("error")
		}
	
 		bodyReader := bytes.NewReader(jsonBytes)
		req, _ := http.NewRequest(
			requesttype,
			endpoint,
			bodyReader,
		)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		//in case the test requires it, verify response body
		if test.checkBody {
			body, err := io.ReadAll(res.Body)

		    assert.Nilf(t, err, test.description)

			expectedBodyJson, parseErr := json.Marshal(test.expectedBody)

			//shouldn't fail
			assert.Nilf(t, parseErr, test.description)

			// Verify, that the reponse body equals the expected body
		    assert.Equalf(t, string(expectedBodyJson), string(body), test.description)

		}

		
	}
}

//helperRequestRunnerWithoutBody is the same as previous, only for function which doesn't send anything in the body
func helperRequestRunnerWithoutBody(endpoint string, requesttype string, tests []Test, t *testing.T){

	for _, test := range tests {

		req, _ := http.NewRequest(
			requesttype,
			endpoint,
			nil,
		)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		res, _ := app.Test(req, -1)

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	}
}


//TestCreateContact will build all relevant test cases for contact creation flow and run the tests.
func TestCreateContact(t *testing.T) {
	var createContacttests =[]Test{
		{
			description: "Create contact- 201",
			expectedCode: 201,
			checkBody: false,
		},
		{
			description: "Create existing contact- 400",
			expectedCode: 400,
			expectedBody: createErrorContactResponse(responses.ContactAlreadyExists, 400, "error"),
			checkBody: true,
		},
		{
			description: "Create contact- missing first name property",
			expectedCode: 400,
			checkBody: false,

		},
		{
			description: "Create contact- missing last name property",
			expectedCode: 400,
			checkBody: false,

		},
		{
			description:  "Create contact- missing phone and address",
			expectedCode: 201,
			checkBody: false,

		},
		{
			description: "Create contact- invalid input- contains forbidden chars",
			expectedCode: 400,
			checkBody: false,

		},
	}
	


	helperRunner("/contact", "POST", createContacttests, contactTestInputs, t)
	
}


//TestSearchContact will build all relevant test cases for contact searching flow and run the tests.
func TestSearchContact(t *testing.T) {
	
	var searchContactTests =[]Test{
		{
			description: "Search contact- 201",
			expectedCode: 200,
			expectedBody: createResultContactResponse(searchContactTestInputs, 0, 200, "success"),
			checkBody: true,
		},
		{
			description: "Search not existed contact- 404",
			expectedCode: 404,
			expectedBody: createErrorContactResponse(responses.ContactNotFound, 404, "error"),
			checkBody: true,

		},
		{
			description: "Search contact- missing first name property",
			expectedCode: 400,
			checkBody: false,

		},
		{
			description: "Search contact- missing last name property",
			expectedCode: 400,
			checkBody: false,

		},
		{
			description:  "Search contact- missing phone and address",
			expectedCode: 200,
			expectedBody: createResultContactResponse(searchContactTestInputs, 4, 200, "success"),
			checkBody: true,

		},
	}


	helperRunner("/getContact", "GET", searchContactTests, searchContactTestInputs, t)

	
}


//TestEditContact will build all relevant test cases for contact editing flow and run the tests.
func TestEditContact(t *testing.T) {

	var editContactTests =[]Test{
		{
			description: "Edit contact- 200",
			expectedCode: 200,
			expectedBody: createResultContactResponse(searchContactTestInputs, 0, 200, "success"),
			checkBody: true,

		},
		{
			description: "Edit not existed contact- 404",
			expectedCode: 404,
			expectedBody: createErrorContactResponse(responses.ContactNotFound, 404, "error"),
			checkBody: true,

		},
		{
			description: "Edit contact- missing first name property",
			expectedCode: 400,
			checkBody: false,
		},
		{
			description: "Edit contact- missing last name property",
			expectedCode: 400,
			checkBody: false,
		},
		{
			description:  "Edit contact- missing phone and address",
			expectedCode: 400,
			expectedBody: createErrorContactResponse(responses.NoUpdatedFields, 400, "error"),
			checkBody: true,
		},
	}


	helperRunner("/editContact", "PUT", editContactTests, searchContactTestInputs, t)

	
}


//TestDeleteContact will build all relevant test cases for contact deleting flow and run the tests.
func TestDeleteContact(t *testing.T) {
	
	var deleteContactTests =[]Test{
		{
			description: "Delete contact- 200",
			expectedCode: 200,
			checkBody: false,
		},
		{
			description: "Delete not existed contact- 404",
			expectedCode: 404,
			expectedBody: createErrorContactResponse(responses.ContactNotFound, 404, "error"),
			checkBody: true,
		},
		{
			description: "Delete contact- missing first name property",
			expectedCode: 400,
			checkBody: false,
		},
		{
			description: "Delete contact- missing last name property",
			expectedCode: 400,
			checkBody: false,
		},
		{
			description:  "Delete contact- missing phone and address",
			expectedCode: 200,
			checkBody: false,
		},
	}

	helperRunner("/deleteContact", "DELETE", deleteContactTests, searchContactTestInputs, t)
	
}

//TestGetContacts will build all relevant test cases for get contacts flow and run the tests.
func TestGetContacts(t *testing.T) {
	
	var getContactTests =[]Test{
		{
			description: "Get contacts- 200",
			expectedCode: 200,
			checkBody: false,
		},
	}

	helperRequestRunnerWithoutBody("/contacts/1", "GET", getContactTests, t)

}