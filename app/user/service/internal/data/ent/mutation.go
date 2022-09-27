// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/user/service/internal/data/ent/predicate"
	"kratos-blog/app/user/service/internal/data/ent/user"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op             Op
	typ            string
	id             *uint32
	create_time    *int64
	addcreate_time *int64
	update_time    *int64
	addupdate_time *int64
	delete_time    *int64
	adddelete_time *int64
	username       *string
	nickname       *string
	email          *string
	avatar         *string
	description    *string
	password       *string
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*User, error)
	predicates     []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uint32) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uint32, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint32{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *UserMutation) SetCreateTime(i int64) {
	m.create_time = &i
	m.addcreate_time = nil
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *UserMutation) CreateTime() (r int64, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreateTime(ctx context.Context) (v *int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// AddCreateTime adds i to the "create_time" field.
func (m *UserMutation) AddCreateTime(i int64) {
	if m.addcreate_time != nil {
		*m.addcreate_time += i
	} else {
		m.addcreate_time = &i
	}
}

// AddedCreateTime returns the value that was added to the "create_time" field in this mutation.
func (m *UserMutation) AddedCreateTime() (r int64, exists bool) {
	v := m.addcreate_time
	if v == nil {
		return
	}
	return *v, true
}

// ClearCreateTime clears the value of the "create_time" field.
func (m *UserMutation) ClearCreateTime() {
	m.create_time = nil
	m.addcreate_time = nil
	m.clearedFields[user.FieldCreateTime] = struct{}{}
}

// CreateTimeCleared returns if the "create_time" field was cleared in this mutation.
func (m *UserMutation) CreateTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldCreateTime]
	return ok
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *UserMutation) ResetCreateTime() {
	m.create_time = nil
	m.addcreate_time = nil
	delete(m.clearedFields, user.FieldCreateTime)
}

// SetUpdateTime sets the "update_time" field.
func (m *UserMutation) SetUpdateTime(i int64) {
	m.update_time = &i
	m.addupdate_time = nil
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *UserMutation) UpdateTime() (r int64, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdateTime(ctx context.Context) (v *int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// AddUpdateTime adds i to the "update_time" field.
func (m *UserMutation) AddUpdateTime(i int64) {
	if m.addupdate_time != nil {
		*m.addupdate_time += i
	} else {
		m.addupdate_time = &i
	}
}

// AddedUpdateTime returns the value that was added to the "update_time" field in this mutation.
func (m *UserMutation) AddedUpdateTime() (r int64, exists bool) {
	v := m.addupdate_time
	if v == nil {
		return
	}
	return *v, true
}

// ClearUpdateTime clears the value of the "update_time" field.
func (m *UserMutation) ClearUpdateTime() {
	m.update_time = nil
	m.addupdate_time = nil
	m.clearedFields[user.FieldUpdateTime] = struct{}{}
}

// UpdateTimeCleared returns if the "update_time" field was cleared in this mutation.
func (m *UserMutation) UpdateTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldUpdateTime]
	return ok
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *UserMutation) ResetUpdateTime() {
	m.update_time = nil
	m.addupdate_time = nil
	delete(m.clearedFields, user.FieldUpdateTime)
}

