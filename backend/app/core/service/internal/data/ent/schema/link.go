package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("链接名").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("url").
			Comment("链接").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("logo").
			Comment("图标").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("description").
			Comment("说明").
			Optional().
			Nillable(),

		field.String("team").
			Comment("分组").
			Optional().
			Nillable(),

		field.Int32("priority").
			Comment("优先级").
			Optional().
			Nillable(),
	}
}

// Mixin of the Link.
func (Link) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}
