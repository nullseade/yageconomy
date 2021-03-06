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

func testEconomyLogs(t *testing.T) {
	t.Parallel()

	query := EconomyLogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEconomyLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
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

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EconomyLogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyLogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EconomyLogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EconomyLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EconomyLogExists to return true, but got false.")
	}
}

func testEconomyLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	economyLogFound, err := FindEconomyLog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if economyLogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEconomyLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EconomyLogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEconomyLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EconomyLogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEconomyLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	economyLogOne := &EconomyLog{}
	economyLogTwo := &EconomyLog{}
	if err = randomize.Struct(seed, economyLogOne, economyLogDBTypes, false, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}
	if err = randomize.Struct(seed, economyLogTwo, economyLogDBTypes, false, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEconomyLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	economyLogOne := &EconomyLog{}
	economyLogTwo := &EconomyLog{}
	if err = randomize.Struct(seed, economyLogOne, economyLogDBTypes, false, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}
	if err = randomize.Struct(seed, economyLogTwo, economyLogDBTypes, false, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEconomyLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(economyLogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
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

func testEconomyLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyLogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEconomyLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	economyLogDBTypes = map[string]string{`ID`: `bigint`, `GuildID`: `bigint`, `AuthorID`: `bigint`, `TargetID`: `bigint`, `Amount`: `bigint`, `Action`: `smallint`}
	_                 = bytes.MinRead
)

func testEconomyLogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(economyLogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(economyLogAllColumns) == len(economyLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEconomyLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(economyLogAllColumns) == len(economyLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyLog{}
	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyLogDBTypes, true, economyLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(economyLogAllColumns, economyLogPrimaryKeyColumns) {
		fields = economyLogAllColumns
	} else {
		fields = strmangle.SetComplement(
			economyLogAllColumns,
			economyLogPrimaryKeyColumns,
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

	slice := EconomyLogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEconomyLogsUpsert(t *testing.T) {
	t.Parallel()

	if len(economyLogAllColumns) == len(economyLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EconomyLog{}
	if err = randomize.Struct(seed, &o, economyLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyLog: %s", err)
	}

	count, err := EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, economyLogDBTypes, false, economyLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyLog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyLog: %s", err)
	}

	count, err = EconomyLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
