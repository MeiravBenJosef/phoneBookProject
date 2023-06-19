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

  /contact

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

___
**Search contact**
___
  Search a contact in phone book, by full name.

* **URL**

  /searchContact

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Body**<br />
JSON<br />
`{
   "name": "first name",
   "lastName":"last name"
}`<br />
**Required:**<br />
  `"name":"string"` AND `"lastName":"string"`
**Required:**<br />
  * First and last name are required to create new contact
  * The following characters are not allowed for any property: !@#?*/$&<>

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
    "status": 200,
    "message": "success",
    "data": {
        "name": "first name",
        "lastName":"last name",
        "phone":"phone",
        "address":"address"
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

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{
    "status": 404,
    "message": "error",
    "data": {
        "data": "Contact wasn't found!"
    }
}`

___
**Edit contact**
___
  Edit phone/address of a contact in phone book, by it's full name.

* **URL**

  /editContact

* **Method:**

  `PUT`
  
*  **URL Params**

   None

* **Body**<br />
JSON<br />
`{
   "name": "first name",
   "lastName":"last name",
   "phone":"phone",
   "address":"address"
}`<br />
**Required:**<br />
  `"name":"string"` AND `"lastName":"string"`
  `"phone":"string"` OR `"address":"string"`
**Required:**<br />
  * First and last name are required to edit new contact
  * The following characters are not allowed for any property: !@#?*/$&<>
  * One of the two is required to make update: phone, address

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
    "status": 200,
    "message": "success",
    "data": {
        "name": "first name",
        "lastName":"last name",
        "phone":"phone",
        "address":"address"
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

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{
    "status": 404,
    "message": "error",
    "data": {
        "data": "Contact wasn't found!"
    }
}`

___
**Delete contact**
___
  Delete a contact in phone book, by it's full name.

* **URL**

  /deleteContact

* **Method:**

  `DELETE`
  
*  **URL Params**

   None

* **Body**<br />
JSON<br />
`{
   "name": "first name",
   "lastName":"last name"
}`<br />
**Required:**<br />
  `"name":"string"` AND `"lastName":"string"`
**Required:**<br />
  * First and last name are required to delete new contact
  * The following characters are not allowed for any property: !@#?*/$&<>

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
    "status": 200,
    "message": "success",
    "data": {
        "data": "Contact successfully deleted!"
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

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{
    "status": 404,
    "message": "error",
    "data": {
        "data": "Contact wasn't found!"
    }
}`

___
**Get contacts- with pagination feature and limit of 10**
___
  Get all contacts in phone book, from a certain page, which is an input.

* **URL**

  /contacts/:page

* **Method:**

  `GET`
  
*  **URL Params**

   None


* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
    "status": 200,
    "message": "success",
    "data": {[{
        "name": "first name",
        "lastName":"last name",
        "phone":"phone",
        "address":"address"
    }]}
}`
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
    "status": 500,
    "message": "error",
    "data": {
        "data": "error message"
    }
}`



