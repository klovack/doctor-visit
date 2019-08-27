package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/volatiletech/null"
)

// Varchar is custom scalar for string that can be null from the database
var Varchar = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "varchar",
		Description: "VARCHAR is a string type like data type from SQL",
		Serialize: func(value interface{}) interface{} {
			val := value.(null.String)

			return val.String
		},
	},
)

// SQLInt is custom scalar for int that can be null from the database
var SQLInt = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "int",
		Description: "Int value from database can be null, this type handles that",
		Serialize: func(value interface{}) interface{} {
			val := value.(null.Int)

			if val.Valid {
				return val.Int
			}

			return nil
		},
	},
)

// Doctor is a graphql object, it is used by graphql field to create a doctor type
var Doctor = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Doctor",
		Fields: graphql.Fields{
			"doctor_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: Varchar,
			},
			"title": &graphql.Field{
				Type: Varchar,
			},
			"specialist": &graphql.Field{
				Type: Varchar,
			},
			"address": &graphql.Field{
				Type: Varchar,
			},
		},
	},
)

// Illness is a graphql object, it is used by graphql field to create an illness type
var Illness = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Illness",
		Fields: graphql.Fields{
			"illness_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: Varchar,
			},
			"risk": &graphql.Field{
				Type: SQLInt,
			},
		},
	},
)

// User is a graphql object. It is used by graphql field to create an User type
func User(resolver *Resolver) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: func() graphql.Fields {
				return graphql.Fields{
					"user_id": &graphql.Field{
						Type: graphql.Int,
					},
					"first_name": &graphql.Field{
						Type: Varchar,
					},
					"last_name": &graphql.Field{
						Type: Varchar,
					},
					"email": &graphql.Field{
						Type: graphql.String,
					},
					"birth_date": &graphql.Field{
						Type: graphql.DateTime,
					},
					"favorite_doctors": &graphql.Field{
						Type:    graphql.NewList(Doctor),
						Resolve: resolver.FindFavoritesDoctors,
					},
					"illness_history": &graphql.Field{
						Type: graphql.NewList(UserIllness(resolver)),
					},
				}
			},
		},
	)
}

func UserIllness(resolver *Resolver) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "userIllness",
			Fields: func() graphql.Fields {
				return graphql.Fields{
					"user_illness_id": &graphql.Field{
						Type: graphql.Int,
					},
					"illness": &graphql.Field{
						Type: Illness,
					},
					"from": &graphql.Field{
						Type: graphql.DateTime,
					},
					"to": &graphql.Field{
						Type: graphql.DateTime,
					},
					"notes": &graphql.Field{
						Type: graphql.String,
					},
					"doctor": &graphql.Field{
						Type: Doctor,
					},
					"user": &graphql.Field{
						Type: User(resolver),
					},
				}
			},
		},
	)
}
