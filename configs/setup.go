package configs

import (
    "context"
    "log"
    "time"
	"os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


//var DB Dbinstance

//ConnectDB will create a mongo db client based on string connection
func ConnectDB() *mongo.Client{
    client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGOURI")))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }
    log.Output(1, "Connected to MongoDB")
	return client
}



//Client instance
var DB *mongo.Client = ConnectDB()


//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database(os.Getenv("MONGOCLUSTERNAME")).Collection(collectionName)
    return collection
}