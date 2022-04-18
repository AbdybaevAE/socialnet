package users

type UserService interface {
	Register(args *RegArgs)
	Login()
	CreatePost()
	GetFeed()
	LikePost()
	CommentPost()
}
type RegArgs struct {
	FirstName  string
	LastName   string
	Email      string
	CityId     int
	LanguageId int
}
