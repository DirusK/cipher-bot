// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cipher-bot/ent/request"
	"cipher-bot/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetTelegramID sets the "telegram_id" field.
func (uc *UserCreate) SetTelegramID(i int64) *UserCreate {
	uc.mutation.SetTelegramID(i)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetLanguage sets the "language" field.
func (uc *UserCreate) SetLanguage(s string) *UserCreate {
	uc.mutation.SetLanguage(s)
	return uc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (uc *UserCreate) SetNillableLanguage(s *string) *UserCreate {
	if s != nil {
		uc.SetLanguage(*s)
	}
	return uc
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (uc *UserCreate) AddRequestIDs(ids ...int) *UserCreate {
	uc.mutation.AddRequestIDs(ids...)
	return uc
}

// AddRequests adds the "requests" edges to the Request entity.
func (uc *UserCreate) AddRequests(r ...*Request) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRequestIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Language(); !ok {
		v := user.DefaultLanguage
		uc.mutation.SetLanguage(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.TelegramID(); !ok {
		return &ValidationError{Name: "telegram_id", err: errors.New(`ent: missing required field "User.telegram_id"`)}
	}
	if _, ok := uc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "User.first_name"`)}
	}
	if _, ok := uc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "User.last_name"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := uc.mutation.TelegramID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramID,
		})
		_node.TelegramID = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := uc.mutation.Language(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLanguage,
		})
		_node.Language = value
	}
	if nodes := uc.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestsTable,
			Columns: []string{user.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsert) SetCreatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreatedAt() *UserUpsert {
	u.SetExcluded(user.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// SetUsername sets the "username" field.
func (u *UserUpsert) SetUsername(v string) *UserUpsert {
	u.Set(user.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsert) UpdateUsername() *UserUpsert {
	u.SetExcluded(user.FieldUsername)
	return u
}

// SetTelegramID sets the "telegram_id" field.
func (u *UserUpsert) SetTelegramID(v int64) *UserUpsert {
	u.Set(user.FieldTelegramID, v)
	return u
}

// UpdateTelegramID sets the "telegram_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateTelegramID() *UserUpsert {
	u.SetExcluded(user.FieldTelegramID)
	return u
}

// AddTelegramID adds v to the "telegram_id" field.
func (u *UserUpsert) AddTelegramID(v int64) *UserUpsert {
	u.Add(user.FieldTelegramID, v)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsert) SetFirstName(v string) *UserUpsert {
	u.Set(user.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsert) UpdateFirstName() *UserUpsert {
	u.SetExcluded(user.FieldFirstName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *UserUpsert) SetLastName(v string) *UserUpsert {
	u.Set(user.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsert) UpdateLastName() *UserUpsert {
	u.SetExcluded(user.FieldLastName)
	return u
}

// SetLanguage sets the "language" field.
func (u *UserUpsert) SetLanguage(v string) *UserUpsert {
	u.Set(user.FieldLanguage, v)
	return u
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *UserUpsert) UpdateLanguage() *UserUpsert {
	u.SetExcluded(user.FieldLanguage)
	return u
}

// ClearLanguage clears the value of the "language" field.
func (u *UserUpsert) ClearLanguage() *UserUpsert {
	u.SetNull(user.FieldLanguage)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(user.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.User.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertOne) SetCreatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertOne) SetUsername(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUsername() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetTelegramID sets the "telegram_id" field.
func (u *UserUpsertOne) SetTelegramID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetTelegramID(v)
	})
}

// AddTelegramID adds v to the "telegram_id" field.
func (u *UserUpsertOne) AddTelegramID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddTelegramID(v)
	})
}

// UpdateTelegramID sets the "telegram_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateTelegramID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateTelegramID()
	})
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsertOne) SetFirstName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateFirstName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *UserUpsertOne) SetLastName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLastName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastName()
	})
}

// SetLanguage sets the "language" field.
func (u *UserUpsertOne) SetLanguage(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLanguage() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLanguage()
	})
}

// ClearLanguage clears the value of the "language" field.
func (u *UserUpsertOne) ClearLanguage() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearLanguage()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(user.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertBulk) SetCreatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertBulk) SetUsername(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUsername() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetTelegramID sets the "telegram_id" field.
func (u *UserUpsertBulk) SetTelegramID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetTelegramID(v)
	})
}

// AddTelegramID adds v to the "telegram_id" field.
func (u *UserUpsertBulk) AddTelegramID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddTelegramID(v)
	})
}

// UpdateTelegramID sets the "telegram_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateTelegramID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateTelegramID()
	})
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsertBulk) SetFirstName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateFirstName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *UserUpsertBulk) SetLastName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLastName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastName()
	})
}

// SetLanguage sets the "language" field.
func (u *UserUpsertBulk) SetLanguage(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLanguage() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLanguage()
	})
}

// ClearLanguage clears the value of the "language" field.
func (u *UserUpsertBulk) ClearLanguage() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearLanguage()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
