package models

import (
	"fmt"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ArticlePage struct {
	Number   int `json: "number"`
	Articles []ArticleList `json: "articles"`
}


type ArticleList struct {
	ID 				  int    		`json:"id"`
	Author    		  UserList    	`json:"author"`
	AuthorId		  int 			`json:"author_id"`
	Content   		  string    	`json:"content"`
	Comments		  []CommentList	`json:"comments"`
	Title     		  string 		`json:"title"`
	Tags	  		  []Tag  		`json:"tags"`
	UpdatedAt 		  string 		`json:"updated_at"`
}

type Article struct {
	ID 		  int       `json:"id, omitempty"`
	Author    int       `json:"author_id"`
	Title     string    `json:"title"`
	Tags	  []Tag		`json:"tags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllArticles() []ArticleList {
	articles := []ArticleList{}

	// 选择所有的文章id
	row, err := utils.GetConn().Query("SELECT a.id FROM Articles a")
	if err != nil {
		fmt.Println("error:", err)
	}
	for row.Next() {
		articleId := -1
		err = row.Scan(&articleId)
		articles = append(articles, *GetArticleByID(articleId))
	}
	return articles
}

func CreateArticle(article Article) bool {
	// 向Articles表添加条目
	stmt, err := utils.GetConn().Prepare("insert into Articles values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("db insert prepare error")
	}
	_, err = stmt.Exec(nil, article.Title, article.Content, time.Now(), article.Author)
	if err != nil {
		panic("db insert error")
	}
	for _, v := range article.Tags {
		CreateTag(v.Content)

		// 向postTags表添加条目
		stmt, err = utils.GetConn().Prepare("insert into postTags values (?, ?, ?, ?)")
		if err != nil {
			panic("db insert prepare error")
		}
		tagid := getTagId(v.Content)
		_, err = stmt.Exec(time.Now(), time.Now(), article.ID, tagid)
		if err != nil {
			panic("db insert error")
		}
	}
	return true
}

func GetArticleByID(id int) *ArticleList {
	article := ArticleList{}
	row, err := utils.GetConn().Query("SELECT * FROM Articles a WHERE a.id=?", id)
	if err != nil {
		fmt.Println("error:", err)
	}
	for row.Next() {
		err = row.Scan(&article.ID, &article.Title, &article.Content, nil, article.UpdatedAt, article.AuthorId)
		article.Author = *GetUserListByID(article.AuthorId)

		// 寻找文章拥有的tag
		row2, err := utils.GetConn().Query("SELECT t.content FROM Tags t, postTags p WHERE t.id=p.TagId and p.ArticleId=?", article.ID)
		if err != nil {
			fmt.Println("error:", err)
		}
		tags := []Tag{}
		for row2.Next() {
			tag := Tag{}
			err = row2.Scan(&tag.Content)
			tags = append(tags, tag)
		}
		article.Tags = tags

		// 寻找文章拥有的评论
		row3, err := utils.GetConn().Query("SELECT c.id, c.content, c.createdAt, c.ArticleId, c.authorId FROM Comment c WHERE c.ArticleId=?", article.ID)
		if err != nil {
			fmt.Println("error:", err)
		}
		comments := []CommentList{}
		for row3.Next() {
			comment := CommentList{}
			err = row3.Scan(&comment.ID, &comment.Content, &comment.ArticleID, &comment.CreatorID)
			comment.Creator = *GetUserListByID(comment.CreatorID)
			comments = append(comments, comment)
		}
		article.Comments = comments
	}

	return &article
}

func GetArticleByUserID(id int) []ArticleList {
	articles := []ArticleList{}
	row, err := utils.GetConn().Query("SELECT a.id FROM Articles a WHERE a.authorId=?", id)
	if err != nil {
		fmt.Println("error:", err)
	}
	for row.Next() {
		articleId := -1
		err = row.Scan(&articleId)
		articles = append(articles, *GetArticleByID(articleId))
	}
	return articles
}

func UpdateArticleByID(article Article) bool {
	stmt, err := utils.GetConn().Prepare("update user set title=?, content=?, createdAt=?, updatedAt=?, authorId=? where id=?")
	if err != nil {
		fmt.Println("error:", err)
	}
	_, err = stmt.Exec(article.Title, article.Content, article.CreatedAt, time.Now(), article.Author, article.ID)
	if err != nil {
		panic(err)
	}

	return true
}

func GetArticlesByTag(tag string) []ArticleList {
	articles := []ArticleList{}

	// 先从tag表中找到tagid
	row, err := utils.GetConn().Query("SELECT t.id FROM Tags t WHERE t.content=?", tag)
	if err != nil {
		fmt.Println("error:", err)
	}
	for row.Next() {
		tagId := -1
		err = row.Scan(&tagId)
		// 再从postTag表中根据tagid找到对应的articleid
		row2, err := utils.GetConn().Query("SELECT p.ArticleId FROM postTags p WHERE p.TagId=?", tagId)
		if err != nil {
			fmt.Println("error:", err)
		}
		for row2.Next() {
			articleId := -1
			err = row2.Scan(&articleId)
			articles = append(articles, *GetArticleByID(articleId))
		}
	}
	return articles
}

func GetArticlesPerPage(pageNum int) ArticlePage {
	articlePage := ArticlePage{}
	allArticles := GetAllArticles()
	articlePage.Number = len(allArticles)
	i := 0
	for range allArticles {
		if i >= (pageNum-1)*6 && i < pageNum * 6 {
			articlePage.Articles = append(articlePage.Articles, allArticles[i])
		}
		i = i + 1
	}
	return articlePage
}
