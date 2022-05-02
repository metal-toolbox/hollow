// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Firmware is an object representing the database table.
type Firmware struct {
	ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ComponentID string      `boil:"component_id" json:"component_id" toml:"component_id" yaml:"component_id"`
	Vendor      null.String `boil:"vendor" json:"vendor,omitempty" toml:"vendor" yaml:"vendor,omitempty"`
	Model       null.String `boil:"model" json:"model,omitempty" toml:"model" yaml:"model,omitempty"`
	Filename    null.String `boil:"filename" json:"filename,omitempty" toml:"filename" yaml:"filename,omitempty"`
	Version     null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	Utility     null.String `boil:"utility" json:"utility,omitempty" toml:"utility" yaml:"utility,omitempty"`
	Sha         null.String `boil:"sha" json:"sha,omitempty" toml:"sha" yaml:"sha,omitempty"`
	UpstreamURL null.String `boil:"upstream_url" json:"upstream_url,omitempty" toml:"upstream_url" yaml:"upstream_url,omitempty"`

	R *firmwareR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L firmwareL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FirmwareColumns = struct {
	ID          string
	ComponentID string
	Vendor      string
	Model       string
	Filename    string
	Version     string
	Utility     string
	Sha         string
	UpstreamURL string
}{
	ID:          "id",
	ComponentID: "component_id",
	Vendor:      "vendor",
	Model:       "model",
	Filename:    "filename",
	Version:     "version",
	Utility:     "utility",
	Sha:         "sha",
	UpstreamURL: "upstream_url",
}

var FirmwareTableColumns = struct {
	ID          string
	ComponentID string
	Vendor      string
	Model       string
	Filename    string
	Version     string
	Utility     string
	Sha         string
	UpstreamURL string
}{
	ID:          "firmwares.id",
	ComponentID: "firmwares.component_id",
	Vendor:      "firmwares.vendor",
	Model:       "firmwares.model",
	Filename:    "firmwares.filename",
	Version:     "firmwares.version",
	Utility:     "firmwares.utility",
	Sha:         "firmwares.sha",
	UpstreamURL: "firmwares.upstream_url",
}

// Generated where

var FirmwareWhere = struct {
	ID          whereHelperstring
	ComponentID whereHelperstring
	Vendor      whereHelpernull_String
	Model       whereHelpernull_String
	Filename    whereHelpernull_String
	Version     whereHelpernull_String
	Utility     whereHelpernull_String
	Sha         whereHelpernull_String
	UpstreamURL whereHelpernull_String
}{
	ID:          whereHelperstring{field: "\"firmwares\".\"id\""},
	ComponentID: whereHelperstring{field: "\"firmwares\".\"component_id\""},
	Vendor:      whereHelpernull_String{field: "\"firmwares\".\"vendor\""},
	Model:       whereHelpernull_String{field: "\"firmwares\".\"model\""},
	Filename:    whereHelpernull_String{field: "\"firmwares\".\"filename\""},
	Version:     whereHelpernull_String{field: "\"firmwares\".\"version\""},
	Utility:     whereHelpernull_String{field: "\"firmwares\".\"utility\""},
	Sha:         whereHelpernull_String{field: "\"firmwares\".\"sha\""},
	UpstreamURL: whereHelpernull_String{field: "\"firmwares\".\"upstream_url\""},
}

// FirmwareRels is where relationship names are stored.
var FirmwareRels = struct {
	Component string
}{
	Component: "Component",
}

// firmwareR is where relationships are stored.
type firmwareR struct {
	Component *ServerComponent `boil:"Component" json:"Component" toml:"Component" yaml:"Component"`
}

// NewStruct creates a new relationship struct
func (*firmwareR) NewStruct() *firmwareR {
	return &firmwareR{}
}

// firmwareL is where Load methods for each relationship are stored.
type firmwareL struct{}

var (
	firmwareAllColumns            = []string{"id", "component_id", "vendor", "model", "filename", "version", "utility", "sha", "upstream_url"}
	firmwareColumnsWithoutDefault = []string{"component_id"}
	firmwareColumnsWithDefault    = []string{"id", "vendor", "model", "filename", "version", "utility", "sha", "upstream_url"}
	firmwarePrimaryKeyColumns     = []string{"id"}
)

type (
	// FirmwareSlice is an alias for a slice of pointers to Firmware.
	// This should almost always be used instead of []Firmware.
	FirmwareSlice []*Firmware
	// FirmwareHook is the signature for custom Firmware hook methods
	FirmwareHook func(context.Context, boil.ContextExecutor, *Firmware) error

	firmwareQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	firmwareType                 = reflect.TypeOf(&Firmware{})
	firmwareMapping              = queries.MakeStructMapping(firmwareType)
	firmwarePrimaryKeyMapping, _ = queries.BindMapping(firmwareType, firmwareMapping, firmwarePrimaryKeyColumns)
	firmwareInsertCacheMut       sync.RWMutex
	firmwareInsertCache          = make(map[string]insertCache)
	firmwareUpdateCacheMut       sync.RWMutex
	firmwareUpdateCache          = make(map[string]updateCache)
	firmwareUpsertCacheMut       sync.RWMutex
	firmwareUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var firmwareBeforeInsertHooks []FirmwareHook
var firmwareBeforeUpdateHooks []FirmwareHook
var firmwareBeforeDeleteHooks []FirmwareHook
var firmwareBeforeUpsertHooks []FirmwareHook

var firmwareAfterInsertHooks []FirmwareHook
var firmwareAfterSelectHooks []FirmwareHook
var firmwareAfterUpdateHooks []FirmwareHook
var firmwareAfterDeleteHooks []FirmwareHook
var firmwareAfterUpsertHooks []FirmwareHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Firmware) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Firmware) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Firmware) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Firmware) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Firmware) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Firmware) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Firmware) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Firmware) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Firmware) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmwareAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFirmwareHook registers your hook function for all future operations.
func AddFirmwareHook(hookPoint boil.HookPoint, firmwareHook FirmwareHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		firmwareBeforeInsertHooks = append(firmwareBeforeInsertHooks, firmwareHook)
	case boil.BeforeUpdateHook:
		firmwareBeforeUpdateHooks = append(firmwareBeforeUpdateHooks, firmwareHook)
	case boil.BeforeDeleteHook:
		firmwareBeforeDeleteHooks = append(firmwareBeforeDeleteHooks, firmwareHook)
	case boil.BeforeUpsertHook:
		firmwareBeforeUpsertHooks = append(firmwareBeforeUpsertHooks, firmwareHook)
	case boil.AfterInsertHook:
		firmwareAfterInsertHooks = append(firmwareAfterInsertHooks, firmwareHook)
	case boil.AfterSelectHook:
		firmwareAfterSelectHooks = append(firmwareAfterSelectHooks, firmwareHook)
	case boil.AfterUpdateHook:
		firmwareAfterUpdateHooks = append(firmwareAfterUpdateHooks, firmwareHook)
	case boil.AfterDeleteHook:
		firmwareAfterDeleteHooks = append(firmwareAfterDeleteHooks, firmwareHook)
	case boil.AfterUpsertHook:
		firmwareAfterUpsertHooks = append(firmwareAfterUpsertHooks, firmwareHook)
	}
}

