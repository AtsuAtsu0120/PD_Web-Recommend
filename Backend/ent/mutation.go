// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Backend/ent/predicate"
	"Backend/ent/website"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser    = "User"
	TypeWebsite = "Website"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
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
func withUserID(id int) userOption {
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

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
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
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
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

// WebsiteMutation represents an operation that mutates the Website nodes in the graph.
type WebsiteMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	filePath      *string
	url           *string
	bright        *float64
	addbright     *float64
	flashy        *float64
	addflashy     *float64
	adult         *float64
	addadult      *float64
	smart         *float64
	addsmart      *float64
	beautiful     *float64
	addbeautiful  *float64
	like          *float64
	addlike       *float64
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Website, error)
	predicates    []predicate.Website
}

var _ ent.Mutation = (*WebsiteMutation)(nil)

// websiteOption allows management of the mutation configuration using functional options.
type websiteOption func(*WebsiteMutation)

// newWebsiteMutation creates new mutation for the Website entity.
func newWebsiteMutation(c config, op Op, opts ...websiteOption) *WebsiteMutation {
	m := &WebsiteMutation{
		config:        c,
		op:            op,
		typ:           TypeWebsite,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWebsiteID sets the ID field of the mutation.
func withWebsiteID(id int) websiteOption {
	return func(m *WebsiteMutation) {
		var (
			err   error
			once  sync.Once
			value *Website
		)
		m.oldValue = func(ctx context.Context) (*Website, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Website.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWebsite sets the old Website of the mutation.
func withWebsite(node *Website) websiteOption {
	return func(m *WebsiteMutation) {
		m.oldValue = func(context.Context) (*Website, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WebsiteMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WebsiteMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WebsiteMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WebsiteMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Website.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *WebsiteMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *WebsiteMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *WebsiteMutation) ResetName() {
	m.name = nil
}

// SetFilePath sets the "filePath" field.
func (m *WebsiteMutation) SetFilePath(s string) {
	m.filePath = &s
}

// FilePath returns the value of the "filePath" field in the mutation.
func (m *WebsiteMutation) FilePath() (r string, exists bool) {
	v := m.filePath
	if v == nil {
		return
	}
	return *v, true
}

// OldFilePath returns the old "filePath" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldFilePath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFilePath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFilePath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFilePath: %w", err)
	}
	return oldValue.FilePath, nil
}

// ResetFilePath resets all changes to the "filePath" field.
func (m *WebsiteMutation) ResetFilePath() {
	m.filePath = nil
}

// SetURL sets the "url" field.
func (m *WebsiteMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *WebsiteMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *WebsiteMutation) ResetURL() {
	m.url = nil
}

// SetBright sets the "bright" field.
func (m *WebsiteMutation) SetBright(f float64) {
	m.bright = &f
	m.addbright = nil
}

// Bright returns the value of the "bright" field in the mutation.
func (m *WebsiteMutation) Bright() (r float64, exists bool) {
	v := m.bright
	if v == nil {
		return
	}
	return *v, true
}

// OldBright returns the old "bright" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldBright(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBright is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBright requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBright: %w", err)
	}
	return oldValue.Bright, nil
}

// AddBright adds f to the "bright" field.
func (m *WebsiteMutation) AddBright(f float64) {
	if m.addbright != nil {
		*m.addbright += f
	} else {
		m.addbright = &f
	}
}

// AddedBright returns the value that was added to the "bright" field in this mutation.
func (m *WebsiteMutation) AddedBright() (r float64, exists bool) {
	v := m.addbright
	if v == nil {
		return
	}
	return *v, true
}

// ResetBright resets all changes to the "bright" field.
func (m *WebsiteMutation) ResetBright() {
	m.bright = nil
	m.addbright = nil
}

// SetFlashy sets the "flashy" field.
func (m *WebsiteMutation) SetFlashy(f float64) {
	m.flashy = &f
	m.addflashy = nil
}

// Flashy returns the value of the "flashy" field in the mutation.
func (m *WebsiteMutation) Flashy() (r float64, exists bool) {
	v := m.flashy
	if v == nil {
		return
	}
	return *v, true
}

// OldFlashy returns the old "flashy" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldFlashy(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFlashy is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFlashy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFlashy: %w", err)
	}
	return oldValue.Flashy, nil
}

// AddFlashy adds f to the "flashy" field.
func (m *WebsiteMutation) AddFlashy(f float64) {
	if m.addflashy != nil {
		*m.addflashy += f
	} else {
		m.addflashy = &f
	}
}

// AddedFlashy returns the value that was added to the "flashy" field in this mutation.
func (m *WebsiteMutation) AddedFlashy() (r float64, exists bool) {
	v := m.addflashy
	if v == nil {
		return
	}
	return *v, true
}

// ResetFlashy resets all changes to the "flashy" field.
func (m *WebsiteMutation) ResetFlashy() {
	m.flashy = nil
	m.addflashy = nil
}

// SetAdult sets the "adult" field.
func (m *WebsiteMutation) SetAdult(f float64) {
	m.adult = &f
	m.addadult = nil
}

// Adult returns the value of the "adult" field in the mutation.
func (m *WebsiteMutation) Adult() (r float64, exists bool) {
	v := m.adult
	if v == nil {
		return
	}
	return *v, true
}

// OldAdult returns the old "adult" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldAdult(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAdult is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAdult requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAdult: %w", err)
	}
	return oldValue.Adult, nil
}

