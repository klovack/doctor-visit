package gql

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/klovack/doctor-visit/db"
	"github.com/klovack/doctor-visit/models"
)

// Resolver struct holds a connection to the database,
// more on [resolver](https://graphql.org/learn/execution/#root-fields-resolvers)
type Resolver struct {
	db *db.DB
}

// ---------- DoctorResolver -------------

// FindDoctor resolves the doctor query and manage to respond with single `doctor`
func (r *Resolver) FindDoctor(p graphql.ResolveParams) (interface{}, error) {
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

// FindDoctors resolves the doctors query and manage call to the `GetDoctorsByAttribute` and respond with list of `doctors`
func (r *Resolver) FindDoctors(p graphql.ResolveParams) (interface{}, error) {
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

// FindIllness resolves the doctor illness and manage to respond with single `illness`.
func (r *Resolver) FindIllness(p graphql.ResolveParams) (interface{}, error) {
	id, idOK := p.Args["id"].(int)

	if idOK {
		return r.db.GetIllnessByID(id)
	}

	return nil, fmt.Errorf("illness ID can't be empty")
}

// FindIllnesses resolves the illnesses query and responds with list of `illnesses`
func (r *Resolver) FindIllnesses(p graphql.ResolveParams) (interface{}, error) {
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

// FindUser resolves the user query and responds with a single `user`
func (r *Resolver) FindUser(p graphql.ResolveParams) (interface{}, error) {
	id, idOK := p.Args["id"].(int)

	if idOK {
		return r.db.GetUserByID(id)
	}

	return nil, fmt.Errorf("User ID can't be empty")
}

// FindUserByEmail resolves the `userByEmail` query and responds with a single `user`
func (r *Resolver) FindUserByEmail(p graphql.ResolveParams) (interface{}, error) {
	email, emailOK := p.Args["email"].(string)

	if emailOK {
		return r.db.GetUserByEmail(email)
	}

	return nil, fmt.Errorf("Email can't be empty")
}

// FindUsers resolves the users query and responds with a list `users`
func (r *Resolver) FindUsers(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	email, emailOK := p.Args["email"].(string)
	birthDate, birthDateOK := p.Args["birthDate"].(time.Time)

	// Here, the error is not relevant, if the limit or offset is not specified
	// it will be set to the default value
	limit, _ := p.Args["limit"].(int)
	offset, _ := p.Args["offset"].(int)

	// Serialize the order
	orderBy, _ := p.Args["orderBy"].(int)
	order := EUsersOrderBy.Serialize(orderBy).(string)

	if nameOK ||
		emailOK ||
		birthDateOK {
		return r.db.GetUsersByAttribute(limit, offset, order, name, email, birthDate)
	}

	return r.db.GetUsers(limit, offset, order)
}

// FindFavoritesDoctors returns all favorite_doctors for the given user.
// This Resolver is not the best, but it works. It heavily needs better implementation to
// optimize request to the database
func (r *Resolver) FindFavoritesDoctors(p graphql.ResolveParams) (interface{}, error) {
	user := p.Source.(*models.User)
	return r.db.GetDoctorsForUser(user)
}
