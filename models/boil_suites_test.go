// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Doctors", testDoctors)
	t.Run("Illnesses", testIllnesses)
	t.Run("Users", testUsers)
	t.Run("UserIllnesses", testUserIllnesses)
	t.Run("Visits", testVisits)
}

func TestDelete(t *testing.T) {
	t.Run("Doctors", testDoctorsDelete)
	t.Run("Illnesses", testIllnessesDelete)
	t.Run("Users", testUsersDelete)
	t.Run("UserIllnesses", testUserIllnessesDelete)
	t.Run("Visits", testVisitsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Doctors", testDoctorsQueryDeleteAll)
	t.Run("Illnesses", testIllnessesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("UserIllnesses", testUserIllnessesQueryDeleteAll)
	t.Run("Visits", testVisitsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Doctors", testDoctorsSliceDeleteAll)
	t.Run("Illnesses", testIllnessesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("UserIllnesses", testUserIllnessesSliceDeleteAll)
	t.Run("Visits", testVisitsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Doctors", testDoctorsExists)
	t.Run("Illnesses", testIllnessesExists)
	t.Run("Users", testUsersExists)
	t.Run("UserIllnesses", testUserIllnessesExists)
	t.Run("Visits", testVisitsExists)
}

func TestFind(t *testing.T) {
	t.Run("Doctors", testDoctorsFind)
	t.Run("Illnesses", testIllnessesFind)
	t.Run("Users", testUsersFind)
	t.Run("UserIllnesses", testUserIllnessesFind)
	t.Run("Visits", testVisitsFind)
}

func TestBind(t *testing.T) {
	t.Run("Doctors", testDoctorsBind)
	t.Run("Illnesses", testIllnessesBind)
	t.Run("Users", testUsersBind)
	t.Run("UserIllnesses", testUserIllnessesBind)
	t.Run("Visits", testVisitsBind)
}

func TestOne(t *testing.T) {
	t.Run("Doctors", testDoctorsOne)
	t.Run("Illnesses", testIllnessesOne)
	t.Run("Users", testUsersOne)
	t.Run("UserIllnesses", testUserIllnessesOne)
	t.Run("Visits", testVisitsOne)
}

func TestAll(t *testing.T) {
	t.Run("Doctors", testDoctorsAll)
	t.Run("Illnesses", testIllnessesAll)
	t.Run("Users", testUsersAll)
	t.Run("UserIllnesses", testUserIllnessesAll)
	t.Run("Visits", testVisitsAll)
}

func TestCount(t *testing.T) {
	t.Run("Doctors", testDoctorsCount)
	t.Run("Illnesses", testIllnessesCount)
	t.Run("Users", testUsersCount)
	t.Run("UserIllnesses", testUserIllnessesCount)
	t.Run("Visits", testVisitsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Doctors", testDoctorsHooks)
	t.Run("Illnesses", testIllnessesHooks)
	t.Run("Users", testUsersHooks)
	t.Run("UserIllnesses", testUserIllnessesHooks)
	t.Run("Visits", testVisitsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Doctors", testDoctorsInsert)
	t.Run("Doctors", testDoctorsInsertWhitelist)
	t.Run("Illnesses", testIllnessesInsert)
	t.Run("Illnesses", testIllnessesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("UserIllnesses", testUserIllnessesInsert)
	t.Run("UserIllnesses", testUserIllnessesInsertWhitelist)
	t.Run("Visits", testVisitsInsert)
	t.Run("Visits", testVisitsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserIllnessToDoctorUsingDoctor", testUserIllnessToOneDoctorUsingDoctor)
	t.Run("UserIllnessToIllnessUsingIllness", testUserIllnessToOneIllnessUsingIllness)
	t.Run("UserIllnessToUserUsingUser", testUserIllnessToOneUserUsingUser)
	t.Run("VisitToDoctorUsingDoctor", testVisitToOneDoctorUsingDoctor)
	t.Run("VisitToUserUsingUser", testVisitToOneUserUsingUser)
	t.Run("VisitToUserIllnessUsingUserIllness", testVisitToOneUserIllnessUsingUserIllness)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DoctorToUsers", testDoctorToManyUsers)
	t.Run("DoctorToUserIllnesses", testDoctorToManyUserIllnesses)
	t.Run("DoctorToVisits", testDoctorToManyVisits)
	t.Run("IllnessToUserIllnesses", testIllnessToManyUserIllnesses)
	t.Run("UserToDoctors", testUserToManyDoctors)
	t.Run("UserToUserIllnesses", testUserToManyUserIllnesses)
	t.Run("UserToVisits", testUserToManyVisits)
	t.Run("UserIllnessToVisits", testUserIllnessToManyVisits)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserIllnessToDoctorUsingUserIllnesses", testUserIllnessToOneSetOpDoctorUsingDoctor)
	t.Run("UserIllnessToIllnessUsingUserIllnesses", testUserIllnessToOneSetOpIllnessUsingIllness)
	t.Run("UserIllnessToUserUsingUserIllnesses", testUserIllnessToOneSetOpUserUsingUser)
	t.Run("VisitToDoctorUsingVisits", testVisitToOneSetOpDoctorUsingDoctor)
	t.Run("VisitToUserUsingVisits", testVisitToOneSetOpUserUsingUser)
	t.Run("VisitToUserIllnessUsingVisits", testVisitToOneSetOpUserIllnessUsingUserIllness)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserIllnessToDoctorUsingUserIllnesses", testUserIllnessToOneRemoveOpDoctorUsingDoctor)
	t.Run("UserIllnessToIllnessUsingUserIllnesses", testUserIllnessToOneRemoveOpIllnessUsingIllness)
	t.Run("UserIllnessToUserUsingUserIllnesses", testUserIllnessToOneRemoveOpUserUsingUser)
	t.Run("VisitToDoctorUsingVisits", testVisitToOneRemoveOpDoctorUsingDoctor)
	t.Run("VisitToUserUsingVisits", testVisitToOneRemoveOpUserUsingUser)
	t.Run("VisitToUserIllnessUsingVisits", testVisitToOneRemoveOpUserIllnessUsingUserIllness)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("DoctorToUsers", testDoctorToManyAddOpUsers)
	t.Run("DoctorToUserIllnesses", testDoctorToManyAddOpUserIllnesses)
	t.Run("DoctorToVisits", testDoctorToManyAddOpVisits)
	t.Run("IllnessToUserIllnesses", testIllnessToManyAddOpUserIllnesses)
	t.Run("UserToDoctors", testUserToManyAddOpDoctors)
	t.Run("UserToUserIllnesses", testUserToManyAddOpUserIllnesses)
	t.Run("UserToVisits", testUserToManyAddOpVisits)
	t.Run("UserIllnessToVisits", testUserIllnessToManyAddOpVisits)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("DoctorToUsers", testDoctorToManySetOpUsers)
	t.Run("DoctorToUserIllnesses", testDoctorToManySetOpUserIllnesses)
	t.Run("DoctorToVisits", testDoctorToManySetOpVisits)
	t.Run("IllnessToUserIllnesses", testIllnessToManySetOpUserIllnesses)
	t.Run("UserToDoctors", testUserToManySetOpDoctors)
	t.Run("UserToUserIllnesses", testUserToManySetOpUserIllnesses)
	t.Run("UserToVisits", testUserToManySetOpVisits)
	t.Run("UserIllnessToVisits", testUserIllnessToManySetOpVisits)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("DoctorToUsers", testDoctorToManyRemoveOpUsers)
	t.Run("DoctorToUserIllnesses", testDoctorToManyRemoveOpUserIllnesses)
	t.Run("DoctorToVisits", testDoctorToManyRemoveOpVisits)
	t.Run("IllnessToUserIllnesses", testIllnessToManyRemoveOpUserIllnesses)
	t.Run("UserToDoctors", testUserToManyRemoveOpDoctors)
	t.Run("UserToUserIllnesses", testUserToManyRemoveOpUserIllnesses)
	t.Run("UserToVisits", testUserToManyRemoveOpVisits)
	t.Run("UserIllnessToVisits", testUserIllnessToManyRemoveOpVisits)
}

func TestReload(t *testing.T) {
	t.Run("Doctors", testDoctorsReload)
	t.Run("Illnesses", testIllnessesReload)
	t.Run("Users", testUsersReload)
	t.Run("UserIllnesses", testUserIllnessesReload)
	t.Run("Visits", testVisitsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Doctors", testDoctorsReloadAll)
	t.Run("Illnesses", testIllnessesReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("UserIllnesses", testUserIllnessesReloadAll)
	t.Run("Visits", testVisitsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Doctors", testDoctorsSelect)
	t.Run("Illnesses", testIllnessesSelect)
	t.Run("Users", testUsersSelect)
	t.Run("UserIllnesses", testUserIllnessesSelect)
	t.Run("Visits", testVisitsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Doctors", testDoctorsUpdate)
	t.Run("Illnesses", testIllnessesUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("UserIllnesses", testUserIllnessesUpdate)
	t.Run("Visits", testVisitsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Doctors", testDoctorsSliceUpdateAll)
	t.Run("Illnesses", testIllnessesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("UserIllnesses", testUserIllnessesSliceUpdateAll)
	t.Run("Visits", testVisitsSliceUpdateAll)
}