// SetDeleteTime sets the "delete_time" field.
func (m *UserMutation) SetDeleteTime(i int64) {
	m.delete_time = &i
	m.adddelete_time = nil
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *UserMutation) DeleteTime() (r int64, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDeleteTime(ctx context.Context) (v *int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// AddDeleteTime adds i to the "delete_time" field.
func (m *UserMutation) AddDeleteTime(i int64) {
	if m.adddelete_time != nil {
		*m.adddelete_time += i
	} else {
		m.adddelete_time = &i
	}
}

// AddedDeleteTime returns the value that was added to the "delete_time" field in this mutation.
func (m *UserMutation) AddedDeleteTime() (r int64, exists bool) {
	v := m.adddelete_time
	if v == nil {
		return
	}
	return *v, true
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (m *UserMutation) ClearDeleteTime() {
	m.delete_time = nil
	m.adddelete_time = nil
	m.clearedFields[user.FieldDeleteTime] = struct{}{}
}

// DeleteTimeCleared returns if the "delete_time" field was cleared in this mutation.
func (m *UserMutation) DeleteTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldDeleteTime]
	return ok
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *UserMutation) ResetDeleteTime() {
	m.delete_time = nil
	m.adddelete_time = nil
	delete(m.clearedFields, user.FieldDeleteTime)
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ClearUsername clears the value of the "username" field.
func (m *UserMutation) ClearUsername() {
	m.username = nil
	m.clearedFields[user.FieldUsername] = struct{}{}
}

// UsernameCleared returns if the "username" field was cleared in this mutation.
func (m *UserMutation) UsernameCleared() bool {
	_, ok := m.clearedFields[user.FieldUsername]
	return ok
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
	delete(m.clearedFields, user.FieldUsername)
}

// SetNickname sets the "nickname" field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickname(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ClearNickname clears the value of the "nickname" field.
func (m *UserMutation) ClearNickname() {
	m.nickname = nil
	m.clearedFields[user.FieldNickname] = struct{}{}
}

// NicknameCleared returns if the "nickname" field was cleared in this mutation.
func (m *UserMutation) NicknameCleared() bool {
	_, ok := m.clearedFields[user.FieldNickname]
	return ok
}

// ResetNickname resets all changes to the "nickname" field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
	delete(m.clearedFields, user.FieldNickname)
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *UserMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[user.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *UserMutation) EmailCleared() bool {
	_, ok := m.clearedFields[user.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, user.FieldEmail)
}

// SetAvatar sets the "avatar" field.
func (m *UserMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the value of the "avatar" field in the mutation.
func (m *UserMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old "avatar" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAvatar(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatar is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ClearAvatar clears the value of the "avatar" field.
func (m *UserMutation) ClearAvatar() {
	m.avatar = nil
	m.clearedFields[user.FieldAvatar] = struct{}{}
}

// AvatarCleared returns if the "avatar" field was cleared in this mutation.
func (m *UserMutation) AvatarCleared() bool {
	_, ok := m.clearedFields[user.FieldAvatar]
	return ok
}

// ResetAvatar resets all changes to the "avatar" field.
func (m *UserMutation) ResetAvatar() {
	m.avatar = nil
	delete(m.clearedFields, user.FieldAvatar)
}

// SetDescription sets the "description" field.
func (m *UserMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *UserMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDescription(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *UserMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[user.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *UserMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[user.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *UserMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, user.FieldDescription)
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ClearPassword clears the value of the "password" field.
func (m *UserMutation) ClearPassword() {
	m.password = nil
	m.clearedFields[user.FieldPassword] = struct{}{}
}

// PasswordCleared returns if the "password" field was cleared in this mutation.
func (m *UserMutation) PasswordCleared() bool {
	_, ok := m.clearedFields[user.FieldPassword]
	return ok
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
	delete(m.clearedFields, user.FieldPassword)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.create_time != nil {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.delete_time != nil {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.avatar != nil {
		fields = append(fields, user.FieldAvatar)
	}
	if m.description != nil {
		fields = append(fields, user.FieldDescription)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreateTime:
		return m.CreateTime()
	case user.FieldUpdateTime:
		return m.UpdateTime()
	case user.FieldDeleteTime:
		return m.DeleteTime()
	case user.FieldUsername:
		return m.Username()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldEmail:
		return m.Email()
	case user.FieldAvatar:
		return m.Avatar()
	case user.FieldDescription:
		return m.Description()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case user.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case user.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldNickname:
		return m.OldNickname(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldAvatar:
		return m.OldAvatar(ctx)
	case user.FieldDescription:
		return m.OldDescription(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreateTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case user.FieldUpdateTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case user.FieldDeleteTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	case user.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addcreate_time != nil {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.addupdate_time != nil {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.adddelete_time != nil {
		fields = append(fields, user.FieldDeleteTime)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreateTime:
		return m.AddedCreateTime()
	case user.FieldUpdateTime:
		return m.AddedUpdateTime()
	case user.FieldDeleteTime:
		return m.AddedDeleteTime()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreateTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreateTime(v)
		return nil
	case user.FieldUpdateTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdateTime(v)
		return nil
	case user.FieldDeleteTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeleteTime(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldCreateTime) {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.FieldCleared(user.FieldUpdateTime) {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.FieldCleared(user.FieldDeleteTime) {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.FieldCleared(user.FieldUsername) {
		fields = append(fields, user.FieldUsername)
	}
	if m.FieldCleared(user.FieldNickname) {
		fields = append(fields, user.FieldNickname)
	}
	if m.FieldCleared(user.FieldEmail) {
		fields = append(fields, user.FieldEmail)
	}
	if m.FieldCleared(user.FieldAvatar) {
		fields = append(fields, user.FieldAvatar)
	}
	if m.FieldCleared(user.FieldDescription) {
		fields = append(fields, user.FieldDescription)
	}
	if m.FieldCleared(user.FieldPassword) {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldCreateTime:
		m.ClearCreateTime()
		return nil
	case user.FieldUpdateTime:
		m.ClearUpdateTime()
		return nil
	case user.FieldDeleteTime:
		m.ClearDeleteTime()
		return nil
	case user.FieldUsername:
		m.ClearUsername()
		return nil
	case user.FieldNickname:
		m.ClearNickname()
		return nil
	case user.FieldEmail:
		m.ClearEmail()
		return nil
	case user.FieldAvatar:
		m.ClearAvatar()
		return nil
	case user.FieldDescription:
		m.ClearDescription()
		return nil
	case user.FieldPassword:
		m.ClearPassword()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case user.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case user.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldAvatar:
		m.ResetAvatar()
		return nil
	case user.FieldDescription:
		m.ResetDescription()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
