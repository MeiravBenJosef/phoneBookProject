**Phone book project**
----
**Running docker project locally instructions**

  * Clone the project to a local folder on your machine.
  * Inside .env file, for MONGOURI env variable, put the mongo db connection string provided to you.
  * Make sure to have docker desktop installed and running. Can be downloaded here: https://www.docker.com/products/docker-desktop/
  * Open a new terminal inside the project and run the following: docker compose up
  * A new docker container should be created and app should run at http://127.0.0.1:3000 
  * In order to send requests to the app, please use your favourite API client, i used Postman desktop: https://www.postman.com/downloads/
  
**Running project's tests**

  * Open a new terminal inside the project.
  * run the following: docker compose run --service-ports web bash
  * Go inside cmd folder: cd cmd/
  * Run the following: go test -v


----
**API documentation**
___
**Add contact**
___
  Adds a new contact to the phone book.

* **URL**

  /contact"

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Body**<br />
JSON<br />
`{
   "name": "first name",
   "lastName":"last name",
   "phone": "phone",
   "address": "address"
}`<br />
**Required:**<br />
  `"name":"string"` AND `"lastName":"string"`
**Required:**<br />
  * First and last name are required to create new contact
  * The following characters are not allowed for any property: !@#?*/$&<>

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** `{
    "status": 201,
    "message": "success",
    "data": {
        "InsertedID": "ID"
    }
}`
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{
    "status": 400,
    "message": "error",
    "data": {
        "data": "error message"
    }
}`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
    "status": 500,
    "message": "error",
    "data": {
        "data": "error message"
    }
}`


