package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/kratos-blog/blog-backend/pkg/util/entgo/mixin"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "category",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("slug").
			Optional().
			Nillable(),

		field.String("description").
			MaxLen(100).
			Optional().
			Nillable(),

		field.String("thumbnail").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("password").
			Optional().
			Nillable(),

		field.String("full_path").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Optional().
			Nillable(),

		field.Int32("priority").
			Optional().
			Nillable(),

		field.Uint32("post_count").
			Optional().
			Nillable(),
	}
}

// Mixin of the Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
