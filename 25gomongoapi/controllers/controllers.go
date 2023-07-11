package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://ozair123:ozair123@cluster0.mcvc4.mongodb.net/netflix?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//IMP
var collection *mongo.Collection

func init(){
	//client Option
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection Instance is ready!")
}

//mongodb Connection
//insert 1st record

func insertFirstMovie(movie models.Movie){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %d document(s)\n", inserted.InsertedID)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"$set": bson.M{"watched": true}}}

	deleteCount, err := collection.UpdateOne(context.Background(), filter, update)

	if err!= nil {
        log.Fatal(err)
    }

	fmt.Println("modified Count: ", deleteCount.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
    filter := bson.M{"_id": id}

    deleteCount, err := collection.DeleteOne(context.Background(), filter)

    if err!= nil {
        log.Fatal(err)
    }

    fmt.Println("deleted Count: ", deleteCount)
}

//delete all movies
func deleteAllMovies() int64{
    deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

    if err!= nil {
        log.Fatal(err)
    }

    fmt.Println("Number of deleted Movie: ", deleteResult)
	return deleteResult.deletedCount
}

//find all movies
func findAllMovies() []primitive.M {
    var movies []primitive.M

    cursor, err := collection.Find(context.Background(), bson.D{{}})

    if err!= nil {
        log.Fatal(err)
    }

    for cursor.Next(context.Background()) {
        var movie bson.M
        err := cursor.Decode(&movie)

        if err!= nil {
            log.Fatal(err)
        }

        movies = append(movies, movie)
    }
	defer cursor.Close()
    return movies
}


//actual Controllers

func getAllMoviesController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := findAllMovies()

	return json.NewEncoder(w).Encode(allMovies)
}