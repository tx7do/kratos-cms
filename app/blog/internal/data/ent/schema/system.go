package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kratos-blog/pkg/util/entgo/mixin"
)

// System holds the schema definition for the System entity.
type System struct {
	ent.Schema
}

func (System) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "system",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the System.
func (System) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Optional().
			Nillable(),

		field.String("keywords").
			Optional().
			Nillable(),

		field.String("description").
			Optional().
			Nillable(),

		field.String("record_number").
			Optional().
			Nillable(),

		field.Int32("Theme").
			Optional().
			Nillable(),
	}
}

// Mixin of the System.
func (System) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the System.
func (System) Edges() []ent.Edge {
	return nil
}

// Indexes of the System.
func (System) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").StorageKey("system_pkey"),
	}
}
