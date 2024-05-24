package db

import (
	"OzonTestTask/OzonTestTask/graph/model"
	"database/sql"
)

func CreateComment(db *sql.DB, newComment model.NewComment) (int64, bool) {
	query := "INSERT INTO comments (author, parentPost, parentComment) VALUES ($1, $2, $3) RETURNING id"
	res, ok := MustExec(db, query, newComment.Author, newComment.ParentPost, newComment.ParentComment)
	if !ok {
		return -1, false
	}
	idInt, err := res.LastInsertId()
	if err != nil {
		return -1, false
	}
	return idInt, true
}

func DeleteComent(db *sql.DB, commentId int64) bool {
	query := "DELETE FROM comments WHERE id = $1"
	_, ok := MustExec(db, query, commentId)
	if !ok {
		return false
	}
	return true
}

func DeleteAllComments(db *sql.DB) bool {
	query := "DELETE FROM comments"
	_, ok := MustExec(db, query)
	return ok
}

func GetCommentByPost(db *sql.DB, parentPost int64) ([]*model.Comment, bool) {
	var comments []*model.Comment
	row, err := db.Query("SELECT * FROM comments WHERE parentPost = $1", parentPost)
	if err != nil {
		return nil, false
	}
	var comment model.Comment
	for row.Next() {
		if err := row.Scan(&comment.ID, &comment.Author, &comment.ParentPost, &comment.ParentPost); err != nil {
			return nil, false
		}
		comments = append(comments, &comment)
	}
	return comments, true
}

func GetAllComments(db *sql.DB) ([]*model.Comment, bool) {
	var comments []*model.Comment
	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		return nil, false
	}
	var comment model.Comment
	for rows.Next() {
		if err := rows.Scan(&comment.ID, &comment.Author, &comment.ParentPost, &comment.ParentComment); err != nil {
			return nil, false
		}
		comments = append(comments, &comment)
	}
	return comments, true
}

func GetComment(db *sql.DB, commId int64) (*model.Comment, bool) {
	var comm model.Comment
	row, err := db.Query("SELECT * FROM comments WHERE id = $1", commId)
	if err != nil {
		return nil, false
	}
	for row.Next() {
		if err := row.Scan(&comm.ID, &comm.Author, &comm.ParentPost, &comm.ParentPost); err != nil {
			return nil, false
		}
	}
	return &comm, true
}
