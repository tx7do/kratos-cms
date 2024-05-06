package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/post"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-cms/gen/api/go/content/service/v1"
)

type PostRepo struct {
	data *Data
	log  *log.Helper
}

func NewPostRepo(data *Data, logger log.Logger) *PostRepo {
	l := log.NewHelper(log.With(logger, "module", "post/repo/core-service"))
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
		CreateTime:      util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:      util.UnixMilliToStringPtr(in.UpdateTime),
		EditTime:        util.UnixMilliToStringPtr(in.EditTime),
	}
}

func (r *PostRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Post.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *PostRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	builder := r.data.db.Client().Post.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), post.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.Post, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListPostResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *PostRepo) Get(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	res, err := r.data.db.Client().Post.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *PostRepo) Create(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error) {
	res, err := r.data.db.Client().Post.Create().
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
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *PostRepo) Update(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error) {
	builder := r.data.db.Client().Post.UpdateOneID(req.Id).
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
		SetEditTime(time.Now().UnixMilli()).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *PostRepo) Delete(ctx context.Context, req *v1.DeletePostRequest) (bool, error) {
	err := r.data.db.Client().Post.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
