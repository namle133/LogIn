package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/namle133/LogIn.git/LogIn_Project/database"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserService_SignIn(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c  context.Context
		ui *domain.UserInit
	}
	tests := []struct {
		fields  fields
		args    args
		want    *domain.Claims
		wantErr bool
	}{
		//case success
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "admin",
					Password: "admin1234",
					Email:    "admin@gmail.com"},
			},
			want: &domain.Claims{
				Username:       "admin",
				Password:       string(hash("admin1234")),
				Email:          "admin@gmail.com",
				StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(5 * time.Minute).Unix()},
			},
			wantErr: false,
		},

		//case failed
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "",
					Password: "Namle311",
					Email:    "Namle@gmail.com"},
			},
			want: &domain.Claims{
				Username:       "",
				Password:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		_, err := us.SignIn(tt.args.c, tt.args.ui)
		if (err != nil) != tt.wantErr {
			t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c context.Context
		u *domain.UserInit
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		//case success
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "Namle",
					Password: "Namle1234",
					Email:    "Namle@gmail.com"},
			},
			wantErr: false,
		},

		//case failed
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "",
					Password: "31231423",
					Email:    ""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.CreateUser(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
			t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_UserAdmin(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	tests := []struct {
		fields  fields
		wantErr bool
	}{
		{
			fields:  fields{Db: database.ConnectDatabase()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.UserAdmin(); (err != nil) != tt.wantErr {
			t.Errorf("UserAdmin() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_CheckRowToken(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.CheckRowToken(tt.args.c); (err != nil) != tt.wantErr {
			t.Errorf("CheckRowToken() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_LogOut(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{Db: database.ConnectDatabase()},
			args: args{
				c: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.LogOut(tt.args.c); (err != nil) != tt.wantErr {
			t.Errorf("LogOut() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
