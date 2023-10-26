package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("目录名").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("url").
			Comment("链接").
			Optional().
			Nillable(),

		field.Int32("priority").
			Comment("优先级").
			Optional().
			Nillable(),

		field.String("target").
			Comment("目标").
			Optional().
			Nillable(),

		field.String("icon").
			Comment("图标").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("父目录ID").
			Optional().
			Nillable(),

		field.String("team").
			Comment("分组").
			Optional().
			Nillable(),
	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