// One returns a single firmware record from the query.
func (q firmwareQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Firmware, error) {
	o := &Firmware{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for firmwares")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Firmware records from the query.
func (q firmwareQuery) All(ctx context.Context, exec boil.ContextExecutor) (FirmwareSlice, error) {
	var o []*Firmware

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Firmware slice")
	}

	if len(firmwareAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Firmware records in the query.
func (q firmwareQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count firmwares rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q firmwareQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if firmwares exists")
	}

	return count > 0, nil
}

// Component pointed to by the foreign key.
func (o *Firmware) Component(mods ...qm.QueryMod) serverComponentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ComponentID),
	}

	queryMods = append(queryMods, mods...)

	query := ServerComponents(queryMods...)
	queries.SetFrom(query.Query, "\"server_components\"")

	return query
}

// LoadComponent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (firmwareL) LoadComponent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFirmware interface{}, mods queries.Applicator) error {
	var slice []*Firmware
	var object *Firmware

	if singular {
		object = maybeFirmware.(*Firmware)
	} else {
		slice = *maybeFirmware.(*[]*Firmware)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &firmwareR{}
		}
		args = append(args, object.ComponentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &firmwareR{}
			}

			for _, a := range args {
				if a == obj.ComponentID {
					continue Outer
				}
			}

			args = append(args, obj.ComponentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`server_components`),
		qm.WhereIn(`server_components.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServerComponent")
	}

	var resultSlice []*ServerComponent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServerComponent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for server_components")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for server_components")
	}

	if len(firmwareAfterSelectHooks) != 0 {
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
		object.R.Component = foreign
		if foreign.R == nil {
			foreign.R = &serverComponentR{}
		}
		foreign.R.ComponentFirmwares = append(foreign.R.ComponentFirmwares, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ComponentID == foreign.ID {
				local.R.Component = foreign
				if foreign.R == nil {
					foreign.R = &serverComponentR{}
				}
				foreign.R.ComponentFirmwares = append(foreign.R.ComponentFirmwares, local)
				break
			}
		}
	}

	return nil
}

// SetComponent of the firmware to the related item.
// Sets o.R.Component to related.
// Adds o to related.R.ComponentFirmwares.
func (o *Firmware) SetComponent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServerComponent) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"firmwares\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"component_id"}),
		strmangle.WhereClause("\"", "\"", 2, firmwarePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ComponentID = related.ID
	if o.R == nil {
		o.R = &firmwareR{
			Component: related,
		}
	} else {
		o.R.Component = related
	}

	if related.R == nil {
		related.R = &serverComponentR{
			ComponentFirmwares: FirmwareSlice{o},
		}
	} else {
		related.R.ComponentFirmwares = append(related.R.ComponentFirmwares, o)
	}

	return nil
}

// Firmwares retrieves all the records using an executor.
func Firmwares(mods ...qm.QueryMod) firmwareQuery {
	mods = append(mods, qm.From("\"firmwares\""))
	return firmwareQuery{NewQuery(mods...)}
}

// FindFirmware retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFirmware(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Firmware, error) {
	firmwareObj := &Firmware{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"firmwares\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, firmwareObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from firmwares")
	}

	if err = firmwareObj.doAfterSelectHooks(ctx, exec); err != nil {
		return firmwareObj, err
	}

	return firmwareObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Firmware) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no firmwares provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(firmwareColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	firmwareInsertCacheMut.RLock()
	cache, cached := firmwareInsertCache[key]
	firmwareInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			firmwareAllColumns,
			firmwareColumnsWithDefault,
			firmwareColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(firmwareType, firmwareMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(firmwareType, firmwareMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"firmwares\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"firmwares\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into firmwares")
	}

	if !cached {
		firmwareInsertCacheMut.Lock()
		firmwareInsertCache[key] = cache
		firmwareInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Firmware.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Firmware) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	firmwareUpdateCacheMut.RLock()
	cache, cached := firmwareUpdateCache[key]
	firmwareUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			firmwareAllColumns,
			firmwarePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update firmwares, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"firmwares\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, firmwarePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(firmwareType, firmwareMapping, append(wl, firmwarePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update firmwares row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for firmwares")
	}

	if !cached {
		firmwareUpdateCacheMut.Lock()
		firmwareUpdateCache[key] = cache
		firmwareUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q firmwareQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for firmwares")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for firmwares")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FirmwareSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmwarePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"firmwares\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, firmwarePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in firmware slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all firmware")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Firmware) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no firmwares provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(firmwareColumnsWithDefault, o)

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

	firmwareUpsertCacheMut.RLock()
	cache, cached := firmwareUpsertCache[key]
	firmwareUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			firmwareAllColumns,
			firmwareColumnsWithDefault,
			firmwareColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			firmwareAllColumns,
			firmwarePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert firmwares, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(firmwarePrimaryKeyColumns))
			copy(conflict, firmwarePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"firmwares\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(firmwareType, firmwareMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(firmwareType, firmwareMapping, ret)
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
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert firmwares")
	}

	if !cached {
		firmwareUpsertCacheMut.Lock()
		firmwareUpsertCache[key] = cache
		firmwareUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Firmware record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Firmware) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Firmware provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), firmwarePrimaryKeyMapping)
	sql := "DELETE FROM \"firmwares\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from firmwares")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for firmwares")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q firmwareQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no firmwareQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from firmwares")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for firmwares")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FirmwareSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(firmwareBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmwarePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"firmwares\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, firmwarePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from firmware slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for firmwares")
	}

	if len(firmwareAfterDeleteHooks) != 0 {
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
func (o *Firmware) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFirmware(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FirmwareSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FirmwareSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmwarePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"firmwares\".* FROM \"firmwares\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, firmwarePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FirmwareSlice")
	}

	*o = slice

	return nil
}

// FirmwareExists checks if the Firmware row exists.
func FirmwareExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"firmwares\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if firmwares exists")
	}

	return exists, nil
}
