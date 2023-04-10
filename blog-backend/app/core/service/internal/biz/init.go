package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAttachmentUseCase,
	NewCategoryUseCase,
	NewCommentUseCase,
	NewLinkUseCase,
	NewPostUseCase,
	NewTagUseCase,
	NewUserUseCase,
	NewUserAuthUseCase,
)
