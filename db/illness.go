package db

import (
	"github.com/klovack/doctor-visit/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetIllnessesByAttribute makes request to the database and returns the slice of illnesses if there's any
func (d *DB) GetIllnessesByAttribute(limit int, offset int, order string, name string, risk int) (models.IllnessSlice, error) {
	validatePage(&limit, &offset)

	var query []qm.QueryMod
	if name != "" {
		name = sqllike(name)
		query = append(query, qm.Where("name ILIKE ?", name))
	}

	if risk >= 0 {
		query = append(query, qm.Where("risk=?", risk))
	}

	query = append(query, qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset))
	return models.Illnesses(query...).All(d.ctx, d)
}

// GetIllnesses returns all `illnesses`
func (d *DB) GetIllnesses(limit, offset int, order string) (models.IllnessSlice, error) {
	validatePage(&limit, &offset)
	return models.Illnesses(qm.OrderBy(order), qm.Limit(limit), qm.Offset(offset)).All(d.ctx, d)
}

// GetIllnessByID prepares the query to get a single illness
func (d *DB) GetIllnessByID(id int) (*models.Illness, error) {
	return models.Illnesses(qm.Where("illness_id=?", id)).One(d.ctx, d)
}
