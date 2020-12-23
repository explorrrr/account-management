package domain


type UserService struct {
	CreateUser(ctx context.Context, userName string, password string)
}


