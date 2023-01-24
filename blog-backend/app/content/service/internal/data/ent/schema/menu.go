package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/kratos-blog/blog-backend/pkg/util/entgo/mixin"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "menu",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("url").
			Optional().
			Nillable(),

		field.Int32("priority").
			Optional().
			Nillable(),

		field.String("target").
			Optional().
			Nillable(),

		field.String("icon").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Optional().
			Nillable(),

		field.String("team").
			Optional().
			Nillable(),
	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
