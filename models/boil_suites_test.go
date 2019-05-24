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
	t.Run("EconomyConfigs", testEconomyConfigs)
	t.Run("EconomyLogs", testEconomyLogs)
	t.Run("EconomyPickImages", testEconomyPickImages)
	t.Run("EconomyPlants", testEconomyPlants)
	t.Run("EconomyShopItems", testEconomyShopItems)
	t.Run("EconomyShopListItems", testEconomyShopListItems)
	t.Run("EconomyUsers", testEconomyUsers)
	t.Run("EconomyWaifuItems", testEconomyWaifuItems)
}

func TestDelete(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsDelete)
	t.Run("EconomyLogs", testEconomyLogsDelete)
	t.Run("EconomyPickImages", testEconomyPickImagesDelete)
	t.Run("EconomyPlants", testEconomyPlantsDelete)
	t.Run("EconomyShopItems", testEconomyShopItemsDelete)
	t.Run("EconomyShopListItems", testEconomyShopListItemsDelete)
	t.Run("EconomyUsers", testEconomyUsersDelete)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsQueryDeleteAll)
	t.Run("EconomyLogs", testEconomyLogsQueryDeleteAll)
	t.Run("EconomyPickImages", testEconomyPickImagesQueryDeleteAll)
	t.Run("EconomyPlants", testEconomyPlantsQueryDeleteAll)
	t.Run("EconomyShopItems", testEconomyShopItemsQueryDeleteAll)
	t.Run("EconomyShopListItems", testEconomyShopListItemsQueryDeleteAll)
	t.Run("EconomyUsers", testEconomyUsersQueryDeleteAll)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsSliceDeleteAll)
	t.Run("EconomyLogs", testEconomyLogsSliceDeleteAll)
	t.Run("EconomyPickImages", testEconomyPickImagesSliceDeleteAll)
	t.Run("EconomyPlants", testEconomyPlantsSliceDeleteAll)
	t.Run("EconomyShopItems", testEconomyShopItemsSliceDeleteAll)
	t.Run("EconomyShopListItems", testEconomyShopListItemsSliceDeleteAll)
	t.Run("EconomyUsers", testEconomyUsersSliceDeleteAll)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsExists)
	t.Run("EconomyLogs", testEconomyLogsExists)
	t.Run("EconomyPickImages", testEconomyPickImagesExists)
	t.Run("EconomyPlants", testEconomyPlantsExists)
	t.Run("EconomyShopItems", testEconomyShopItemsExists)
	t.Run("EconomyShopListItems", testEconomyShopListItemsExists)
	t.Run("EconomyUsers", testEconomyUsersExists)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsExists)
}

func TestFind(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsFind)
	t.Run("EconomyLogs", testEconomyLogsFind)
	t.Run("EconomyPickImages", testEconomyPickImagesFind)
	t.Run("EconomyPlants", testEconomyPlantsFind)
	t.Run("EconomyShopItems", testEconomyShopItemsFind)
	t.Run("EconomyShopListItems", testEconomyShopListItemsFind)
	t.Run("EconomyUsers", testEconomyUsersFind)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsFind)
}

func TestBind(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsBind)
	t.Run("EconomyLogs", testEconomyLogsBind)
	t.Run("EconomyPickImages", testEconomyPickImagesBind)
	t.Run("EconomyPlants", testEconomyPlantsBind)
	t.Run("EconomyShopItems", testEconomyShopItemsBind)
	t.Run("EconomyShopListItems", testEconomyShopListItemsBind)
	t.Run("EconomyUsers", testEconomyUsersBind)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsBind)
}

func TestOne(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsOne)
	t.Run("EconomyLogs", testEconomyLogsOne)
	t.Run("EconomyPickImages", testEconomyPickImagesOne)
	t.Run("EconomyPlants", testEconomyPlantsOne)
	t.Run("EconomyShopItems", testEconomyShopItemsOne)
	t.Run("EconomyShopListItems", testEconomyShopListItemsOne)
	t.Run("EconomyUsers", testEconomyUsersOne)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsOne)
}

func TestAll(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsAll)
	t.Run("EconomyLogs", testEconomyLogsAll)
	t.Run("EconomyPickImages", testEconomyPickImagesAll)
	t.Run("EconomyPlants", testEconomyPlantsAll)
	t.Run("EconomyShopItems", testEconomyShopItemsAll)
	t.Run("EconomyShopListItems", testEconomyShopListItemsAll)
	t.Run("EconomyUsers", testEconomyUsersAll)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsAll)
}

