package db

import (
	"fmt"

	"github.com/klovack/doctor-visit/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetDoctorsByAttribute is called within our doctor query for graphql
// If limit is set to 0, the function will use the DefaultLimit instead.
func (d *DB) GetDoctorsByAttribute(limit int, offset int, order, name, specialist, title, address string) (models.DoctorSlice, error) {

	validatePage(&limit, &offset)

	var query []qm.QueryMod

	if name != "" {
		name = sqllike(name)
		query = append(query, qm.Where("name ILIKE ?", name))
	}

	if title != "" {
		title = sqllike(title)
		query = append(query, qm.Where("title ILIKE ?", title))
	}

	if specialist != "" {
		specialist = sqllike(specialist)
		query = append(query, qm.Where("specialist ILIKE ?", specialist))
	}

	if address != "" {
		address = sqllike(address)
		query = append(query, qm.Where("address ILIKE ?", address))
	}

	query = append(query, qm.OrderBy("name ASC"), qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset))

	// Prepare query
	doctors, err := models.Doctors(query...).All(d.ctx, d)
	// doctors, err := models.Doctors(models.DoctorWhere.Name.EQ(null.StringFrom(name))).All(d.ctx, d)
	if err != nil {
		return nil, fmt.Errorf("Failed to get the list of doctors: %v", err)
	}

	return doctors, nil
}

// GetDoctors returns all doctors
func (d *DB) GetDoctors(limit, offset int, order string) (models.DoctorSlice, error) {
	validatePage(&limit, &offset)
	return models.Doctors(qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset)).All(d.ctx, d)
}

// GetDoctorByID prepares the query to get a single doctor
func (d *DB) GetDoctorByID(id int) (*models.Doctor, error) {
	return models.Doctors(qm.Where("doctor_id=?", id)).One(d.ctx, d)
}
