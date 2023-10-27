package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("wallet_address").NotEmpty(),
	}
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the AdminUser.
func (AdminUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wallet_address"),
	}
}

// Mixin of the AdminUser.
func (AdminUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
