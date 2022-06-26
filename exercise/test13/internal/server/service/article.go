package service

import "context"

type article struct {
repo domain.IArticleRepo
}

// NewArticleUsecase init
func NewArticleUsecase(repo domain.IArticleRepo) domain.IArticleUsecase {
return &article{repo: repo}
}

func (u *article) GetArticle(ctx context.Context, id int) (*domain.Article, error) {
// 这里可能有其他业务逻辑...
return u.repo.GetArticle(ctx, id)
}

func (u *article) CreateArticle(ctx context.Context, article *domain.Article) error {
return u.repo.CreateArticle(ctx, article)
}

// 确保实现了对应的接口
var _ v1.BlogServiceHTTPServer = &Artcile{}

// Artcile Artcile
type Artcile struct {
	usecase domain.IArticleUsecase
}

// NewArticleService 初始化方法
func NewArticleService(usecase domain.IArticleUsecase) *Artcile {
	return &Artcile{usecase: usecase}
}

// CreateArticle 创建一篇文章
func (a *Artcile) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleResp, error) {
	article := &domain.Article{
		Title:   req.Title,
		Content: req.Content,
	}
	err := a.usecase.CreateArticle(ctx, article)
	return &v1.CreateArticleResp{}, err
}