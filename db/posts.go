package db

import (
	"OzonTestTask/OzonTestTask/graph/model"
	"database/sql"
)

func CreatePost(db *sql.DB, newPost model.NewPost) (int64, bool) {
	query := "INSERT INTO posts (author, commentsAllowed) VALUES ($1, $2) RETURNING id"
	res, ok := MustExec(db, query, newPost.Author, newPost.CommentsAllowed)
	if !ok {
		return -1, false
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, false
	}
	return id, true
}

func DeletePost(db *sql.DB, postId int64) bool {
	query := "DELETE FROM posts WHERE id = $1"
	_, ok := MustExec(db, query, postId)
	if !ok {
		return false
	}
	return true
}

func DeleteAllPosts(db *sql.DB) bool {
	query := "DELETE FROM posts"
	_, ok := MustExec(db, query)
	return ok
}

func GetAllPosts(db *sql.DB) ([]*model.Post, bool) {
	var posts []*model.Post
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, false
	}
	var post model.Post
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Author, &post.CommentsAllowed); err != nil {
			return nil, false
		}
		posts = append(posts, &post)
	}
	return posts, true
}

func GetPost(db *sql.DB, postId int64) (*model.Post, bool) {
	var post model.Post
	row, err := db.Query("SELECT * FROM posts WHERE id = $1", postId)
	if err != nil {
		return nil, false
	}
	for row.Next() {
		if err := row.Scan(&post.ID, &post.Author, &post.CommentsAllowed); err != nil {
			return nil, false
		}
	}
	return &post, true
}
