package db

import (
	"time"

	"github.com/klovack/doctor-visit/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetUserByID returns a single user from the database based off of its ID
func (d *DB) GetUserByID(id int) (*models.User, error) {
	return models.Users(qm.Where("user_id=?", id)).One(d.ctx, d)
}

// GetUsersByAttribute returns a list of users from the database based off of its attributes
func (d *DB) GetUsersByAttribute(
	limit, offset int,
	order, name, email string,
	birthdate time.Time,
) (models.UserSlice, error) {
	validatePage(&limit, &offset)

	var query []qm.QueryMod

	if name != "" {
		name = sqllike(name)
		query = append(query, qm.Where("first_name=? OR last_name=?", name))
	}

	// It may be later need refactoring
	// Because search by email should be unique
	if email != "" {
		email = sqllike(email)
		query = append(query, qm.Where("email=?", email))
	}

	query = append(
		query,
		qm.Where("birth_date > make_date(?)", birthdate.Unix()),
		qm.OrderBy(order),
		qm.Limit(limit),
		qm.Offset(offset),
	)

	return models.Users(query...).All(d.ctx, d)
}

// GetUsers returns a list of users from the database
func (d *DB) GetUsers(limit, offset int, order string) (models.UserSlice, error) {
	validatePage(&limit, &offset)

	return models.Users(qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset)).All(d.ctx, d)
}

// GetUserByEmail returns one single User if email matched
func (d *DB) GetUserByEmail(email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email)).One(d.ctx, d)
}
