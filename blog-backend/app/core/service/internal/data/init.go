package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewEntClient,
	NewRedisClient,

	NewAttachmentRepo,
	NewCategoryRepo,
	NewCommentRepo,
	NewLinkRepo,
	NewPostRepo,
	NewTagRepo,
	NewUserRepo,
	NewUserTokenRepo,
)
