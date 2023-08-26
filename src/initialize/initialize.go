package initialize


import(
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


var (
	mongoClient *mongo.Client
	mongoDB *mongo.Collection
)

func initClient(){
	// Set client options to specify the MongoDB connection URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error

	// Connect to MongoDB
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to mongoDB: ",err)
		return
	}

	// Ping the MongoDB server to check if it's responsive
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to mongoDB: ",err)
	}

	// If the Ping operation was successful, the MongoDB server is responsive
	log.Println("Connection to MongoDB inititalized")

}

func getClient() (*mongo.Client){
	return mongoClient
}



func InitDB() {

	initClient()
	DBclient := getClient();
	mongoDB = DBclient.Database("PswdMngr").Collection("EntryDB")

	log.Println("Created EntryDB database for PswdMngr")
}

func GetDB() (*mongo.Collection){
	return mongoDB
}

