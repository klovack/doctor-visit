// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testIllnesses(t *testing.T) {
	t.Parallel()

	query := Illnesses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIllnessesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIllnessesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Illnesses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIllnessesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IllnessSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIllnessesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IllnessExists(ctx, tx, o.IllnessID)
	if err != nil {
		t.Errorf("Unable to check if Illness exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IllnessExists to return true, but got false.")
	}
}

func testIllnessesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	illnessFound, err := FindIllness(ctx, tx, o.IllnessID)
	if err != nil {
		t.Error(err)
	}

	if illnessFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIllnessesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Illnesses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIllnessesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Illnesses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIllnessesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	illnessOne := &Illness{}
	illnessTwo := &Illness{}
	if err = randomize.Struct(seed, illnessOne, illnessDBTypes, false, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}
	if err = randomize.Struct(seed, illnessTwo, illnessDBTypes, false, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = illnessOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = illnessTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Illnesses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIllnessesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	illnessOne := &Illness{}
	illnessTwo := &Illness{}
	if err = randomize.Struct(seed, illnessOne, illnessDBTypes, false, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}
	if err = randomize.Struct(seed, illnessTwo, illnessDBTypes, false, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = illnessOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = illnessTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func illnessBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func illnessAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Illness) error {
	*o = Illness{}
	return nil
}

func testIllnessesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Illness{}
	o := &Illness{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, illnessDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Illness object: %s", err)
	}

	AddIllnessHook(boil.BeforeInsertHook, illnessBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	illnessBeforeInsertHooks = []IllnessHook{}

	AddIllnessHook(boil.AfterInsertHook, illnessAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	illnessAfterInsertHooks = []IllnessHook{}

	AddIllnessHook(boil.AfterSelectHook, illnessAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	illnessAfterSelectHooks = []IllnessHook{}

	AddIllnessHook(boil.BeforeUpdateHook, illnessBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	illnessBeforeUpdateHooks = []IllnessHook{}

	AddIllnessHook(boil.AfterUpdateHook, illnessAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	illnessAfterUpdateHooks = []IllnessHook{}

	AddIllnessHook(boil.BeforeDeleteHook, illnessBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	illnessBeforeDeleteHooks = []IllnessHook{}

	AddIllnessHook(boil.AfterDeleteHook, illnessAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	illnessAfterDeleteHooks = []IllnessHook{}

	AddIllnessHook(boil.BeforeUpsertHook, illnessBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	illnessBeforeUpsertHooks = []IllnessHook{}

	AddIllnessHook(boil.AfterUpsertHook, illnessAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	illnessAfterUpsertHooks = []IllnessHook{}
}

func testIllnessesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIllnessesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(illnessColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIllnessToManyUserIllnesses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Illness
	var b, c UserIllness

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userIllnessDBTypes, false, userIllnessColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userIllnessDBTypes, false, userIllnessColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.IllnessID, a.IllnessID)
	queries.Assign(&c.IllnessID, a.IllnessID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserIllnesses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.IllnessID, b.IllnessID) {
			bFound = true
		}
		if queries.Equal(v.IllnessID, c.IllnessID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := IllnessSlice{&a}
	if err = a.L.LoadUserIllnesses(ctx, tx, false, (*[]*Illness)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserIllnesses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserIllnesses = nil
	if err = a.L.LoadUserIllnesses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserIllnesses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testIllnessToManyAddOpUserIllnesses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Illness
	var b, c, d, e UserIllness

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, illnessDBTypes, false, strmangle.SetComplement(illnessPrimaryKeyColumns, illnessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIllness{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIllnessDBTypes, false, strmangle.SetComplement(userIllnessPrimaryKeyColumns, userIllnessColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UserIllness{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserIllnesses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.IllnessID, first.IllnessID) {
			t.Error("foreign key was wrong value", a.IllnessID, first.IllnessID)
		}
		if !queries.Equal(a.IllnessID, second.IllnessID) {
			t.Error("foreign key was wrong value", a.IllnessID, second.IllnessID)
		}

		if first.R.Illness != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Illness != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserIllnesses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserIllnesses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserIllnesses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testIllnessToManySetOpUserIllnesses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Illness
	var b, c, d, e UserIllness

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, illnessDBTypes, false, strmangle.SetComplement(illnessPrimaryKeyColumns, illnessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIllness{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIllnessDBTypes, false, strmangle.SetComplement(userIllnessPrimaryKeyColumns, userIllnessColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUserIllnesses(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserIllnesses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserIllnesses(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserIllnesses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.IllnessID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.IllnessID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.IllnessID, d.IllnessID) {
		t.Error("foreign key was wrong value", a.IllnessID, d.IllnessID)
	}
	if !queries.Equal(a.IllnessID, e.IllnessID) {
		t.Error("foreign key was wrong value", a.IllnessID, e.IllnessID)
	}

	if b.R.Illness != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Illness != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Illness != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Illness != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserIllnesses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserIllnesses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testIllnessToManyRemoveOpUserIllnesses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Illness
	var b, c, d, e UserIllness

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, illnessDBTypes, false, strmangle.SetComplement(illnessPrimaryKeyColumns, illnessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIllness{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIllnessDBTypes, false, strmangle.SetComplement(userIllnessPrimaryKeyColumns, userIllnessColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserIllnesses(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserIllnesses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserIllnesses(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserIllnesses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.IllnessID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.IllnessID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Illness != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Illness != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Illness != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Illness != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserIllnesses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserIllnesses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserIllnesses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testIllnessesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIllnessesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IllnessSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIllnessesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Illnesses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	illnessDBTypes = map[string]string{`IllnessID`: `integer`, `Name`: `character varying`, `Risk`: `integer`}
	_              = bytes.MinRead
)

func testIllnessesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(illnessPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(illnessAllColumns) == len(illnessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIllnessesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(illnessAllColumns) == len(illnessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Illness{}
	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, illnessDBTypes, true, illnessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(illnessAllColumns, illnessPrimaryKeyColumns) {
		fields = illnessAllColumns
	} else {
		fields = strmangle.SetComplement(
			illnessAllColumns,
			illnessPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := IllnessSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIllnessesUpsert(t *testing.T) {
	t.Parallel()

	if len(illnessAllColumns) == len(illnessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Illness{}
	if err = randomize.Struct(seed, &o, illnessDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Illness: %s", err)
	}

	count, err := Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, illnessDBTypes, false, illnessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Illness struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Illness: %s", err)
	}

	count, err = Illnesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}