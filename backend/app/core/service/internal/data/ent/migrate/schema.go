// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentColumns holds the columns for the "attachment" table.
	AttachmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Comment: "文件名"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "文件路径"},
		{Name: "file_key", Type: field.TypeString, Nullable: true, Size: 100, Comment: "文件键"},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "缩略图"},
		{Name: "media_type", Type: field.TypeString, Nullable: true, Comment: "媒体类型"},
		{Name: "suffix", Type: field.TypeString, Nullable: true, Comment: "后缀"},
		{Name: "width", Type: field.TypeInt32, Nullable: true, Comment: "图片宽"},
		{Name: "height", Type: field.TypeInt32, Nullable: true, Comment: "图片高"},
		{Name: "size", Type: field.TypeUint64, Nullable: true, Comment: "文件大小"},
		{Name: "type", Type: field.TypeInt32, Nullable: true, Comment: "类型"},
	}
	// AttachmentTable holds the schema information for the "attachment" table.
	AttachmentTable = &schema.Table{
		Name:       "attachment",
		Columns:    AttachmentColumns,
		PrimaryKey: []*schema.Column{AttachmentColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "attachment_id",
				Unique:  false,
				Columns: []*schema.Column{AttachmentColumns[0]},
			},
		},
	}
	// CategoryColumns holds the columns for the "category" table.
	CategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Comment: "分类名"},
		{Name: "slug", Type: field.TypeString, Nullable: true, Comment: "链接别名"},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 100, Comment: "描述"},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "缩略图"},
		{Name: "password", Type: field.TypeString, Nullable: true, Comment: "密码"},
		{Name: "full_path", Type: field.TypeString, Nullable: true, Comment: "完整路径"},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "父分类ID"},
		{Name: "priority", Type: field.TypeInt32, Nullable: true, Comment: "优先级"},
		{Name: "post_count", Type: field.TypeUint32, Nullable: true, Comment: "博文计数"},
	}
	// CategoryTable holds the schema information for the "category" table.
	CategoryTable = &schema.Table{
		Name:       "category",
		Columns:    CategoryColumns,
		PrimaryKey: []*schema.Column{CategoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "category_id",
				Unique:  false,
				Columns: []*schema.Column{CategoryColumns[0]},
			},
		},
	}
	// CommentColumns holds the columns for the "comment" table.
	CommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "author", Type: field.TypeString, Nullable: true, Comment: "作者"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "邮箱地址"},
		{Name: "ip_address", Type: field.TypeString, Nullable: true, Comment: "IP地址"},
		{Name: "author_url", Type: field.TypeString, Nullable: true, Comment: "作者链接"},
		{Name: "gravatar_md5", Type: field.TypeString, Nullable: true, Comment: "MD5"},
		{Name: "content", Type: field.TypeString, Nullable: true, Comment: "内容"},
		{Name: "user_agent", Type: field.TypeString, Nullable: true, Comment: "用户浏览器信息"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "头像"},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "父评论ID"},
		{Name: "status", Type: field.TypeUint32, Nullable: true, Comment: "状态"},
		{Name: "is_admin", Type: field.TypeBool, Nullable: true, Comment: "是否管理员"},
		{Name: "allow_notification", Type: field.TypeBool, Nullable: true, Comment: "允许通知"},
	}
	// CommentTable holds the schema information for the "comment" table.
	CommentTable = &schema.Table{
		Name:       "comment",
		Columns:    CommentColumns,
		PrimaryKey: []*schema.Column{CommentColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "comment_id",
				Unique:  false,
				Columns: []*schema.Column{CommentColumns[0]},
			},
		},
	}
	// LinkColumns holds the columns for the "link" table.
	LinkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Comment: "链接名"},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "链接"},
		{Name: "logo", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "图标"},
		{Name: "description", Type: field.TypeString, Nullable: true, Comment: "说明"},
		{Name: "team", Type: field.TypeString, Nullable: true, Comment: "分组"},
		{Name: "priority", Type: field.TypeInt32, Nullable: true, Comment: "优先级"},
	}
	// LinkTable holds the schema information for the "link" table.
	LinkTable = &schema.Table{
		Name:       "link",
		Columns:    LinkColumns,
		PrimaryKey: []*schema.Column{LinkColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "link_id",
				Unique:  false,
				Columns: []*schema.Column{LinkColumns[0]},
			},
		},
	}
	// MenuColumns holds the columns for the "menu" table.
	MenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Comment: "目录名"},
		{Name: "url", Type: field.TypeString, Nullable: true, Comment: "链接"},
		{Name: "priority", Type: field.TypeInt32, Nullable: true, Comment: "优先级"},
		{Name: "target", Type: field.TypeString, Nullable: true, Comment: "目标"},
		{Name: "icon", Type: field.TypeString, Nullable: true, Comment: "图标"},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "父目录ID"},
		{Name: "team", Type: field.TypeString, Nullable: true, Comment: "分组"},
	}
	// MenuTable holds the schema information for the "menu" table.
	MenuTable = &schema.Table{
		Name:       "menu",
		Columns:    MenuColumns,
		PrimaryKey: []*schema.Column{MenuColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "menu_id",
				Unique:  false,
				Columns: []*schema.Column{MenuColumns[0]},
			},
		},
	}
	// PhotoColumns holds the columns for the "photo" table.
	PhotoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Nullable: true, Comment: "图片名"},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Comment: "缩略图"},
		{Name: "take_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "url", Type: field.TypeString, Nullable: true, Comment: "链接"},
		{Name: "team", Type: field.TypeString, Nullable: true, Comment: "分组"},
		{Name: "location", Type: field.TypeString, Nullable: true, Comment: "地理位置"},
		{Name: "description", Type: field.TypeString, Nullable: true, Comment: "描述"},
		{Name: "likes", Type: field.TypeInt32, Nullable: true, Comment: "点赞数"},
	}
	// PhotoTable holds the schema information for the "photo" table.
	PhotoTable = &schema.Table{
		Name:       "photo",
		Columns:    PhotoColumns,
		PrimaryKey: []*schema.Column{PhotoColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "photo_id",
				Unique:  false,
				Columns: []*schema.Column{PhotoColumns[0]},
			},
		},
	}
	// PostColumns holds the columns for the "post" table.
	PostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "title", Type: field.TypeString, Nullable: true, Comment: "博文标题"},
		{Name: "slug", Type: field.TypeString, Nullable: true, Comment: "链接别名"},
		{Name: "meta_keywords", Type: field.TypeString, Nullable: true, Comment: "关键词列表"},
		{Name: "meta_description", Type: field.TypeString, Nullable: true, Comment: "描述信息"},
		{Name: "full_path", Type: field.TypeString, Nullable: true, Comment: "全路径"},
		{Name: "original_content", Type: field.TypeString, Nullable: true, Comment: "原始内容"},
		{Name: "content", Type: field.TypeString, Nullable: true, Comment: "内容"},
		{Name: "summary", Type: field.TypeString, Nullable: true, Comment: "概要信息"},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Comment: "缩略图"},
		{Name: "password", Type: field.TypeString, Nullable: true, Comment: "密码"},
		{Name: "template", Type: field.TypeString, Nullable: true, Comment: "模板"},
		{Name: "comment_count", Type: field.TypeInt32, Nullable: true, Comment: "评论数"},
		{Name: "visits", Type: field.TypeInt32, Nullable: true, Comment: "访问数"},
		{Name: "likes", Type: field.TypeInt32, Nullable: true, Comment: "点赞数"},
		{Name: "word_count", Type: field.TypeInt32, Nullable: true, Comment: "文章字数"},
		{Name: "top_priority", Type: field.TypeInt32, Nullable: true, Comment: "优先级"},
		{Name: "status", Type: field.TypeInt32, Nullable: true, Comment: "状态"},
		{Name: "editor_type", Type: field.TypeInt32, Nullable: true, Comment: "编辑器类型"},
		{Name: "edit_time", Type: field.TypeInt64, Nullable: true, Comment: "编辑时间"},
		{Name: "disallow_comment", Type: field.TypeBool, Nullable: true, Comment: "不允许评论"},
		{Name: "in_progress", Type: field.TypeBool, Nullable: true, Comment: "审核中"},
	}
	// PostTable holds the schema information for the "post" table.
	PostTable = &schema.Table{
		Name:       "post",
		Columns:    PostColumns,
		PrimaryKey: []*schema.Column{PostColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "post_id",
				Unique:  false,
				Columns: []*schema.Column{PostColumns[0]},
			},
		},
	}
	// TagColumns holds the columns for the "tag" table.
	TagColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Comment: "标签名"},
		{Name: "color", Type: field.TypeString, Nullable: true, Comment: "颜色"},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Comment: "缩略图"},
		{Name: "slug", Type: field.TypeString, Nullable: true, Comment: "链接别名"},
		{Name: "slug_name", Type: field.TypeString, Nullable: true, Comment: "链接别名"},
		{Name: "post_count", Type: field.TypeUint32, Nullable: true, Comment: "博文计数"},
	}
	// TagTable holds the schema information for the "tag" table.
	TagTable = &schema.Table{
		Name:       "tag",
		Comment:    "标签",
		Columns:    TagColumns,
		PrimaryKey: []*schema.Column{TagColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_id",
				Unique:  false,
				Columns: []*schema.Column{TagColumns[0]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true, Size: 50, Comment: "用户名"},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, Comment: "登陆密码"},
		{Name: "nickname", Type: field.TypeString, Nullable: true, Size: 128, Comment: "昵称"},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 127, Comment: "电子邮箱"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "头像"},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "个人说明"},
		{Name: "authority", Type: field.TypeEnum, Nullable: true, Comment: "授权", Enums: []string{"SYS_ADMIN", "CUSTOMER_USER", "GUEST_USER", "REFRESH_TOKEN"}, Default: "CUSTOMER_USER"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[0]},
			},
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentTable,
		CategoryTable,
		CommentTable,
		LinkTable,
		MenuTable,
		PhotoTable,
		PostTable,
		TagTable,
		UserTable,
	}
)

func init() {
	AttachmentTable.Annotation = &entsql.Annotation{
		Table:     "attachment",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	CategoryTable.Annotation = &entsql.Annotation{
		Table:     "category",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	CommentTable.Annotation = &entsql.Annotation{
		Table:     "comment",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	LinkTable.Annotation = &entsql.Annotation{
		Table:     "link",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	MenuTable.Annotation = &entsql.Annotation{
		Table:     "menu",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PhotoTable.Annotation = &entsql.Annotation{
		Table:     "photo",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PostTable.Annotation = &entsql.Annotation{
		Table:     "post",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	TagTable.Annotation = &entsql.Annotation{
		Table:     "tag",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table:     "user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}