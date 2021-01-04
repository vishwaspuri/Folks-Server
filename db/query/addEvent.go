package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//AddEvent - adds new event to db.
func AddEvent(name string, description string, destination string, locationLatitude float64, locationLongitude float64, datetime time.Time, hostedBy primitive.ObjectID, inviteList []primitive.ObjectID, picturesUrls []string, admins []primitive.ObjectID, client *mongo.Client) (models.Event, error) {
	var err error
	var event models.Event
	emptyEventObject := models.Event{}

	event.ID = primitive.NewObjectID()
	event.Name = name
	event.Description = description
	event.Destination = destination
	event.LocationLatitude = locationLatitude
	event.LocationLongitude = locationLongitude
	event.Datetime = datetime
	event.CreatedOn = time.Now()
	event.HostedBy = hostedBy
	event.InviteList = inviteList
	event.Admins = admins
	event.PicturesUrls = picturesUrls

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	_, err = collection.InsertOne(ctx, event)
	if err != nil {
		return emptyEventObject, err
	}

	return event, nil
}
