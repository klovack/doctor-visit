// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Visit is an object representing the database table.
type Visit struct {
	VisitID       int       `boil:"visit_id" json:"visit_id" toml:"visit_id" yaml:"visit_id"`
	UserIllnessID null.Int  `boil:"user_illness_id" json:"user_illness_id,omitempty" toml:"user_illness_id" yaml:"user_illness_id,omitempty"`
	DoctorID      null.Int  `boil:"doctor_id" json:"doctor_id,omitempty" toml:"doctor_id" yaml:"doctor_id,omitempty"`
	UserID        null.Int  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	VisitDate     time.Time `boil:"visit_date" json:"visit_date" toml:"visit_date" yaml:"visit_date"`

	R *visitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L visitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VisitColumns = struct {
	VisitID       string
	UserIllnessID string
	DoctorID      string
	UserID        string
	VisitDate     string
}{
	VisitID:       "visit_id",
	UserIllnessID: "user_illness_id",
	DoctorID:      "doctor_id",
	UserID:        "user_id",
	VisitDate:     "visit_date",
}

// Generated where

var VisitWhere = struct {
	VisitID       whereHelperint
	UserIllnessID whereHelpernull_Int
	DoctorID      whereHelpernull_Int
	UserID        whereHelpernull_Int
	VisitDate     whereHelpertime_Time
}{
	VisitID:       whereHelperint{field: "\"visit\".\"visit_id\""},
	UserIllnessID: whereHelpernull_Int{field: "\"visit\".\"user_illness_id\""},
	DoctorID:      whereHelpernull_Int{field: "\"visit\".\"doctor_id\""},
	UserID:        whereHelpernull_Int{field: "\"visit\".\"user_id\""},
	VisitDate:     whereHelpertime_Time{field: "\"visit\".\"visit_date\""},
}

// VisitRels is where relationship names are stored.
var VisitRels = struct {
	Doctor      string
	User        string
	UserIllness string
}{
	Doctor:      "Doctor",
	User:        "User",
	UserIllness: "UserIllness",
}

// visitR is where relationships are stored.
type visitR struct {
	Doctor      *Doctor
	User        *User
	UserIllness *UserIllness
}

// NewStruct creates a new relationship struct
func (*visitR) NewStruct() *visitR {
	return &visitR{}
}

// visitL is where Load methods for each relationship are stored.
type visitL struct{}

var (
	visitAllColumns            = []string{"visit_id", "user_illness_id", "doctor_id", "user_id", "visit_date"}
	visitColumnsWithoutDefault = []string{"user_illness_id", "doctor_id", "user_id"}
	visitColumnsWithDefault    = []string{"visit_id", "visit_date"}
	visitPrimaryKeyColumns     = []string{"visit_id"}
)

type (
	// VisitSlice is an alias for a slice of pointers to Visit.
	// This should generally be used opposed to []Visit.
	VisitSlice []*Visit
	// VisitHook is the signature for custom Visit hook methods
	VisitHook func(context.Context, boil.ContextExecutor, *Visit) error

	visitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	visitType                 = reflect.TypeOf(&Visit{})
	visitMapping              = queries.MakeStructMapping(visitType)
	visitPrimaryKeyMapping, _ = queries.BindMapping(visitType, visitMapping, visitPrimaryKeyColumns)
	visitInsertCacheMut       sync.RWMutex
	visitInsertCache          = make(map[string]insertCache)
	visitUpdateCacheMut       sync.RWMutex
	visitUpdateCache          = make(map[string]updateCache)
	visitUpsertCacheMut       sync.RWMutex
	visitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var visitBeforeInsertHooks []VisitHook
var visitBeforeUpdateHooks []VisitHook
var visitBeforeDeleteHooks []VisitHook
var visitBeforeUpsertHooks []VisitHook

var visitAfterInsertHooks []VisitHook
var visitAfterSelectHooks []VisitHook
var visitAfterUpdateHooks []VisitHook
var visitAfterDeleteHooks []VisitHook
var visitAfterUpsertHooks []VisitHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Visit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Visit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Visit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Visit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Visit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Visit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Visit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Visit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Visit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVisitHook registers your hook function for all future operations.
func AddVisitHook(hookPoint boil.HookPoint, visitHook VisitHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		visitBeforeInsertHooks = append(visitBeforeInsertHooks, visitHook)
	case boil.BeforeUpdateHook:
		visitBeforeUpdateHooks = append(visitBeforeUpdateHooks, visitHook)
	case boil.BeforeDeleteHook:
		visitBeforeDeleteHooks = append(visitBeforeDeleteHooks, visitHook)
	case boil.BeforeUpsertHook:
		visitBeforeUpsertHooks = append(visitBeforeUpsertHooks, visitHook)
	case boil.AfterInsertHook:
		visitAfterInsertHooks = append(visitAfterInsertHooks, visitHook)
	case boil.AfterSelectHook:
		visitAfterSelectHooks = append(visitAfterSelectHooks, visitHook)
	case boil.AfterUpdateHook:
		visitAfterUpdateHooks = append(visitAfterUpdateHooks, visitHook)
	case boil.AfterDeleteHook:
		visitAfterDeleteHooks = append(visitAfterDeleteHooks, visitHook)
	case boil.AfterUpsertHook:
		visitAfterUpsertHooks = append(visitAfterUpsertHooks, visitHook)
	}
}

// One returns a single visit record from the query.
func (q visitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Visit, error) {
	o := &Visit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for visit")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Visit records from the query.
func (q visitQuery) All(ctx context.Context, exec boil.ContextExecutor) (VisitSlice, error) {
	var o []*Visit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Visit slice")
	}

	if len(visitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Visit records in the query.
func (q visitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count visit rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q visitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if visit exists")
	}

	return count > 0, nil
}

// Doctor pointed to by the foreign key.
func (o *Visit) Doctor(mods ...qm.QueryMod) doctorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("doctor_id=?", o.DoctorID),
	}

	queryMods = append(queryMods, mods...)

	query := Doctors(queryMods...)
	queries.SetFrom(query.Query, "\"doctor\"")

	return query
}

