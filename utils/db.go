package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:limzhonglin@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	// Users
	users := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
        username VARCHAR(64) UNIQUE  NOT NULL,
        email VARCHAR(64) NOT NULL,
        password VARCHAR(64) NOT NULL,
        iconPath VARCHAR(64) NOT NULL
    );
    `
	_ ,err = DB.Exec(users)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}
	// UserRelations
	userRelations := `
    CREATE TABLE IF NOT EXISTS userRelations(
        UserId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE ,
        followerId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
    `

	_ ,err = DB.Exec(userRelations)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	// Tags
	tags := `
    CREATE TABLE IF NOT EXISTS Tags(
        content VARCHAR(64) NOT NULL PRIMARY KEY
    );
    `
	_ ,err = DB.Exec(tags)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	// Articles
	articles := `
    CREATE TABLE IF NOT EXISTS Aticles(
        id INTEGER  PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(64),
        content VARCHAR(64),
      	createdAt DATE,
      	updatedAt DATE,
      	authorId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
    `
	_ ,err = DB.Exec(articles)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	// Comments
	comments := `
    CREATE TABLE IF NOT EXISTS Comments(
        id INTEGER  PRIMARY KEY AUTO_INCREMENT,
        content VARCHAR(64),
      	createdAt DATE,
      	updatedAt DATE,
      	articleId INTEGER NOT NULL REFERENCES Articles(id) ON DELETE CASCADE ON UPDATE CASCADE,
      	authorId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
    `
	_ ,err = DB.Exec(comments)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	// postTags
	postTags := `
    CREATE TABLE IF NOT EXISTS postTags(
      	ArticleId INTEGER NOT NULL REFERENCES Articles(id) ON DELETE CASCADE ON UPDATE CASCADE,
      	TagId INTEGER NOT NULL REFERENCES Tags(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
    `
	_ ,err = DB.Exec(postTags)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
		//panic(err)
	}

	fmt.Println("Create database successfully!")
}

func GetConn() *sql.DB {
	return DB;
}

