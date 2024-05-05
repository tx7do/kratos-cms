package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
		entsql.WithComments(true),
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("author").
			Comment("作者").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("email").
			Comment("邮箱地址").
			Optional().
			Nillable(),

		field.String("ip_address").
			Comment("IP地址").
			Optional().
			Nillable(),

		field.String("author_url").
			Comment("作者链接").
			Optional().
			Nillable(),

		field.String("gravatar_md5").
			Comment("MD5").
			Optional().
			Nillable(),

		field.String("content").
			Comment("内容").
			Optional().
			Nillable(),

		field.String("user_agent").
			Comment("用户浏览器信息").
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("头像").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("父评论ID").
			Optional().
			Nillable(),

		field.Uint32("status").
			Comment("状态").
			Optional().
			Nillable(),

		field.Bool("is_admin").
			Comment("是否管理员").
			Optional().
			Nillable(),

		field.Bool("allow_notification").
			Comment("允许通知").
			Optional().
			Nillable(),
	}
}

// Mixin of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
