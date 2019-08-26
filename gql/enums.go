package gql

import "github.com/graphql-go/graphql"

// EIllnessOrderBy enum to be used in `illness.OrderBy`
var EIllnessOrderBy = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "IllnessOrderBy",
		Values: graphql.EnumValueConfigMap{
			"illness_id": &graphql.EnumValueConfig{
				Value: 0,
			},
			"name": &graphql.EnumValueConfig{
				Value: 1,
			},
			"risk": &graphql.EnumValueConfig{
				Value: 2,
			},
		},
	},
)

// EDoctorsOrderBy enum to be used in `doctors.OrderBy`
var EDoctorsOrderBy = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "DoctorOrderBy",
		Values: graphql.EnumValueConfigMap{
			"doctor_id": &graphql.EnumValueConfig{
				Value: 0,
			},
			"name": &graphql.EnumValueConfig{
				Value: 1,
			},
			"title": &graphql.EnumValueConfig{
				Value: 2,
			},
			"specialist": &graphql.EnumValueConfig{
				Value: 3,
			},
		},
	},
)

// EUsersOrderBy to be sued in `users.OrderBy`
var EUsersOrderBy = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "UserOrderBy",
		Values: graphql.EnumValueConfigMap{
			"user_id": &graphql.EnumValueConfig{
				Value: 0,
			},
			"first_name": &graphql.EnumValueConfig{
				Value: 1,
			},
			"last_name": &graphql.EnumValueConfig{
				Value: 2,
			},
			"email": &graphql.EnumValueConfig{
				Value: 3,
			},
			"birthdate": &graphql.EnumValueConfig{
				Value: 4,
			},
		},
	},
)
