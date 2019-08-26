package gql

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/klovack/doctor-visit/db"
)

// Resolver struct holds a connection to the database,
// more on [resolver](https://graphql.org/learn/execution/#root-fields-resolvers)
type Resolver struct {
	db *db.DB
}

// DoctorResolver resolves the doctor query and manage to respond with single `doctor`
func (r *Resolver) DoctorResolver(p graphql.ResolveParams) (interface{}, error) {
	id, idOK := p.Args["id"].(int)

	if idOK {
		doctors, err := r.db.GetDoctorByID(id)
		if err != nil {
			return nil, err
		}

		return doctors, nil
	}

	return nil, fmt.Errorf("Doctor ID has to be specified")
}

// DoctorsResolver resolves the doctors query and manage call to the `GetDoctorsByAttribute` and respond with list of `doctors`
func (r *Resolver) DoctorsResolver(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	specialist, specialistOK := p.Args["specialist"].(string)
	title, titleOK := p.Args["title"].(string)
	address, addressOK := p.Args["address"].(string)

	// Here, the error is not relevant, if the limit or offset is not specified
	// it will be set to the default value
	limit, _ := p.Args["limit"].(int)
	offset, _ := p.Args["offset"].(int)

	// Serialize the order
	orderBy, _ := p.Args["orderBy"].(int)
	order := EDoctorsOrderBy.Serialize(orderBy).(string)

	if nameOK || specialistOK || titleOK || addressOK {
		doctors, err := r.db.GetDoctorsByAttribute(limit, offset, order, name, specialist, title, address)
		if err != nil {
			return nil, err
		}

		return doctors, nil
	}

	doctors, err := r.db.GetDoctors(limit, offset, order)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}

// IllnessResolver resolves the doctor illness and manage to respond with single `illness`.
func (r *Resolver) IllnessResolver(p graphql.ResolveParams) (interface{}, error) {
	id, idOK := p.Args["id"].(int)

	if idOK {
		return r.db.GetIllnessByID(id)
	}

	return nil, fmt.Errorf("illness ID can't be empty")
}

// IllnessesResolver resolves the illnesses query and responds with list of `illnesses`
func (r *Resolver) IllnessesResolver(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	risk, riskOK := p.Args["risk"].(int)

	// Here, the error is not relevant, if the limit or offset is not specified
	// it will be set to the default value
	limit, _ := p.Args["limit"].(int)
	offset, _ := p.Args["offset"].(int)

	// Serialize the order
	orderBy, _ := p.Args["orderBy"].(int)
	order := EIllnessOrderBy.Serialize(orderBy).(string)

	if nameOK || riskOK {
		// Because the default value of int is always 0 if not initialized,
		// It will be problematic if there's risk with value of 0.
		// risk should be > 0
		if !riskOK {
			risk = -1
		}

		// call `GetIllnessesByAttribute` to filter the illnesses
		// and then return them
		return r.db.GetIllnessesByAttribute(limit, offset, order, name, risk)
	}

	return r.db.GetIllnesses(limit, offset, order)
}

// UserResolver resolves the user query and responds with a single `user`
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, idOK := p.Args["id"].(int)

	if idOK {
		return r.db.GetUserByID(id)
	}

	return nil, fmt.Errorf("User ID can't be empty")
}

// UsersResolver resolves the users query and responds with a list `users`
func (r *Resolver) UsersResolver(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	older, olderOK := p.Args["older"].(int)
	younger, youngerOK := p.Args["younger"].(int)
	email, emailOK := p.Args["email"].(string)
	age, ageOK := p.Args["age"].(int)
	birthDate, birthDateOK := p.Args["birthDate"].(time.Time)

	// Here, the error is not relevant, if the limit or offset is not specified
	// it will be set to the default value
	limit, _ := p.Args["limit"].(int)
	offset, _ := p.Args["offset"].(int)

	// Serialize the order
	orderBy, _ := p.Args["orderBy"].(int)
	order := EUsersOrderBy.Serialize(orderBy).(string)

	if nameOK ||
		olderOK ||
		youngerOK ||
		emailOK ||
		ageOK ||
		birthDateOK {
		return r.db.GetUsersByAttribute(limit, offset, order, name, email, older, younger, age, birthDate)
	}

	return r.db.GetUsers(limit, offset, order)
}
