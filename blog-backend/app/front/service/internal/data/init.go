package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewDiscovery,

	NewAuthenticator,
	NewAuthorizer,

	NewUserServiceClient,
	NewAttachmentServiceClient,
	NewCommentServiceClient,
	NewCategoryServiceClient,
	NewLinkServiceClient,
	NewPostServiceClient,
	NewTagServiceClient,
)
