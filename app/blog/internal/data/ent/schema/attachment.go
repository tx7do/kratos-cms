package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kratos-blog/pkg/util/entgo/mixin"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

func (Attachment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "attachment",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Nillable(),

		field.String("path").
			Optional().
			Nillable(),

		field.String("file_key").
			Optional().
			Nillable(),

		field.String("thumb_path").
			Optional().
			Nillable(),

		field.String("media_type").
			Optional().
			Nillable(),

		field.String("suffix").
			Optional().
			Nillable(),

		field.Int32("width").
			Optional().
			Nillable(),

		field.Int32("height").
			Optional().
			Nillable(),

		field.Int64("size").
			Optional().
			Nillable(),

		field.Int32("type").
			Optional().
			Nillable(),
	}
}

// Mixin of the Attachment.
func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return nil
}
