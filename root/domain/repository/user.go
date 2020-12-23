package domain

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	Find(ctx context.Context, userID int) (*model.User, error)
}

type UserRepository struct {
	db sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{*db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Query("")
}