// AddAdult adds f to the "adult" field.
func (m *WebsiteMutation) AddAdult(f float64) {
	if m.addadult != nil {
		*m.addadult += f
	} else {
		m.addadult = &f
	}
}

// AddedAdult returns the value that was added to the "adult" field in this mutation.
func (m *WebsiteMutation) AddedAdult() (r float64, exists bool) {
	v := m.addadult
	if v == nil {
		return
	}
	return *v, true
}

// ResetAdult resets all changes to the "adult" field.
func (m *WebsiteMutation) ResetAdult() {
	m.adult = nil
	m.addadult = nil
}

// SetSmart sets the "smart" field.
func (m *WebsiteMutation) SetSmart(f float64) {
	m.smart = &f
	m.addsmart = nil
}

// Smart returns the value of the "smart" field in the mutation.
func (m *WebsiteMutation) Smart() (r float64, exists bool) {
	v := m.smart
	if v == nil {
		return
	}
	return *v, true
}

// OldSmart returns the old "smart" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldSmart(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSmart is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSmart requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSmart: %w", err)
	}
	return oldValue.Smart, nil
}

// AddSmart adds f to the "smart" field.
func (m *WebsiteMutation) AddSmart(f float64) {
	if m.addsmart != nil {
		*m.addsmart += f
	} else {
		m.addsmart = &f
	}
}

// AddedSmart returns the value that was added to the "smart" field in this mutation.
func (m *WebsiteMutation) AddedSmart() (r float64, exists bool) {
	v := m.addsmart
	if v == nil {
		return
	}
	return *v, true
}

// ResetSmart resets all changes to the "smart" field.
func (m *WebsiteMutation) ResetSmart() {
	m.smart = nil
	m.addsmart = nil
}

// SetBeautiful sets the "beautiful" field.
func (m *WebsiteMutation) SetBeautiful(f float64) {
	m.beautiful = &f
	m.addbeautiful = nil
}

// Beautiful returns the value of the "beautiful" field in the mutation.
func (m *WebsiteMutation) Beautiful() (r float64, exists bool) {
	v := m.beautiful
	if v == nil {
		return
	}
	return *v, true
}

// OldBeautiful returns the old "beautiful" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldBeautiful(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBeautiful is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBeautiful requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBeautiful: %w", err)
	}
	return oldValue.Beautiful, nil
}

// AddBeautiful adds f to the "beautiful" field.
func (m *WebsiteMutation) AddBeautiful(f float64) {
	if m.addbeautiful != nil {
		*m.addbeautiful += f
	} else {
		m.addbeautiful = &f
	}
}

// AddedBeautiful returns the value that was added to the "beautiful" field in this mutation.
func (m *WebsiteMutation) AddedBeautiful() (r float64, exists bool) {
	v := m.addbeautiful
	if v == nil {
		return
	}
	return *v, true
}

// ResetBeautiful resets all changes to the "beautiful" field.
func (m *WebsiteMutation) ResetBeautiful() {
	m.beautiful = nil
	m.addbeautiful = nil
}

// SetLike sets the "like" field.
func (m *WebsiteMutation) SetLike(f float64) {
	m.like = &f
	m.addlike = nil
}

// Like returns the value of the "like" field in the mutation.
func (m *WebsiteMutation) Like() (r float64, exists bool) {
	v := m.like
	if v == nil {
		return
	}
	return *v, true
}

// OldLike returns the old "like" field's value of the Website entity.
// If the Website object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WebsiteMutation) OldLike(ctx context.Context) (v *float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLike is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLike requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLike: %w", err)
	}
	return oldValue.Like, nil
}

// AddLike adds f to the "like" field.
func (m *WebsiteMutation) AddLike(f float64) {
	if m.addlike != nil {
		*m.addlike += f
	} else {
		m.addlike = &f
	}
}

// AddedLike returns the value that was added to the "like" field in this mutation.
func (m *WebsiteMutation) AddedLike() (r float64, exists bool) {
	v := m.addlike
	if v == nil {
		return
	}
	return *v, true
}

// ResetLike resets all changes to the "like" field.
func (m *WebsiteMutation) ResetLike() {
	m.like = nil
	m.addlike = nil
}

