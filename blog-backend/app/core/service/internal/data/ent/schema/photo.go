package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("图片名").
			Optional().
			Nillable(),

		field.String("thumbnail").
			Comment("缩略图").
			Optional().
			Nillable(),

		field.Int64("take_time").
			Comment("更新时间").
			Optional().
			Nillable(),

		field.String("url").
			Comment("链接").
			Optional().
			Nillable(),

		field.String("team").
			Comment("分组").
			Optional().
			Nillable(),

		field.String("location").
			Comment("地理位置").
			Optional().
			Nillable(),

		field.String("description").
			Comment("描述").
			Optional().
			Nillable(),

		field.Int32("likes").
			Comment("点赞数").
			Optional().
			Nillable(),
	}
}

// Mixin of the Photo.
func (Photo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Photo.
func (Photo) Edges() []ent.Edge {
	return nil
}
