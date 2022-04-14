package service

import (
	//"context"

	//"time"

	"fmt"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
)

type Manager interface {
	CreateUser(usr *domain.User) *domain.Error
	UpdateUser(usr *domain.User) error
	ListUser(limit int32, name string) []*domain.User
	DeleteUser(id string) *domain.Error
	ViewUser(string) (*domain.User, error)
}

func NewManager(dbType string) Manager {

	/*	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

		if err != nil {
			fmt.Println("The error is", err)
		}

		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = client.Ping(ctx, readpref.Primary())

		if err != nil {
			fmt.Println("The error is", err)
			return nil
		}

		collection := client.Database("users_db").Collection("users")

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 	{Key:"_id",Value:"1234"},
		// 	{Key:"name", Value:"Shivani"},
		// 	{Key:"address", Value:"ABC Park Colony"},
		// })

		// Insert
		res, err := collection.InsertOne(ctx, bson.D{
			{Key: "id", Value: "4"},
			{Key: "name", Value: "Kuldeep"},
			{Key: "address", Value: "Kuldeep Address"},
		})
		id := res.InsertedID
		fmt.Println(id)

		// Find

		ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.D
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
			// do something with result....
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
	*/
	// // // //
	ds, err := db.NewDataStore(dbType)
	if err != nil {
		fmt.Println("The err is", err)
		return nil
	}

	return &mgr{ds: ds}

	//return &mgr{ds: make(map[int]*domain.User)}
}
