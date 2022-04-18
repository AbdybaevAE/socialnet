package users

type UserService interface {
	Register(args *RegArgs) (err error)
	Login(args *LoginArgs) (err error)
	CreatePost(args *CreatePostArgs) (err error)
	GetFeed() (items []PostItemRes, err error)
	LikePost(args *LikePostArgs) (err error)
	CommentPost(args *CommentPostArgs) (err error)
}
type RegArgs struct {
	FirstName  string
	LastName   string
	Email      string
	Password   string
	CityId     int
	LanguageId int
}
type LoginArgs struct {
	Email    string
	Password string
}
type CreatePostArgs struct {
	Title   string
	Content string
}
type PostItemRes struct {
	Id      int
	Title   string
	Content string
}
type CommentPostArgs struct {
	PostId  int
	Content string
}
type LikePostArgs struct {
	PostId int
}
