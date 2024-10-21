//go:build wireinject
// +build wireinject

package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewEntClient,
	NewRedisClient,
	NewMinIoClient,

	NewAttachmentRepo,
	NewCategoryRepo,
	NewCommentRepo,
	NewLinkRepo,
	NewPostRepo,
	NewTagRepo,
	NewUserRepo,
)
