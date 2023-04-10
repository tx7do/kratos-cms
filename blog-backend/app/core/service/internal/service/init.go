package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAttachmentService,
	NewCategoryService,
	NewCommentService,
	NewLinkService,
	NewPostService,
	NewTagService,
	NewUserService,
)
