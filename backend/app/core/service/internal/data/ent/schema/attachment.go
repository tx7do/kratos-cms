package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("文件名").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("path").
			Comment("文件路径").
			Optional().
			Nillable(),

		field.String("file_key").
			Comment("文件键").
			MaxLen(100).
			Optional().
			Nillable(),

		field.String("thumbnail").
			Comment("缩略图").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("media_type").
			Comment("媒体类型").
			Optional().
			Nillable(),

		field.String("suffix").
			Comment("后缀").
			Optional().
			Nillable(),

		field.Int32("width").
			Comment("图片宽").
			Optional().
			Nillable(),

		field.Int32("height").
			Comment("图片高").
			Optional().
			Nillable(),

		field.Uint64("Size").
			Comment("文件大小").
			Optional().
			Nillable(),

		field.Int32("type").
			Comment("类型").
			Optional().
			Nillable(),
	}
}

// Mixin of the Attachment.
func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return nil
}
