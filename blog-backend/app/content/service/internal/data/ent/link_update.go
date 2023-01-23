// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/content/service/internal/data/ent/link"
	"kratos-blog/app/content/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks     []Hook
	mutation  *LinkMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdateTime sets the "update_time" field.
func (lu *LinkUpdate) SetUpdateTime(i int64) *LinkUpdate {
	lu.mutation.ResetUpdateTime()
	lu.mutation.SetUpdateTime(i)
	return lu
}

// AddUpdateTime adds i to the "update_time" field.
func (lu *LinkUpdate) AddUpdateTime(i int64) *LinkUpdate {
	lu.mutation.AddUpdateTime(i)
	return lu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (lu *LinkUpdate) ClearUpdateTime() *LinkUpdate {
	lu.mutation.ClearUpdateTime()
	return lu
}

// SetDeleteTime sets the "delete_time" field.
func (lu *LinkUpdate) SetDeleteTime(i int64) *LinkUpdate {
	lu.mutation.ResetDeleteTime()
	lu.mutation.SetDeleteTime(i)
	return lu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableDeleteTime(i *int64) *LinkUpdate {
	if i != nil {
		lu.SetDeleteTime(*i)
	}
	return lu
}

// AddDeleteTime adds i to the "delete_time" field.
func (lu *LinkUpdate) AddDeleteTime(i int64) *LinkUpdate {
	lu.mutation.AddDeleteTime(i)
	return lu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (lu *LinkUpdate) ClearDeleteTime() *LinkUpdate {
	lu.mutation.ClearDeleteTime()
	return lu
}

// SetName sets the "name" field.
func (lu *LinkUpdate) SetName(s string) *LinkUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableName(s *string) *LinkUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// ClearName clears the value of the "name" field.
func (lu *LinkUpdate) ClearName() *LinkUpdate {
	lu.mutation.ClearName()
	return lu
}

// SetURL sets the "url" field.
func (lu *LinkUpdate) SetURL(s string) *LinkUpdate {
	lu.mutation.SetURL(s)
	return lu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableURL(s *string) *LinkUpdate {
	if s != nil {
		lu.SetURL(*s)
	}
	return lu
}

// ClearURL clears the value of the "url" field.
func (lu *LinkUpdate) ClearURL() *LinkUpdate {
	lu.mutation.ClearURL()
	return lu
}

// SetLogo sets the "logo" field.
func (lu *LinkUpdate) SetLogo(s string) *LinkUpdate {
	lu.mutation.SetLogo(s)
	return lu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableLogo(s *string) *LinkUpdate {
	if s != nil {
		lu.SetLogo(*s)
	}
	return lu
}

// ClearLogo clears the value of the "logo" field.
func (lu *LinkUpdate) ClearLogo() *LinkUpdate {
	lu.mutation.ClearLogo()
	return lu
}

// SetDescription sets the "description" field.
func (lu *LinkUpdate) SetDescription(s string) *LinkUpdate {
	lu.mutation.SetDescription(s)
	return lu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableDescription(s *string) *LinkUpdate {
	if s != nil {
		lu.SetDescription(*s)
	}
	return lu
}

// ClearDescription clears the value of the "description" field.
func (lu *LinkUpdate) ClearDescription() *LinkUpdate {
	lu.mutation.ClearDescription()
	return lu
}

// SetTeam sets the "team" field.
func (lu *LinkUpdate) SetTeam(s string) *LinkUpdate {
	lu.mutation.SetTeam(s)
	return lu
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableTeam(s *string) *LinkUpdate {
	if s != nil {
		lu.SetTeam(*s)
	}
	return lu
}

// ClearTeam clears the value of the "team" field.
func (lu *LinkUpdate) ClearTeam() *LinkUpdate {
	lu.mutation.ClearTeam()
	return lu
}

// SetPriority sets the "priority" field.
func (lu *LinkUpdate) SetPriority(i int32) *LinkUpdate {
	lu.mutation.ResetPriority()
	lu.mutation.SetPriority(i)
	return lu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (lu *LinkUpdate) SetNillablePriority(i *int32) *LinkUpdate {
	if i != nil {
		lu.SetPriority(*i)
	}
	return lu
}

// AddPriority adds i to the "priority" field.
func (lu *LinkUpdate) AddPriority(i int32) *LinkUpdate {
	lu.mutation.AddPriority(i)
	return lu
}

// ClearPriority clears the value of the "priority" field.
func (lu *LinkUpdate) ClearPriority() *LinkUpdate {
	lu.mutation.ClearPriority()
	return lu
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks[int, LinkMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LinkUpdate) defaults() {
	if _, ok := lu.mutation.UpdateTime(); !ok && !lu.mutation.UpdateTimeCleared() {
		v := link.UpdateDefaultUpdateTime()
		lu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LinkUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := link.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Link.name": %w`, err)}
		}
	}
	if v, ok := lu.mutation.URL(); ok {
		if err := link.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Link.url": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Logo(); ok {
		if err := link.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "Link.logo": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *LinkUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LinkUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: link.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lu.mutation.CreateTimeCleared() {
		_spec.ClearField(link.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := lu.mutation.UpdateTime(); ok {
		_spec.SetField(link.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := lu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(link.FieldUpdateTime, field.TypeInt64, value)
	}
	if lu.mutation.UpdateTimeCleared() {
		_spec.ClearField(link.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := lu.mutation.DeleteTime(); ok {
		_spec.SetField(link.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := lu.mutation.AddedDeleteTime(); ok {
		_spec.AddField(link.FieldDeleteTime, field.TypeInt64, value)
	}
	if lu.mutation.DeleteTimeCleared() {
		_spec.ClearField(link.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(link.FieldName, field.TypeString, value)
	}
	if lu.mutation.NameCleared() {
		_spec.ClearField(link.FieldName, field.TypeString)
	}
	if value, ok := lu.mutation.URL(); ok {
		_spec.SetField(link.FieldURL, field.TypeString, value)
	}
	if lu.mutation.URLCleared() {
		_spec.ClearField(link.FieldURL, field.TypeString)
	}
	if value, ok := lu.mutation.Logo(); ok {
		_spec.SetField(link.FieldLogo, field.TypeString, value)
	}
	if lu.mutation.LogoCleared() {
		_spec.ClearField(link.FieldLogo, field.TypeString)
	}
	if value, ok := lu.mutation.Description(); ok {
		_spec.SetField(link.FieldDescription, field.TypeString, value)
	}
	if lu.mutation.DescriptionCleared() {
		_spec.ClearField(link.FieldDescription, field.TypeString)
	}
	if value, ok := lu.mutation.Team(); ok {
		_spec.SetField(link.FieldTeam, field.TypeString, value)
	}
	if lu.mutation.TeamCleared() {
		_spec.ClearField(link.FieldTeam, field.TypeString)
	}
	if value, ok := lu.mutation.Priority(); ok {
		_spec.SetField(link.FieldPriority, field.TypeInt32, value)
	}
	if value, ok := lu.mutation.AddedPriority(); ok {
		_spec.AddField(link.FieldPriority, field.TypeInt32, value)
	}
	if lu.mutation.PriorityCleared() {
		_spec.ClearField(link.FieldPriority, field.TypeInt32)
	}
	_spec.AddModifiers(lu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LinkMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (luo *LinkUpdateOne) SetUpdateTime(i int64) *LinkUpdateOne {
	luo.mutation.ResetUpdateTime()
	luo.mutation.SetUpdateTime(i)
	return luo
}

// AddUpdateTime adds i to the "update_time" field.
func (luo *LinkUpdateOne) AddUpdateTime(i int64) *LinkUpdateOne {
	luo.mutation.AddUpdateTime(i)
	return luo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (luo *LinkUpdateOne) ClearUpdateTime() *LinkUpdateOne {
	luo.mutation.ClearUpdateTime()
	return luo
}

// SetDeleteTime sets the "delete_time" field.
func (luo *LinkUpdateOne) SetDeleteTime(i int64) *LinkUpdateOne {
	luo.mutation.ResetDeleteTime()
	luo.mutation.SetDeleteTime(i)
	return luo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableDeleteTime(i *int64) *LinkUpdateOne {
	if i != nil {
		luo.SetDeleteTime(*i)
	}
	return luo
}

// AddDeleteTime adds i to the "delete_time" field.
func (luo *LinkUpdateOne) AddDeleteTime(i int64) *LinkUpdateOne {
	luo.mutation.AddDeleteTime(i)
	return luo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (luo *LinkUpdateOne) ClearDeleteTime() *LinkUpdateOne {
	luo.mutation.ClearDeleteTime()
	return luo
}

// SetName sets the "name" field.
func (luo *LinkUpdateOne) SetName(s string) *LinkUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableName(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// ClearName clears the value of the "name" field.
func (luo *LinkUpdateOne) ClearName() *LinkUpdateOne {
	luo.mutation.ClearName()
	return luo
}

// SetURL sets the "url" field.
func (luo *LinkUpdateOne) SetURL(s string) *LinkUpdateOne {
	luo.mutation.SetURL(s)
	return luo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableURL(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetURL(*s)
	}
	return luo
}

// ClearURL clears the value of the "url" field.
func (luo *LinkUpdateOne) ClearURL() *LinkUpdateOne {
	luo.mutation.ClearURL()
	return luo
}

// SetLogo sets the "logo" field.
func (luo *LinkUpdateOne) SetLogo(s string) *LinkUpdateOne {
	luo.mutation.SetLogo(s)
	return luo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableLogo(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetLogo(*s)
	}
	return luo
}

// ClearLogo clears the value of the "logo" field.
func (luo *LinkUpdateOne) ClearLogo() *LinkUpdateOne {
	luo.mutation.ClearLogo()
	return luo
}

// SetDescription sets the "description" field.
func (luo *LinkUpdateOne) SetDescription(s string) *LinkUpdateOne {
	luo.mutation.SetDescription(s)
	return luo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableDescription(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetDescription(*s)
	}
	return luo
}

// ClearDescription clears the value of the "description" field.
func (luo *LinkUpdateOne) ClearDescription() *LinkUpdateOne {
	luo.mutation.ClearDescription()
	return luo
}

// SetTeam sets the "team" field.
func (luo *LinkUpdateOne) SetTeam(s string) *LinkUpdateOne {
	luo.mutation.SetTeam(s)
	return luo
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableTeam(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetTeam(*s)
	}
	return luo
}

// ClearTeam clears the value of the "team" field.
func (luo *LinkUpdateOne) ClearTeam() *LinkUpdateOne {
	luo.mutation.ClearTeam()
	return luo
}

// SetPriority sets the "priority" field.
func (luo *LinkUpdateOne) SetPriority(i int32) *LinkUpdateOne {
	luo.mutation.ResetPriority()
	luo.mutation.SetPriority(i)
	return luo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillablePriority(i *int32) *LinkUpdateOne {
	if i != nil {
		luo.SetPriority(*i)
	}
	return luo
}

// AddPriority adds i to the "priority" field.
func (luo *LinkUpdateOne) AddPriority(i int32) *LinkUpdateOne {
	luo.mutation.AddPriority(i)
	return luo
}

// ClearPriority clears the value of the "priority" field.
func (luo *LinkUpdateOne) ClearPriority() *LinkUpdateOne {
	luo.mutation.ClearPriority()
	return luo
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	luo.defaults()
	return withHooks[*Link, LinkMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LinkUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdateTime(); !ok && !luo.mutation.UpdateTimeCleared() {
		v := link.UpdateDefaultUpdateTime()
		luo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LinkUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := link.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Link.name": %w`, err)}
		}
	}
	if v, ok := luo.mutation.URL(); ok {
		if err := link.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Link.url": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Logo(); ok {
		if err := link.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "Link.logo": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *LinkUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LinkUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: link.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if luo.mutation.CreateTimeCleared() {
		_spec.ClearField(link.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := luo.mutation.UpdateTime(); ok {
		_spec.SetField(link.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := luo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(link.FieldUpdateTime, field.TypeInt64, value)
	}
	if luo.mutation.UpdateTimeCleared() {
		_spec.ClearField(link.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := luo.mutation.DeleteTime(); ok {
		_spec.SetField(link.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := luo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(link.FieldDeleteTime, field.TypeInt64, value)
	}
	if luo.mutation.DeleteTimeCleared() {
		_spec.ClearField(link.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(link.FieldName, field.TypeString, value)
	}
	if luo.mutation.NameCleared() {
		_spec.ClearField(link.FieldName, field.TypeString)
	}
	if value, ok := luo.mutation.URL(); ok {
		_spec.SetField(link.FieldURL, field.TypeString, value)
	}
	if luo.mutation.URLCleared() {
		_spec.ClearField(link.FieldURL, field.TypeString)
	}
	if value, ok := luo.mutation.Logo(); ok {
		_spec.SetField(link.FieldLogo, field.TypeString, value)
	}
	if luo.mutation.LogoCleared() {
		_spec.ClearField(link.FieldLogo, field.TypeString)
	}
	if value, ok := luo.mutation.Description(); ok {
		_spec.SetField(link.FieldDescription, field.TypeString, value)
	}
	if luo.mutation.DescriptionCleared() {
		_spec.ClearField(link.FieldDescription, field.TypeString)
	}
	if value, ok := luo.mutation.Team(); ok {
		_spec.SetField(link.FieldTeam, field.TypeString, value)
	}
	if luo.mutation.TeamCleared() {
		_spec.ClearField(link.FieldTeam, field.TypeString)
	}
	if value, ok := luo.mutation.Priority(); ok {
		_spec.SetField(link.FieldPriority, field.TypeInt32, value)
	}
	if value, ok := luo.mutation.AddedPriority(); ok {
		_spec.AddField(link.FieldPriority, field.TypeInt32, value)
	}
	if luo.mutation.PriorityCleared() {
		_spec.ClearField(link.FieldPriority, field.TypeInt32)
	}
	_spec.AddModifiers(luo.modifiers...)
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
