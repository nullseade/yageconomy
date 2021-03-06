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

func testEconomyWaifuItems(t *testing.T) {
	t.Parallel()

	query := EconomyWaifuItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEconomyWaifuItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
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

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyWaifuItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EconomyWaifuItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyWaifuItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyWaifuItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyWaifuItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EconomyWaifuItemExists(ctx, tx, o.GuildID, o.LocalID)
	if err != nil {
		t.Errorf("Unable to check if EconomyWaifuItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EconomyWaifuItemExists to return true, but got false.")
	}
}

func testEconomyWaifuItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	economyWaifuItemFound, err := FindEconomyWaifuItem(ctx, tx, o.GuildID, o.LocalID)
	if err != nil {
		t.Error(err)
	}

	if economyWaifuItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEconomyWaifuItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EconomyWaifuItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEconomyWaifuItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EconomyWaifuItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEconomyWaifuItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	economyWaifuItemOne := &EconomyWaifuItem{}
	economyWaifuItemTwo := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, economyWaifuItemOne, economyWaifuItemDBTypes, false, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}
	if err = randomize.Struct(seed, economyWaifuItemTwo, economyWaifuItemDBTypes, false, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyWaifuItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyWaifuItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyWaifuItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEconomyWaifuItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	economyWaifuItemOne := &EconomyWaifuItem{}
	economyWaifuItemTwo := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, economyWaifuItemOne, economyWaifuItemDBTypes, false, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}
	if err = randomize.Struct(seed, economyWaifuItemTwo, economyWaifuItemDBTypes, false, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyWaifuItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyWaifuItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEconomyWaifuItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyWaifuItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(economyWaifuItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyWaifuItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
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

func testEconomyWaifuItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyWaifuItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEconomyWaifuItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyWaifuItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	economyWaifuItemDBTypes = map[string]string{`GuildID`: `bigint`, `LocalID`: `bigint`, `Name`: `text`, `Icon`: `text`, `Price`: `integer`, `GamblingBoost`: `integer`}
	_                       = bytes.MinRead
)

func testEconomyWaifuItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(economyWaifuItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(economyWaifuItemAllColumns) == len(economyWaifuItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEconomyWaifuItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(economyWaifuItemAllColumns) == len(economyWaifuItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyWaifuItem{}
	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyWaifuItemDBTypes, true, economyWaifuItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(economyWaifuItemAllColumns, economyWaifuItemPrimaryKeyColumns) {
		fields = economyWaifuItemAllColumns
	} else {
		fields = strmangle.SetComplement(
			economyWaifuItemAllColumns,
			economyWaifuItemPrimaryKeyColumns,
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

	slice := EconomyWaifuItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEconomyWaifuItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(economyWaifuItemAllColumns) == len(economyWaifuItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EconomyWaifuItem{}
	if err = randomize.Struct(seed, &o, economyWaifuItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyWaifuItem: %s", err)
	}

	count, err := EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, economyWaifuItemDBTypes, false, economyWaifuItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyWaifuItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyWaifuItem: %s", err)
	}

	count, err = EconomyWaifuItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
