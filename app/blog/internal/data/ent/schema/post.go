package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kratos-blog/pkg/util/entgo/mixin"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "post",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("summary").
			Optional().
			Nillable(),

		field.String("original").
			Optional().
			Nillable(),

		field.String("content").
			Optional().
			Nillable(),

		field.String("password").
			Optional().
			Nillable(),

		field.Int64("user_id").
			Optional().
			Nillable(),
	}
}

// Mixin of the Post.
func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

// Indexes of the Post.
func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").StorageKey("post_pkey"),
	}
}
