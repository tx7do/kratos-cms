package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("博文标题").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("slug").
			Comment("链接别名").
			Optional().
			Nillable(),

		field.String("meta_keywords").
			Comment("关键词列表").
			Optional().
			Nillable(),

		field.String("meta_description").
			Comment("描述信息").
			Optional().
			Nillable(),

		field.String("full_path").
			Comment("全路径").
			Optional().
			Nillable(),

		field.String("original_content").
			Comment("原始内容").
			Optional().
			Nillable(),

		field.String("content").
			Comment("内容").
			Optional().
			Nillable(),

		field.String("summary").
			Comment("概要信息").
			Optional().
			Nillable(),

		field.String("thumbnail").
			Comment("缩略图").
			Optional().
			Nillable(),

		field.String("password").
			Comment("密码").
			Optional().
			Nillable(),

		field.String("template").
			Comment("模板").
			Optional().
			Nillable(),

		field.Int32("comment_count").
			Comment("评论数").
			Optional().
			Nillable(),

		field.Int32("visits").
			Comment("访问数").
			Optional().
			Nillable(),

		field.Int32("likes").
			Comment("点赞数").
			Optional().
			Nillable(),

		field.Int32("word_count").
			Comment("文章字数").
			Optional().
			Nillable(),

		field.Int32("top_priority").
			Comment("优先级").
			Optional().
			Nillable(),

		field.Int32("status").
			Comment("状态").
			Optional().
			Nillable(),

		field.Int32("editor_type").
			Comment("编辑器类型").
			Optional().
			Nillable(),

		field.Int64("edit_time").
			Comment("编辑时间").
			Optional().
			Nillable(),

		field.Bool("disallow_comment").
			Comment("不允许评论").
			Optional().
			Nillable(),

		field.Bool("in_progress").
			Comment("审核中").
			Optional().
			Nillable(),
	}
}

// Mixin of the Post.
func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
