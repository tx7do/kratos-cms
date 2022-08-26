package mixin

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type AutoIncrementId struct {
	mixin.Schema
}

func (AutoIncrementId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Comment("id").
			StructTag(`json:"id,omitempty"`).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigserial",
			}).
			Annotations(
				entproto.Field(1),
			).
			Positive().
			Immutable().
			Unique(),
	}
}

// Indexes of the AutoIncrementId.
func (AutoIncrementId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
