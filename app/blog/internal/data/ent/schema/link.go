package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kratos-blog/pkg/util/entgo/mixin"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

func (Link) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "link",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("link").
			Optional().
			Nillable(),

		field.Int32("order_id").
			Optional().
			Nillable(),
	}
}

// Mixin of the Link.
func (Link) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}

// Indexes of the Link.
func (Link) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").StorageKey("link_pkey"),
	}
}
