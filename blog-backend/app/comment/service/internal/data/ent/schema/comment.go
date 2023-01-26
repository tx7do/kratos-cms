package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kratos-blog/pkg/util/entgo/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "comment",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("author").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("email").
			Optional().
			Nillable(),

		field.String("ip_address").
			Optional().
			Nillable(),

		field.String("author_url").
			Optional().
			Nillable(),

		field.String("gravatar_md5").
			Optional().
			Nillable(),

		field.String("content").
			Optional().
			Nillable(),

		field.String("user_agent").
			Optional().
			Nillable(),

		field.String("avatar").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Optional().
			Nillable(),

		field.Uint32("status").
			Optional().
			Nillable(),

		field.Bool("is_admin").
			Optional().
			Nillable(),

		field.Bool("allow_notification").
			Optional().
			Nillable(),
	}
}

// Mixin of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
