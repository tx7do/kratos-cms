package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserUseCase,
	NewPostUseCase,
	NewCategoryUseCase,
	NewLinkUseCase,
	NewTagUseCase,
	NewUserAuthUseCase,
	NewCommentUseCase,
	NewAttachmentUseCase,
)
