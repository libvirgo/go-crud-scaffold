package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// UserActivity holds the schema definition for the UserActivity entity.
type UserActivity struct {
	ent.Schema
}

// Fields of the UserActivity.
func (UserActivity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type"), // 1 register 2 login
	}
}

// Edges of the UserActivity.
func (UserActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_activity").Unique(),
	}
}

// Indexes of the UserActivity.
func (UserActivity) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of the UserActivity.
func (UserActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
