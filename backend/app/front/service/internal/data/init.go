//go:build wireinject
// +build wireinject

package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewDiscovery,

	NewAuthenticator,
	NewAuthorizer,

	NewUserTokenRepo,

	NewUserServiceClient,
	NewAttachmentServiceClient,
	NewCommentServiceClient,
	NewCategoryServiceClient,
	NewLinkServiceClient,
	NewPostServiceClient,
	NewTagServiceClient,
	NewFileServiceClient,
)
