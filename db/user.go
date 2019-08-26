package db

import (
	"time"

	"github.com/klovack/doctor-visit/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// User is a type for the program, not to be confused with models.User.
// This type has same attributes with models.User but with few more additions
type User struct {
	models.User     `boil:",bind"`
	FavoriteDoctors models.DoctorSlice `boil:"favorite_doctors" json:"favorite_doctors"`
}

// GetUserByID returns a single user from the database based off of its ID
func (d *DB) GetUserByID(id int) (*models.User, error) {
	return models.Users(qm.Where("user_id=?", id)).One(d.ctx, d)
}

// GetUsersByAttribute returns a list of users from the database based off of its attributes
func (d *DB) GetUsersByAttribute(
	limit, offset int,
	order, name, email string,
	older, younger, age int,
	birthdate time.Time,
) (models.UserSlice, error) {
	validatePage(&limit, &offset)

	return nil, nil
}

// GetUsers returns a list of users from the database
func (d *DB) GetUsers(limit, offset int, order string) (models.UserSlice, error) {
	validatePage(&limit, &offset)
	return models.Users(qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset)).All(d.ctx, d)
}
