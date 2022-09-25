package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/app/content/service/internal/data/ent"
	"kratos-blog/app/content/service/internal/data/ent/post"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.PostRepo = (*PostRepo)(nil)

type PostRepo struct {
	data *Data
	log  *log.Helper
}

func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	l := log.NewHelper(log.With(logger, "module", "post/repo"))
	return &PostRepo{
		data: data,
		log:  l,
	}
}

func (r *PostRepo) convertEntToProto(in *ent.Post) *v1.Post {
	if in == nil {
		return nil
	}
	return &v1.Post{
		Id:              in.ID,
		Title:           in.Title,
		Status:          in.Status,
		Slug:            in.Slug,
		EditorType:      in.EditorType,
		MetaKeywords:    in.MetaKeywords,
		MetaDescription: in.MetaDescription,
		FullPath:        in.FullPath,
		Summary:         in.Summary,
		Thumbnail:       in.Thumbnail,
		Password:        in.Password,
		Template:        in.Template,
		Content:         in.Content,
		OriginalContent: in.OriginalContent,
		Visits:          in.Visits,
		TopPriority:     in.TopPriority,
		Likes:           in.Likes,
		WordCount:       in.WordCount,
		CommentCount:    in.CommentCount,
		DisallowComment: in.DisallowComment,
		InProgress:      in.InProgress,
		CreateTime:      entgo.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:      entgo.UnixMilliToStringPtr(in.UpdateTime),
		EditTime:        entgo.UnixMilliToStringPtr(in.EditTime),
	}
}

func (r *PostRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Post.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder1.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder1.Order(v)
		}
	} else {
		builder1.Order(ent.Desc(post.FieldCreateTime))
	}
	if req.Page != nil && req.PageSize != nil {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	posts, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Post.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder2.Where(v)
		}
	}
	count, err := builder2.
		Select(post.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Post, 0, len(posts))
	for _, po := range posts {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListPostResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *PostRepo) Get(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	po, err := r.data.db.Post.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Create(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error) {
	po, err := r.data.db.Post.Create().
		SetNillableTitle(req.Post.Title).
		SetNillableStatus(req.Post.Status).
		SetNillableSlug(req.Post.Slug).
		SetNillableEditorType(req.Post.EditorType).
		SetNillableMetaKeywords(req.Post.MetaKeywords).
		SetNillableMetaDescription(req.Post.MetaDescription).
		SetNillableFullPath(req.Post.FullPath).
		SetNillableSummary(req.Post.Summary).
		SetNillableThumbnail(req.Post.Thumbnail).
		SetNillablePassword(req.Post.Password).
		SetNillableTemplate(req.Post.Template).
		SetNillableContent(req.Post.Content).
		SetNillableOriginalContent(req.Post.OriginalContent).
		SetNillableVisits(req.Post.Visits).
		SetNillableTopPriority(req.Post.TopPriority).
		SetNillableLikes(req.Post.Likes).
		SetNillableWordCount(req.Post.WordCount).
		SetNillableCommentCount(req.Post.CommentCount).
		SetNillableDisallowComment(req.Post.DisallowComment).
		SetNillableInProgress(req.Post.InProgress).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Update(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error) {
	builder := r.data.db.Post.UpdateOneID(req.Id).
		SetNillableTitle(req.Post.Title).
		SetNillableStatus(req.Post.Status).
		SetNillableSlug(req.Post.Slug).
		SetNillableEditorType(req.Post.EditorType).
		SetNillableMetaKeywords(req.Post.MetaKeywords).
		SetNillableMetaDescription(req.Post.MetaDescription).
		SetNillableFullPath(req.Post.FullPath).
		SetNillableSummary(req.Post.Summary).
		SetNillableThumbnail(req.Post.Thumbnail).
		SetNillablePassword(req.Post.Password).
		SetNillableTemplate(req.Post.Template).
		SetNillableContent(req.Post.Content).
		SetNillableOriginalContent(req.Post.OriginalContent).
		SetNillableVisits(req.Post.Visits).
		SetNillableTopPriority(req.Post.TopPriority).
		SetNillableLikes(req.Post.Likes).
		SetNillableWordCount(req.Post.WordCount).
		SetNillableCommentCount(req.Post.CommentCount).
		SetNillableDisallowComment(req.Post.DisallowComment).
		SetNillableInProgress(req.Post.InProgress).
		SetEditTime(time.Now().UnixMilli())

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Delete(ctx context.Context, req *v1.DeletePostRequest) (bool, error) {
	err := r.data.db.Post.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
