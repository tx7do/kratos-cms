package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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

		field.String("slug").
			Optional().
			Nillable(),

		field.String("meta_keywords").
			Optional().
			Nillable(),

		field.String("meta_description").
			Optional().
			Nillable(),

		field.String("full_path").
			Optional().
			Nillable(),

		field.String("original_content").
			Optional().
			Nillable(),

		field.String("content").
			Optional().
			Nillable(),

		field.String("summary").
			Optional().
			Nillable(),

		field.String("thumbnail").
			Optional().
			Nillable(),

		field.String("password").
			Optional().
			Nillable(),

		field.String("template").
			Optional().
			Nillable(),

		field.Int32("comment_count").
			Optional().
			Nillable(),

		field.Int32("visits").
			Optional().
			Nillable(),

		field.Int32("likes").
			Optional().
			Nillable(),

		field.Int32("word_count").
			Optional().
			Nillable(),

		field.Int32("top_priority").
			Optional().
			Nillable(),

		field.Int32("status").
			Optional().
			Nillable(),

		field.Int32("editor_type").
			Optional().
			Nillable(),

		field.Int64("edit_time").
			Comment("编辑时间").
			Optional().
			Nillable(),

		field.Bool("disallow_comment").
			Optional().
			Nillable(),

		field.Bool("in_progress").
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
