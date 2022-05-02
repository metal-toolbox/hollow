// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Attributes", testAttributes)
	t.Run("Firmwares", testFirmwares)
	t.Run("ServerComponentTypes", testServerComponentTypes)
	t.Run("ServerComponents", testServerComponents)
	t.Run("Servers", testServers)
	t.Run("VersionedAttributes", testVersionedAttributes)
}

func TestSoftDelete(t *testing.T) {
	t.Run("Servers", testServersSoftDelete)
}

func TestQuerySoftDeleteAll(t *testing.T) {
	t.Run("Servers", testServersQuerySoftDeleteAll)
}

func TestSliceSoftDeleteAll(t *testing.T) {
	t.Run("Servers", testServersSliceSoftDeleteAll)
}

func TestDelete(t *testing.T) {
	t.Run("Attributes", testAttributesDelete)
	t.Run("Firmwares", testFirmwaresDelete)
	t.Run("ServerComponentTypes", testServerComponentTypesDelete)
	t.Run("ServerComponents", testServerComponentsDelete)
	t.Run("Servers", testServersDelete)
	t.Run("VersionedAttributes", testVersionedAttributesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Attributes", testAttributesQueryDeleteAll)
	t.Run("Firmwares", testFirmwaresQueryDeleteAll)
	t.Run("ServerComponentTypes", testServerComponentTypesQueryDeleteAll)
	t.Run("ServerComponents", testServerComponentsQueryDeleteAll)
	t.Run("Servers", testServersQueryDeleteAll)
	t.Run("VersionedAttributes", testVersionedAttributesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Attributes", testAttributesSliceDeleteAll)
	t.Run("Firmwares", testFirmwaresSliceDeleteAll)
	t.Run("ServerComponentTypes", testServerComponentTypesSliceDeleteAll)
	t.Run("ServerComponents", testServerComponentsSliceDeleteAll)
	t.Run("Servers", testServersSliceDeleteAll)
	t.Run("VersionedAttributes", testVersionedAttributesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Attributes", testAttributesExists)
	t.Run("Firmwares", testFirmwaresExists)
	t.Run("ServerComponentTypes", testServerComponentTypesExists)
	t.Run("ServerComponents", testServerComponentsExists)
	t.Run("Servers", testServersExists)
	t.Run("VersionedAttributes", testVersionedAttributesExists)
}

func TestFind(t *testing.T) {
	t.Run("Attributes", testAttributesFind)
	t.Run("Firmwares", testFirmwaresFind)
	t.Run("ServerComponentTypes", testServerComponentTypesFind)
	t.Run("ServerComponents", testServerComponentsFind)
	t.Run("Servers", testServersFind)
	t.Run("VersionedAttributes", testVersionedAttributesFind)
}

func TestBind(t *testing.T) {
	t.Run("Attributes", testAttributesBind)
	t.Run("Firmwares", testFirmwaresBind)
	t.Run("ServerComponentTypes", testServerComponentTypesBind)
	t.Run("ServerComponents", testServerComponentsBind)
	t.Run("Servers", testServersBind)
	t.Run("VersionedAttributes", testVersionedAttributesBind)
}

func TestOne(t *testing.T) {
	t.Run("Attributes", testAttributesOne)
	t.Run("Firmwares", testFirmwaresOne)
	t.Run("ServerComponentTypes", testServerComponentTypesOne)
	t.Run("ServerComponents", testServerComponentsOne)
	t.Run("Servers", testServersOne)
	t.Run("VersionedAttributes", testVersionedAttributesOne)
}

func TestAll(t *testing.T) {
	t.Run("Attributes", testAttributesAll)
	t.Run("Firmwares", testFirmwaresAll)
	t.Run("ServerComponentTypes", testServerComponentTypesAll)
	t.Run("ServerComponents", testServerComponentsAll)
	t.Run("Servers", testServersAll)
	t.Run("VersionedAttributes", testVersionedAttributesAll)
}

func TestCount(t *testing.T) {
	t.Run("Attributes", testAttributesCount)
	t.Run("Firmwares", testFirmwaresCount)
	t.Run("ServerComponentTypes", testServerComponentTypesCount)
	t.Run("ServerComponents", testServerComponentsCount)
	t.Run("Servers", testServersCount)
	t.Run("VersionedAttributes", testVersionedAttributesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Attributes", testAttributesHooks)
	t.Run("Firmwares", testFirmwaresHooks)
	t.Run("ServerComponentTypes", testServerComponentTypesHooks)
	t.Run("ServerComponents", testServerComponentsHooks)
	t.Run("Servers", testServersHooks)
	t.Run("VersionedAttributes", testVersionedAttributesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Attributes", testAttributesInsert)
	t.Run("Attributes", testAttributesInsertWhitelist)
	t.Run("Firmwares", testFirmwaresInsert)
	t.Run("Firmwares", testFirmwaresInsertWhitelist)
	t.Run("ServerComponentTypes", testServerComponentTypesInsert)
	t.Run("ServerComponentTypes", testServerComponentTypesInsertWhitelist)
	t.Run("ServerComponents", testServerComponentsInsert)
	t.Run("ServerComponents", testServerComponentsInsertWhitelist)
	t.Run("Servers", testServersInsert)
	t.Run("Servers", testServersInsertWhitelist)
	t.Run("VersionedAttributes", testVersionedAttributesInsert)
	t.Run("VersionedAttributes", testVersionedAttributesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AttributeToServerUsingServer", testAttributeToOneServerUsingServer)
	t.Run("AttributeToServerComponentUsingServerComponent", testAttributeToOneServerComponentUsingServerComponent)
	t.Run("FirmwareToServerComponentUsingComponent", testFirmwareToOneServerComponentUsingComponent)
	t.Run("ServerComponentToServerUsingServer", testServerComponentToOneServerUsingServer)
	t.Run("ServerComponentToServerComponentTypeUsingServerComponentType", testServerComponentToOneServerComponentTypeUsingServerComponentType)
	t.Run("VersionedAttributeToServerUsingServer", testVersionedAttributeToOneServerUsingServer)
	t.Run("VersionedAttributeToServerComponentUsingServerComponent", testVersionedAttributeToOneServerComponentUsingServerComponent)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ServerComponentTypeToServerComponents", testServerComponentTypeToManyServerComponents)
	t.Run("ServerComponentToAttributes", testServerComponentToManyAttributes)
	t.Run("ServerComponentToComponentFirmwares", testServerComponentToManyComponentFirmwares)
	t.Run("ServerComponentToVersionedAttributes", testServerComponentToManyVersionedAttributes)
	t.Run("ServerToAttributes", testServerToManyAttributes)
	t.Run("ServerToServerComponents", testServerToManyServerComponents)
	t.Run("ServerToVersionedAttributes", testServerToManyVersionedAttributes)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AttributeToServerUsingAttributes", testAttributeToOneSetOpServerUsingServer)
	t.Run("AttributeToServerComponentUsingAttributes", testAttributeToOneSetOpServerComponentUsingServerComponent)
	t.Run("FirmwareToServerComponentUsingComponentFirmwares", testFirmwareToOneSetOpServerComponentUsingComponent)
	t.Run("ServerComponentToServerUsingServerComponents", testServerComponentToOneSetOpServerUsingServer)
	t.Run("ServerComponentToServerComponentTypeUsingServerComponents", testServerComponentToOneSetOpServerComponentTypeUsingServerComponentType)
	t.Run("VersionedAttributeToServerUsingVersionedAttributes", testVersionedAttributeToOneSetOpServerUsingServer)
	t.Run("VersionedAttributeToServerComponentUsingVersionedAttributes", testVersionedAttributeToOneSetOpServerComponentUsingServerComponent)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("AttributeToServerUsingAttributes", testAttributeToOneRemoveOpServerUsingServer)
	t.Run("AttributeToServerComponentUsingAttributes", testAttributeToOneRemoveOpServerComponentUsingServerComponent)
	t.Run("VersionedAttributeToServerUsingVersionedAttributes", testVersionedAttributeToOneRemoveOpServerUsingServer)
	t.Run("VersionedAttributeToServerComponentUsingVersionedAttributes", testVersionedAttributeToOneRemoveOpServerComponentUsingServerComponent)
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
	t.Run("ServerComponentTypeToServerComponents", testServerComponentTypeToManyAddOpServerComponents)
	t.Run("ServerComponentToAttributes", testServerComponentToManyAddOpAttributes)
	t.Run("ServerComponentToComponentFirmwares", testServerComponentToManyAddOpComponentFirmwares)
	t.Run("ServerComponentToVersionedAttributes", testServerComponentToManyAddOpVersionedAttributes)
	t.Run("ServerToAttributes", testServerToManyAddOpAttributes)
	t.Run("ServerToServerComponents", testServerToManyAddOpServerComponents)
	t.Run("ServerToVersionedAttributes", testServerToManyAddOpVersionedAttributes)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ServerComponentToAttributes", testServerComponentToManySetOpAttributes)
	t.Run("ServerComponentToVersionedAttributes", testServerComponentToManySetOpVersionedAttributes)
	t.Run("ServerToAttributes", testServerToManySetOpAttributes)
	t.Run("ServerToVersionedAttributes", testServerToManySetOpVersionedAttributes)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ServerComponentToAttributes", testServerComponentToManyRemoveOpAttributes)
	t.Run("ServerComponentToVersionedAttributes", testServerComponentToManyRemoveOpVersionedAttributes)
	t.Run("ServerToAttributes", testServerToManyRemoveOpAttributes)
	t.Run("ServerToVersionedAttributes", testServerToManyRemoveOpVersionedAttributes)
}

func TestReload(t *testing.T) {
	t.Run("Attributes", testAttributesReload)
	t.Run("Firmwares", testFirmwaresReload)
	t.Run("ServerComponentTypes", testServerComponentTypesReload)
	t.Run("ServerComponents", testServerComponentsReload)
	t.Run("Servers", testServersReload)
	t.Run("VersionedAttributes", testVersionedAttributesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Attributes", testAttributesReloadAll)
	t.Run("Firmwares", testFirmwaresReloadAll)
	t.Run("ServerComponentTypes", testServerComponentTypesReloadAll)
	t.Run("ServerComponents", testServerComponentsReloadAll)
	t.Run("Servers", testServersReloadAll)
	t.Run("VersionedAttributes", testVersionedAttributesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Attributes", testAttributesSelect)
	t.Run("Firmwares", testFirmwaresSelect)
	t.Run("ServerComponentTypes", testServerComponentTypesSelect)
	t.Run("ServerComponents", testServerComponentsSelect)
	t.Run("Servers", testServersSelect)
	t.Run("VersionedAttributes", testVersionedAttributesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Attributes", testAttributesUpdate)
	t.Run("Firmwares", testFirmwaresUpdate)
	t.Run("ServerComponentTypes", testServerComponentTypesUpdate)
	t.Run("ServerComponents", testServerComponentsUpdate)
	t.Run("Servers", testServersUpdate)
	t.Run("VersionedAttributes", testVersionedAttributesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Attributes", testAttributesSliceUpdateAll)
	t.Run("Firmwares", testFirmwaresSliceUpdateAll)
	t.Run("ServerComponentTypes", testServerComponentTypesSliceUpdateAll)
	t.Run("ServerComponents", testServerComponentsSliceUpdateAll)
	t.Run("Servers", testServersSliceUpdateAll)
	t.Run("VersionedAttributes", testVersionedAttributesSliceUpdateAll)
}
