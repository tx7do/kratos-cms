package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/kratos-blog/blog-backend/pkg/util/entgo/mixin"
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

		field.String("url").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("logo").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("description").
			Comment("说明").
			Optional().
			Nillable(),

		field.String("team").
			Optional().
			Nillable(),

		field.Int32("priority").
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
