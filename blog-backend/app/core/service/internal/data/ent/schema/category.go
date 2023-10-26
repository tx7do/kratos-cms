package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "category",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("分类名").
			Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("slug").
			Comment("链接别名").
			Optional().
			Nillable(),

		field.String("description").
			Comment("描述").
			MaxLen(100).
			Optional().
			Nillable(),

		field.String("thumbnail").
			Comment("缩略图").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("password").
			Comment("密码").
			Optional().
			Nillable(),

		field.String("full_path").
			Comment("完整路径").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("父分类ID").
			Optional().
			Nillable(),

		field.Int32("priority").
			Comment("优先级").
			Optional().
			Nillable(),

		field.Uint32("post_count").
			Comment("博文计数").
			Optional().
			Nillable(),
	}
}

// Mixin of the Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
