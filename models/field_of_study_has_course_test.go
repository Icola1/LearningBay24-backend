// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFieldOfStudyHasCourses(t *testing.T) {
	t.Parallel()

	query := FieldOfStudyHasCourses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFieldOfStudyHasCoursesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
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

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFieldOfStudyHasCoursesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FieldOfStudyHasCourses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFieldOfStudyHasCoursesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FieldOfStudyHasCourseSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFieldOfStudyHasCoursesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FieldOfStudyHasCourseExists(ctx, tx, o.FieldOfStudyID, o.CourseID)
	if err != nil {
		t.Errorf("Unable to check if FieldOfStudyHasCourse exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FieldOfStudyHasCourseExists to return true, but got false.")
	}
}

func testFieldOfStudyHasCoursesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fieldOfStudyHasCourseFound, err := FindFieldOfStudyHasCourse(ctx, tx, o.FieldOfStudyID, o.CourseID)
	if err != nil {
		t.Error(err)
	}

	if fieldOfStudyHasCourseFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFieldOfStudyHasCoursesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FieldOfStudyHasCourses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFieldOfStudyHasCoursesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FieldOfStudyHasCourses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFieldOfStudyHasCoursesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fieldOfStudyHasCourseOne := &FieldOfStudyHasCourse{}
	fieldOfStudyHasCourseTwo := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, fieldOfStudyHasCourseOne, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}
	if err = randomize.Struct(seed, fieldOfStudyHasCourseTwo, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fieldOfStudyHasCourseOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fieldOfStudyHasCourseTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FieldOfStudyHasCourses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFieldOfStudyHasCoursesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fieldOfStudyHasCourseOne := &FieldOfStudyHasCourse{}
	fieldOfStudyHasCourseTwo := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, fieldOfStudyHasCourseOne, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}
	if err = randomize.Struct(seed, fieldOfStudyHasCourseTwo, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fieldOfStudyHasCourseOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fieldOfStudyHasCourseTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fieldOfStudyHasCourseBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func fieldOfStudyHasCourseAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FieldOfStudyHasCourse) error {
	*o = FieldOfStudyHasCourse{}
	return nil
}

func testFieldOfStudyHasCoursesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FieldOfStudyHasCourse{}
	o := &FieldOfStudyHasCourse{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse object: %s", err)
	}

	AddFieldOfStudyHasCourseHook(boil.BeforeInsertHook, fieldOfStudyHasCourseBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseBeforeInsertHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.AfterInsertHook, fieldOfStudyHasCourseAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseAfterInsertHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.AfterSelectHook, fieldOfStudyHasCourseAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseAfterSelectHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.BeforeUpdateHook, fieldOfStudyHasCourseBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseBeforeUpdateHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.AfterUpdateHook, fieldOfStudyHasCourseAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseAfterUpdateHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.BeforeDeleteHook, fieldOfStudyHasCourseBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseBeforeDeleteHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.AfterDeleteHook, fieldOfStudyHasCourseAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseAfterDeleteHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.BeforeUpsertHook, fieldOfStudyHasCourseBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseBeforeUpsertHooks = []FieldOfStudyHasCourseHook{}

	AddFieldOfStudyHasCourseHook(boil.AfterUpsertHook, fieldOfStudyHasCourseAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fieldOfStudyHasCourseAfterUpsertHooks = []FieldOfStudyHasCourseHook{}
}

func testFieldOfStudyHasCoursesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFieldOfStudyHasCoursesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fieldOfStudyHasCourseColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFieldOfStudyHasCourseToOneCourseUsingCourse(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local FieldOfStudyHasCourse
	var foreign Course

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, courseDBTypes, false, courseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Course struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CourseID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Course().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FieldOfStudyHasCourseSlice{&local}
	if err = local.L.LoadCourse(ctx, tx, false, (*[]*FieldOfStudyHasCourse)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Course = nil
	if err = local.L.LoadCourse(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFieldOfStudyHasCourseToOneFieldOfStudyUsingFieldOfStudy(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local FieldOfStudyHasCourse
	var foreign FieldOfStudy

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, fieldOfStudyDBTypes, false, fieldOfStudyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudy struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FieldOfStudyID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FieldOfStudy().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FieldOfStudyHasCourseSlice{&local}
	if err = local.L.LoadFieldOfStudy(ctx, tx, false, (*[]*FieldOfStudyHasCourse)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FieldOfStudy == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FieldOfStudy = nil
	if err = local.L.LoadFieldOfStudy(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FieldOfStudy == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFieldOfStudyHasCourseToOneSetOpCourseUsingCourse(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FieldOfStudyHasCourse
	var b, c Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fieldOfStudyHasCourseDBTypes, false, strmangle.SetComplement(fieldOfStudyHasCoursePrimaryKeyColumns, fieldOfStudyHasCourseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Course{&b, &c} {
		err = a.SetCourse(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Course != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FieldOfStudyHasCourses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CourseID != x.ID {
			t.Error("foreign key was wrong value", a.CourseID)
		}

		if exists, err := FieldOfStudyHasCourseExists(ctx, tx, a.FieldOfStudyID, a.CourseID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testFieldOfStudyHasCourseToOneSetOpFieldOfStudyUsingFieldOfStudy(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FieldOfStudyHasCourse
	var b, c FieldOfStudy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fieldOfStudyHasCourseDBTypes, false, strmangle.SetComplement(fieldOfStudyHasCoursePrimaryKeyColumns, fieldOfStudyHasCourseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, fieldOfStudyDBTypes, false, strmangle.SetComplement(fieldOfStudyPrimaryKeyColumns, fieldOfStudyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fieldOfStudyDBTypes, false, strmangle.SetComplement(fieldOfStudyPrimaryKeyColumns, fieldOfStudyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FieldOfStudy{&b, &c} {
		err = a.SetFieldOfStudy(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FieldOfStudy != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FieldOfStudyHasCourses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FieldOfStudyID != x.ID {
			t.Error("foreign key was wrong value", a.FieldOfStudyID)
		}

		if exists, err := FieldOfStudyHasCourseExists(ctx, tx, a.FieldOfStudyID, a.CourseID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testFieldOfStudyHasCoursesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
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

func testFieldOfStudyHasCoursesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FieldOfStudyHasCourseSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFieldOfStudyHasCoursesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FieldOfStudyHasCourses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fieldOfStudyHasCourseDBTypes = map[string]string{`FieldOfStudyID`: `int`, `CourseID`: `int`, `Semester`: `int`}
	_                            = bytes.MinRead
)

func testFieldOfStudyHasCoursesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(fieldOfStudyHasCoursePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fieldOfStudyHasCourseAllColumns) == len(fieldOfStudyHasCoursePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCoursePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFieldOfStudyHasCoursesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fieldOfStudyHasCourseAllColumns) == len(fieldOfStudyHasCoursePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCourseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fieldOfStudyHasCourseDBTypes, true, fieldOfStudyHasCoursePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fieldOfStudyHasCourseAllColumns, fieldOfStudyHasCoursePrimaryKeyColumns) {
		fields = fieldOfStudyHasCourseAllColumns
	} else {
		fields = strmangle.SetComplement(
			fieldOfStudyHasCourseAllColumns,
			fieldOfStudyHasCoursePrimaryKeyColumns,
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

	slice := FieldOfStudyHasCourseSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFieldOfStudyHasCoursesUpsert(t *testing.T) {
	t.Parallel()

	if len(fieldOfStudyHasCourseAllColumns) == len(fieldOfStudyHasCoursePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFieldOfStudyHasCourseUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FieldOfStudyHasCourse{}
	if err = randomize.Struct(seed, &o, fieldOfStudyHasCourseDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FieldOfStudyHasCourse: %s", err)
	}

	count, err := FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fieldOfStudyHasCourseDBTypes, false, fieldOfStudyHasCoursePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FieldOfStudyHasCourse struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FieldOfStudyHasCourse: %s", err)
	}

	count, err = FieldOfStudyHasCourses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
