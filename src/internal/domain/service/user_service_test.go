package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"

	"project03/internal/domain/models"
)

type userRepoStub struct {
	users     map[string]models.User
	createErr error
}

func newUserRepoStub() *userRepoStub {
	return &userRepoStub{
		users: make(map[string]models.User),
	}
}

func (r *userRepoStub) Create(_ context.Context, user models.User) error {
	if r.createErr != nil {
		return r.createErr
	}

	if _, exists := r.users[user.Login]; exists {
		return models.ErrLoginAlreadyTaken
	}

	r.users[user.Login] = user
	return nil
}

func (r *userRepoStub) GetByLogin(_ context.Context, login string) (models.User, error) {
	user, ok := r.users[login]
	if !ok {
		return models.User{}, models.ErrUserNotFound
	}
	return user, nil
}

func (r *userRepoStub) GetByID(_ context.Context, id string) (models.User, error) {
	for _, user := range r.users {
		if user.ID.String() == id {
			return user, nil
		}
	}
	return models.User{}, models.ErrUserNotFound
}

func TestUserServiceRegister(t *testing.T) {
	repo := newUserRepoStub()
	svc := NewUserService(repo)

	user, err := svc.Register(context.Background(), "alice", "securepass123")
	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if user.ID == uuid.Nil {
		t.Fatal("expected generated user UUID")
	}
	if user.Login != "alice" {
		t.Fatalf("login = %q, want alice", user.Login)
	}
	if user.PasswordHash != "" {
		t.Fatal("Register() must not return the password hash")
	}

	stored, err := repo.GetByLogin(context.Background(), "alice")
	if err != nil {
		t.Fatalf("GetByLogin() error = %v", err)
	}
	if stored.PasswordHash == "" || stored.PasswordHash == "securepass123" {
		t.Fatal("expected a hashed password in repository")
	}
}

func TestUserServiceRegisterValidation(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		password string
		wantErr  error
	}{
		{
			name:     "empty login",
			login:    "   ",
			password: "securepass123",
			wantErr:  ErrInvalidLogin,
		},
		{
			name:     "short password",
			login:    "alice",
			password: "123",
			wantErr:  ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewUserService(newUserRepoStub())

			_, err := svc.Register(context.Background(), tt.login, tt.password)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Register() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserServiceRegisterDuplicateLogin(t *testing.T) {
	repo := newUserRepoStub()
	svc := NewUserService(repo)

	_, err := svc.Register(context.Background(), "alice", "securepass123")
	if err != nil {
		t.Fatalf("first Register() error = %v", err)
	}

	_, err = svc.Register(context.Background(), "alice", "anotherpass123")
	if !errors.Is(err, models.ErrLoginAlreadyTaken) {
		t.Fatalf("second Register() error = %v, want ErrLoginAlreadyTaken", err)
	}
}

func TestUserServiceRepositoryError(t *testing.T) {
	repo := newUserRepoStub()
	repo.createErr = errors.New("database unavailable")

	svc := NewUserService(repo)
	_, err := svc.Register(context.Background(), "alice", "securepass123")
	if err == nil {
		t.Fatal("expected repository error")
	}
}
