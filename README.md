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

  /users/:id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : 12, name : "Michael Bloom" }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