// Where appends a list predicates to the WebsiteMutation builder.
func (m *WebsiteMutation) Where(ps ...predicate.Website) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the WebsiteMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *WebsiteMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Website, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *WebsiteMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *WebsiteMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Website).
func (m *WebsiteMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WebsiteMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.name != nil {
		fields = append(fields, website.FieldName)
	}
	if m.filePath != nil {
		fields = append(fields, website.FieldFilePath)
	}
	if m.url != nil {
		fields = append(fields, website.FieldURL)
	}
	if m.bright != nil {
		fields = append(fields, website.FieldBright)
	}
	if m.flashy != nil {
		fields = append(fields, website.FieldFlashy)
	}
	if m.adult != nil {
		fields = append(fields, website.FieldAdult)
	}
	if m.smart != nil {
		fields = append(fields, website.FieldSmart)
	}
	if m.beautiful != nil {
		fields = append(fields, website.FieldBeautiful)
	}
	if m.like != nil {
		fields = append(fields, website.FieldLike)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WebsiteMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case website.FieldName:
		return m.Name()
	case website.FieldFilePath:
		return m.FilePath()
	case website.FieldURL:
		return m.URL()
	case website.FieldBright:
		return m.Bright()
	case website.FieldFlashy:
		return m.Flashy()
	case website.FieldAdult:
		return m.Adult()
	case website.FieldSmart:
		return m.Smart()
	case website.FieldBeautiful:
		return m.Beautiful()
	case website.FieldLike:
		return m.Like()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WebsiteMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case website.FieldName:
		return m.OldName(ctx)
	case website.FieldFilePath:
		return m.OldFilePath(ctx)
	case website.FieldURL:
		return m.OldURL(ctx)
	case website.FieldBright:
		return m.OldBright(ctx)
	case website.FieldFlashy:
		return m.OldFlashy(ctx)
	case website.FieldAdult:
		return m.OldAdult(ctx)
	case website.FieldSmart:
		return m.OldSmart(ctx)
	case website.FieldBeautiful:
		return m.OldBeautiful(ctx)
	case website.FieldLike:
		return m.OldLike(ctx)
	}
	return nil, fmt.Errorf("unknown Website field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WebsiteMutation) SetField(name string, value ent.Value) error {
	switch name {
	case website.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case website.FieldFilePath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFilePath(v)
		return nil
	case website.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case website.FieldBright:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBright(v)
		return nil
	case website.FieldFlashy:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFlashy(v)
		return nil
	case website.FieldAdult:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAdult(v)
		return nil
	case website.FieldSmart:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSmart(v)
		return nil
	case website.FieldBeautiful:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBeautiful(v)
		return nil
	case website.FieldLike:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLike(v)
		return nil
	}
	return fmt.Errorf("unknown Website field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WebsiteMutation) AddedFields() []string {
	var fields []string
	if m.addbright != nil {
		fields = append(fields, website.FieldBright)
	}
	if m.addflashy != nil {
		fields = append(fields, website.FieldFlashy)
	}
	if m.addadult != nil {
		fields = append(fields, website.FieldAdult)
	}
	if m.addsmart != nil {
		fields = append(fields, website.FieldSmart)
	}
	if m.addbeautiful != nil {
		fields = append(fields, website.FieldBeautiful)
	}
	if m.addlike != nil {
		fields = append(fields, website.FieldLike)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WebsiteMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case website.FieldBright:
		return m.AddedBright()
	case website.FieldFlashy:
		return m.AddedFlashy()
	case website.FieldAdult:
		return m.AddedAdult()
	case website.FieldSmart:
		return m.AddedSmart()
	case website.FieldBeautiful:
		return m.AddedBeautiful()
	case website.FieldLike:
		return m.AddedLike()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WebsiteMutation) AddField(name string, value ent.Value) error {
	switch name {
	case website.FieldBright:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBright(v)
		return nil
	case website.FieldFlashy:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFlashy(v)
		return nil
	case website.FieldAdult:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAdult(v)
		return nil
	case website.FieldSmart:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSmart(v)
		return nil
	case website.FieldBeautiful:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBeautiful(v)
		return nil
	case website.FieldLike:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLike(v)
		return nil
	}
	return fmt.Errorf("unknown Website numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WebsiteMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WebsiteMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WebsiteMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Website nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WebsiteMutation) ResetField(name string) error {
	switch name {
	case website.FieldName:
		m.ResetName()
		return nil
	case website.FieldFilePath:
		m.ResetFilePath()
		return nil
	case website.FieldURL:
		m.ResetURL()
		return nil
	case website.FieldBright:
		m.ResetBright()
		return nil
	case website.FieldFlashy:
		m.ResetFlashy()
		return nil
	case website.FieldAdult:
		m.ResetAdult()
		return nil
	case website.FieldSmart:
		m.ResetSmart()
		return nil
	case website.FieldBeautiful:
		m.ResetBeautiful()
		return nil
	case website.FieldLike:
		m.ResetLike()
		return nil
	}
	return fmt.Errorf("unknown Website field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WebsiteMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WebsiteMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WebsiteMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WebsiteMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WebsiteMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WebsiteMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WebsiteMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Website unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WebsiteMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Website edge %s", name)
}
