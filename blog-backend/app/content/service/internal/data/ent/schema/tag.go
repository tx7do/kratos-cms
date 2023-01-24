package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/kratos-blog/blog-backend/pkg/util/entgo/mixin"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "tag",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("slug").
			Optional().
			Nillable(),

		field.String("color").
			Optional().
			Nillable(),

		field.String("thumbnail").
			Optional().
			Nillable(),

		field.String("slug_name").
			Optional().
			Nillable(),

		field.Uint32("post_count").
			Optional().
			Nillable(),
	}
}

// Mixin of the Tag.
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
