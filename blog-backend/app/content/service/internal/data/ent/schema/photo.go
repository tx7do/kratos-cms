package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/kratos-blog/blog-backend/pkg/util/entgo/mixin"
)

// Photo holds the schema definition for the Photo entity.
type Photo struct {
	ent.Schema
}

func (Photo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "photo",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Nillable(),

		field.String("thumbnail").
			Optional().
			Nillable(),

		field.Int64("take_time").
			Comment("更新时间").
			Optional().
			Nillable(),

		field.String("url").
			Optional().
			Nillable(),

		field.String("team").
			Optional().
			Nillable(),

		field.String("location").
			Optional().
			Nillable(),

		field.String("description").
			Optional().
			Nillable(),

		field.Int32("likes").
			Optional().
			Nillable(),
	}
}

// Mixin of the Photo.
func (Photo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Photo.
func (Photo) Edges() []ent.Edge {
	return nil
}
