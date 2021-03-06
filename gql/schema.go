package gql

import (
	"github.com/olivere/elastic/v7"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/wefolks/backend/db"
)

var mongo *db.MongoDB
var elasti *elastic.Client
// InitSchema - defines complete graphql schema
func InitSchema(d *db.MongoDB, ec *elastic.Client) graphql.Schema {
	mongo = d
	elasti = ec
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllEvents": &graphql.Field{
					Type:    graphql.NewList(EventType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllEvents,
				},
				"getAllUsers": &graphql.Field{
					Type:    graphql.NewList(UserType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllUsers,
				},
				"getAllSquads": &graphql.Field{
					Type:    graphql.NewList(SquadType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllSquads,
				},
				"getEventByID": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: getEvent,
				},
				"getUserByID": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: getUser,
				},
				"getSquadByID": &graphql.Field{
					Type: SquadType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: getSquad,
				},
				"myProfile": &graphql.Field{
					Type:    UserType,
					Args:    graphql.FieldConfigArgument{},
					Resolve: myProfile,
				},
				"getNearByEvents": &graphql.Field{
					Type: graphql.NewList(EventType),
					Args: graphql.FieldConfigArgument{
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"radius": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
					},
					Resolve: getNearByEvents,
				},
				"getNearByEventsWithImages": &graphql.Field{
					Type: graphql.NewList(EventType),
					Args: graphql.FieldConfigArgument{
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"radius": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
					},
					Resolve: getNearByEventsWithImages,
				},
				"getPastEvents": &graphql.Field{
					Type: graphql.NewList(EventType),
					Args: graphql.FieldConfigArgument{},
					Resolve: getPastEvents,
				},
				"getUpcommingEvents": &graphql.Field{
					Type: graphql.NewList(EventType),
					Args: graphql.FieldConfigArgument{},
					Resolve: getUpcommingEvents,
				},
				"myNotifications": &graphql.Field{
					Type: graphql.NewList(NotificationType),
					Args: graphql.FieldConfigArgument{},
					Resolve: myNotifications,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addEvent": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"destination": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"datetime": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"inviteList": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"picturesUrls": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: addEvent,
				},
				"addSquad": &graphql.Field{
					Type: SquadType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"groupImages": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: addSquad,
				},
				"followUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: followUser,
				},
				"acceptRequest": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: acceptRequest,
				},
				"declineRequest": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: declineRequest,
				},
				"requestEvent": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: requestEvent,
				},
				"acceptParticipants": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"userID": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: acceptParticipants,
				},
				"declineParticipants": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"userID": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: declineParticipants,
				},
				"inviteParticipants": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"userID": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: inviteParticipants,
				},
				"assignAdmin": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"admins": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: assignAdmin,
				},
				"acceptInvite": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: acceptInvite,
				},
				"declineInvite": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"eventID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: declineInvite,
				},
				"updateEvent": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"destination": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"datetime": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"picturesUrls": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: updateEvent,
				},
				"changePassword": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"oldPassword": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"newPassword": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: changePassword,
				},
				"updateUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"phoneNo": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"firstName": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"lastName": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"isPublic": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
					},
					Resolve: updateUser,
				},
			},
		}),
		Types: []graphql.Type{ID, UserType, EventType, SquadType},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema
}
