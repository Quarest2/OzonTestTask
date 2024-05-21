package OzonTestTask

type Post struct {
	ID              uint32
	Author          string
	CommentsAllowed bool
}

type Comment struct {
	ID              uint32
	Author          string
	ParentPostID    uint32
	ParentCommentID uint32
}