// User pointed to by the foreign key.
func (o *Visit) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"user\"")

	return query
}

// UserIllness pointed to by the foreign key.
func (o *Visit) UserIllness(mods ...qm.QueryMod) userIllnessQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_illness_id=?", o.UserIllnessID),
	}

	queryMods = append(queryMods, mods...)

	query := UserIllnesses(queryMods...)
	queries.SetFrom(query.Query, "\"user_illness\"")

	return query
}

// LoadDoctor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (visitL) LoadDoctor(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVisit interface{}, mods queries.Applicator) error {
	var slice []*Visit
	var object *Visit

	if singular {
		object = maybeVisit.(*Visit)
	} else {
		slice = *maybeVisit.(*[]*Visit)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &visitR{}
		}
		if !queries.IsNil(object.DoctorID) {
			args = append(args, object.DoctorID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &visitR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DoctorID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DoctorID) {
				args = append(args, obj.DoctorID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`doctor`), qm.WhereIn(`doctor_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Doctor")
	}

	var resultSlice []*Doctor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Doctor")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for doctor")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for doctor")
	}

	if len(visitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Doctor = foreign
		if foreign.R == nil {
			foreign.R = &doctorR{}
		}
		foreign.R.Visits = append(foreign.R.Visits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DoctorID, foreign.DoctorID) {
				local.R.Doctor = foreign
				if foreign.R == nil {
					foreign.R = &doctorR{}
				}
				foreign.R.Visits = append(foreign.R.Visits, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (visitL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVisit interface{}, mods queries.Applicator) error {
	var slice []*Visit
	var object *Visit

	if singular {
		object = maybeVisit.(*Visit)
	} else {
		slice = *maybeVisit.(*[]*Visit)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &visitR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &visitR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`user`), qm.WhereIn(`user_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(visitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Visits = append(foreign.R.Visits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.UserID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Visits = append(foreign.R.Visits, local)
				break
			}
		}
	}

	return nil
}

// LoadUserIllness allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (visitL) LoadUserIllness(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVisit interface{}, mods queries.Applicator) error {
	var slice []*Visit
	var object *Visit

	if singular {
		object = maybeVisit.(*Visit)
	} else {
		slice = *maybeVisit.(*[]*Visit)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &visitR{}
		}
		if !queries.IsNil(object.UserIllnessID) {
			args = append(args, object.UserIllnessID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &visitR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserIllnessID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserIllnessID) {
				args = append(args, obj.UserIllnessID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`user_illness`), qm.WhereIn(`user_illness_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserIllness")
	}

	var resultSlice []*UserIllness
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserIllness")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_illness")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_illness")
	}

	if len(visitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.UserIllness = foreign
		if foreign.R == nil {
			foreign.R = &userIllnessR{}
		}
		foreign.R.Visits = append(foreign.R.Visits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserIllnessID, foreign.UserIllnessID) {
				local.R.UserIllness = foreign
				if foreign.R == nil {
					foreign.R = &userIllnessR{}
				}
				foreign.R.Visits = append(foreign.R.Visits, local)
				break
			}
		}
	}

	return nil
}

// SetDoctor of the visit to the related item.
// Sets o.R.Doctor to related.
// Adds o to related.R.Visits.
func (o *Visit) SetDoctor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Doctor) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"visit\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"doctor_id"}),
		strmangle.WhereClause("\"", "\"", 2, visitPrimaryKeyColumns),
	)
	values := []interface{}{related.DoctorID, o.VisitID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.DoctorID, related.DoctorID)
	if o.R == nil {
		o.R = &visitR{
			Doctor: related,
		}
	} else {
		o.R.Doctor = related
	}

	if related.R == nil {
		related.R = &doctorR{
			Visits: VisitSlice{o},
		}
	} else {
		related.R.Visits = append(related.R.Visits, o)
	}

	return nil
}

// RemoveDoctor relationship.
// Sets o.R.Doctor to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Visit) RemoveDoctor(ctx context.Context, exec boil.ContextExecutor, related *Doctor) error {
	var err error

	queries.SetScanner(&o.DoctorID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("doctor_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Doctor = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Visits {
		if queries.Equal(o.DoctorID, ri.DoctorID) {
			continue
		}

		ln := len(related.R.Visits)
		if ln > 1 && i < ln-1 {
			related.R.Visits[i] = related.R.Visits[ln-1]
		}
		related.R.Visits = related.R.Visits[:ln-1]
		break
	}
	return nil
}

// SetUser of the visit to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Visits.
func (o *Visit) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"visit\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, visitPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.VisitID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.UserID)
	if o.R == nil {
		o.R = &visitR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Visits: VisitSlice{o},
		}
	} else {
		related.R.Visits = append(related.R.Visits, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Visit) RemoveUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.User = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Visits {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.Visits)
		if ln > 1 && i < ln-1 {
			related.R.Visits[i] = related.R.Visits[ln-1]
		}
		related.R.Visits = related.R.Visits[:ln-1]
		break
	}
	return nil
}

// SetUserIllness of the visit to the related item.
// Sets o.R.UserIllness to related.
// Adds o to related.R.Visits.
func (o *Visit) SetUserIllness(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserIllness) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"visit\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_illness_id"}),
		strmangle.WhereClause("\"", "\"", 2, visitPrimaryKeyColumns),
	)
	values := []interface{}{related.UserIllnessID, o.VisitID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserIllnessID, related.UserIllnessID)
	if o.R == nil {
		o.R = &visitR{
			UserIllness: related,
		}
	} else {
		o.R.UserIllness = related
	}

	if related.R == nil {
		related.R = &userIllnessR{
			Visits: VisitSlice{o},
		}
	} else {
		related.R.Visits = append(related.R.Visits, o)
	}

	return nil
}

// RemoveUserIllness relationship.
// Sets o.R.UserIllness to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Visit) RemoveUserIllness(ctx context.Context, exec boil.ContextExecutor, related *UserIllness) error {
	var err error

	queries.SetScanner(&o.UserIllnessID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_illness_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.UserIllness = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Visits {
		if queries.Equal(o.UserIllnessID, ri.UserIllnessID) {
			continue
		}

		ln := len(related.R.Visits)
		if ln > 1 && i < ln-1 {
			related.R.Visits[i] = related.R.Visits[ln-1]
		}
		related.R.Visits = related.R.Visits[:ln-1]
		break
	}
	return nil
}

// Visits retrieves all the records using an executor.
func Visits(mods ...qm.QueryMod) visitQuery {
	mods = append(mods, qm.From("\"visit\""))
	return visitQuery{NewQuery(mods...)}
}

// FindVisit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVisit(ctx context.Context, exec boil.ContextExecutor, visitID int, selectCols ...string) (*Visit, error) {
	visitObj := &Visit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"visit\" where \"visit_id\"=$1", sel,
	)

	q := queries.Raw(query, visitID)

	err := q.Bind(ctx, exec, visitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from visit")
	}

	return visitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Visit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no visit provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(visitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	visitInsertCacheMut.RLock()
	cache, cached := visitInsertCache[key]
	visitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			visitAllColumns,
			visitColumnsWithDefault,
			visitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(visitType, visitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"visit\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"visit\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into visit")
	}

	if !cached {
		visitInsertCacheMut.Lock()
		visitInsertCache[key] = cache
		visitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Visit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Visit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	visitUpdateCacheMut.RLock()
	cache, cached := visitUpdateCache[key]
	visitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			visitAllColumns,
			visitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update visit, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"visit\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, visitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, append(wl, visitPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update visit row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for visit")
	}

	if !cached {
		visitUpdateCacheMut.Lock()
		visitUpdateCache[key] = cache
		visitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q visitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for visit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for visit")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VisitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"visit\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, visitPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in visit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all visit")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Visit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no visit provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(visitColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	visitUpsertCacheMut.RLock()
	cache, cached := visitUpsertCache[key]
	visitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			visitAllColumns,
			visitColumnsWithDefault,
			visitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			visitAllColumns,
			visitPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert visit, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(visitPrimaryKeyColumns))
			copy(conflict, visitPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"visit\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(visitType, visitMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert visit")
	}

	if !cached {
		visitUpsertCacheMut.Lock()
		visitUpsertCache[key] = cache
		visitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Visit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Visit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Visit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), visitPrimaryKeyMapping)
	sql := "DELETE FROM \"visit\" WHERE \"visit_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from visit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for visit")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q visitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no visitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from visit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for visit")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VisitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(visitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"visit\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, visitPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from visit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for visit")
	}

	if len(visitAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Visit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVisit(ctx, exec, o.VisitID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VisitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VisitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"visit\".* FROM \"visit\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, visitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VisitSlice")
	}

	*o = slice

	return nil
}

// VisitExists checks if the Visit row exists.
func VisitExists(ctx context.Context, exec boil.ContextExecutor, visitID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"visit\" where \"visit_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, visitID)
	}

	row := exec.QueryRowContext(ctx, sql, visitID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if visit exists")
	}

	return exists, nil
}