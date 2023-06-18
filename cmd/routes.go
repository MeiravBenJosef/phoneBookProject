package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MeiravBenJosef/phoneBookProject/handlers"
	"github.com/MeiravBenJosef/phoneBookProject/configs"

	"go.mongodb.org/mongo-driver/mongo"

)

var contactCollection *mongo.Collection = configs.GetCollection(configs.DB, "contacts")

//This function will navigate requests to the currect function 
func setupRoutes(app *fiber.App){
	app.Post("/contact", handlers.CreateContact)
	app.Get("/getContact", handlers.GetAContact)
	app.Put("/editContact", handlers.EditAContact) 
	app.Delete("/deleteContact", handlers.DeleteAContact)
	app.Get("/contacts/:page", handlers.GetContacts)
}