func TestCount(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsCount)
	t.Run("EconomyLogs", testEconomyLogsCount)
	t.Run("EconomyPickImages", testEconomyPickImagesCount)
	t.Run("EconomyPlants", testEconomyPlantsCount)
	t.Run("EconomyShopItems", testEconomyShopItemsCount)
	t.Run("EconomyShopListItems", testEconomyShopListItemsCount)
	t.Run("EconomyUsers", testEconomyUsersCount)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsCount)
}

func TestInsert(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsInsert)
	t.Run("EconomyConfigs", testEconomyConfigsInsertWhitelist)
	t.Run("EconomyLogs", testEconomyLogsInsert)
	t.Run("EconomyLogs", testEconomyLogsInsertWhitelist)
	t.Run("EconomyPickImages", testEconomyPickImagesInsert)
	t.Run("EconomyPickImages", testEconomyPickImagesInsertWhitelist)
	t.Run("EconomyPlants", testEconomyPlantsInsert)
	t.Run("EconomyPlants", testEconomyPlantsInsertWhitelist)
	t.Run("EconomyShopItems", testEconomyShopItemsInsert)
	t.Run("EconomyShopItems", testEconomyShopItemsInsertWhitelist)
	t.Run("EconomyShopListItems", testEconomyShopListItemsInsert)
	t.Run("EconomyShopListItems", testEconomyShopListItemsInsertWhitelist)
	t.Run("EconomyUsers", testEconomyUsersInsert)
	t.Run("EconomyUsers", testEconomyUsersInsertWhitelist)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsInsert)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsReload)
	t.Run("EconomyLogs", testEconomyLogsReload)
	t.Run("EconomyPickImages", testEconomyPickImagesReload)
	t.Run("EconomyPlants", testEconomyPlantsReload)
	t.Run("EconomyShopItems", testEconomyShopItemsReload)
	t.Run("EconomyShopListItems", testEconomyShopListItemsReload)
	t.Run("EconomyUsers", testEconomyUsersReload)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsReloadAll)
	t.Run("EconomyLogs", testEconomyLogsReloadAll)
	t.Run("EconomyPickImages", testEconomyPickImagesReloadAll)
	t.Run("EconomyPlants", testEconomyPlantsReloadAll)
	t.Run("EconomyShopItems", testEconomyShopItemsReloadAll)
	t.Run("EconomyShopListItems", testEconomyShopListItemsReloadAll)
	t.Run("EconomyUsers", testEconomyUsersReloadAll)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsSelect)
	t.Run("EconomyLogs", testEconomyLogsSelect)
	t.Run("EconomyPickImages", testEconomyPickImagesSelect)
	t.Run("EconomyPlants", testEconomyPlantsSelect)
	t.Run("EconomyShopItems", testEconomyShopItemsSelect)
	t.Run("EconomyShopListItems", testEconomyShopListItemsSelect)
	t.Run("EconomyUsers", testEconomyUsersSelect)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsUpdate)
	t.Run("EconomyLogs", testEconomyLogsUpdate)
	t.Run("EconomyPickImages", testEconomyPickImagesUpdate)
	t.Run("EconomyPlants", testEconomyPlantsUpdate)
	t.Run("EconomyShopItems", testEconomyShopItemsUpdate)
	t.Run("EconomyShopListItems", testEconomyShopListItemsUpdate)
	t.Run("EconomyUsers", testEconomyUsersUpdate)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("EconomyConfigs", testEconomyConfigsSliceUpdateAll)
	t.Run("EconomyLogs", testEconomyLogsSliceUpdateAll)
	t.Run("EconomyPickImages", testEconomyPickImagesSliceUpdateAll)
	t.Run("EconomyPlants", testEconomyPlantsSliceUpdateAll)
	t.Run("EconomyShopItems", testEconomyShopItemsSliceUpdateAll)
	t.Run("EconomyShopListItems", testEconomyShopListItemsSliceUpdateAll)
	t.Run("EconomyUsers", testEconomyUsersSliceUpdateAll)
	t.Run("EconomyWaifuItems", testEconomyWaifuItemsSliceUpdateAll)
}
