package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/klovack/doctor-visit/db"
)

// Root holds the query for the graphql
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *db.DB) *Root {

	// Create reslver for holding our database.
	// More on resolver https://graphql.org/learn/execution/#root-fields-resolvers
	resolver := Resolver{db: db}

	// Create a new Root
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"doctor":      doctorQuery(&resolver),
					"doctors":     doctorsQuery(&resolver),
					"illness":     illnessQuery(&resolver),
					"illnesses":   illnessesQuery(&resolver),
					"user":        userQuery(&resolver),
					"users":       usersQuery(&resolver),
					"userByEmail": userByEmailQuery(&resolver),
				},
			},
		),
	}

	return &root
}

func doctorQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "doctor",
		// Slice of Doctor
		Type: Doctor,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Get the list of all doctors",
		Resolve:     resolver.FindDoctor,
	}
}

func doctorsQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "doctors",
		// Slice of Doctor
		Type: graphql.NewList(Doctor),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"specialist": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"address": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "If not set, then the default limit will be used",
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: EDoctorsOrderBy,
			},
		},
		Description: "Get the list of all doctors. If limit is not set, the default limit will be used",
		Resolve:     resolver.FindDoctors,
	}
}

func illnessQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "illness",
		Type: Illness,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolver.FindIllness,
	}
}

func illnessesQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "illnesses",
		Type: graphql.NewList(Illness),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"risk": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: EIllnessOrderBy,
			},
		},
		Description: "Get the list of illnesses. If limit is not set, the default limit will be used",
		Resolve:     resolver.FindIllnesses,
	}
}

func userQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "user",
		Type: User(resolver),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolver.FindUser,
	}
}

func usersQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "users",
		Type: graphql.NewList(User(resolver)),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"birthDate": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
		},
		Resolve: resolver.FindUsers,
	}
}

func userByEmailQuery(resolver *Resolver) *graphql.Field {
	return &graphql.Field{
		Name: "userByEmail",
		Type: User(resolver),
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolver.FindUserByEmail,
	}
}
