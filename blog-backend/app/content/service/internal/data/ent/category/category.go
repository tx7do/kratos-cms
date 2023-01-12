// Code generated by ent, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFullPath holds the string denoting the full_path field in the database.
	FieldFullPath = "full_path"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldPostCount holds the string denoting the post_count field in the database.
	FieldPostCount = "post_count"
	// Table holds the table name of the category in the database.
	Table = "category"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldName,
	FieldSlug,
	FieldDescription,
	FieldThumbnail,
	FieldPassword,
	FieldFullPath,
	FieldParentID,
	FieldPriority,
	FieldPostCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() int64
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() int64
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	ThumbnailValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)