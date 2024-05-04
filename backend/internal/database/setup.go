package database

import (
	"context"
	"log"

	"github.com/ultimateshadsform/pi-dashboard/internal/helpers"
	"github.com/ultimateshadsform/pi-dashboard/internal/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckSetup() {
	// Check if setup is complete
	// If not, create setup code and print it

	var dbsettings structs.DBSettings

	if err := db.Collection("settings").FindOne(context.Background(), bson.D{}).Decode(&dbsettings); err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	if dbsettings.IsSetupComplete {
		return
	}

	if dbsettings.SetupCode == "" {
		setupCode, err := helpers.GenerateRandomString(6)
		if err != nil {
			panic(err)
		}

		dbsettings.SetupCode = setupCode
		dbsettings.IsSetupComplete = false

		if _, err := db.Collection("settings").InsertOne(context.Background(), dbsettings); err != nil {
			panic(err)
		}

		log.Println("Setup code:", setupCode)
	} else {
		log.Println("Setup code:", dbsettings.SetupCode)
	}

}
