package article

import (
	"awesomeProject/goweb/Handler"
	"awesomeProject/goweb/models"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return Handler.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
