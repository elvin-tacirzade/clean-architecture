package repositories

import (
	"github.com/elvin-tacirzade/clean-architecture/pkg/models"
	"github.com/jmoiron/sqlx"
)

type (
	UsersRepository interface {
		GetAllUsers() ([]models.Users, error)
		FindById(id string) (*models.Users, error)
		InsertUser(user *models.Users) error
		DeleteUser(id string) error
	}

	usersRepository struct {
		DB *sqlx.DB
	}
)

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		DB: db,
	}
}

func (u *usersRepository) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	err := u.DB.Select(&users, "SELECT public.users.id, public.users.name, public.users.email, public.users.password, "+
		"public.users.created_at, public.users.updated_at from public.users order by id DESC")
	if err != nil {
		return nil, err
	}
	return users, err
}

func (u *usersRepository) FindById(id string) (*models.Users, error) {
	var user models.Users
	err := u.DB.Get(&user, "SELECT public.users.id, public.users.name, public.users.email, "+
		"public.users.password, public.users.created_at, public.users.updated_at from public.users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersRepository) InsertUser(user *models.Users) error {
	_, err := u.DB.NamedExec("INSERT INTO public.users (name, email, password, created_at, updated_at)"+
		" VALUES (:name, :email, :password, :created_at, :updated_at)", user)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepository) DeleteUser(id string) error {
	_, err := u.DB.Exec("DELETE from public.users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
