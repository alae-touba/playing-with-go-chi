package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Unique().
			Immutable(),
		field.String("username").
			Unique().
			NotEmpty().
			MinLen(3),
		field.String("password").
			Sensitive().
			NotEmpty().
			MinLen(6),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
