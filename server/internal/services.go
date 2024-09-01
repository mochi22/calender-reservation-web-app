// /home/ec2-user/calendar-reservation-web-app/server/internal/services.go
package internal

import (
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    UserRepo *UserRepository
}

func (s *AuthService) SignUp(user *User) error {
    // ユーザー登録ロジック
    // パスワードのハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    // UserRepositoryを使ってユーザー作成
    err = s.UserRepo.Create(user)
    if err != nil {
        return err
    }

    return nil
}

func (s *AuthService) SignIn(email, password string) (*User, error) {
    // ユーザー認証ロジック
    // UserRepositoryを使ってユーザー取得
    user, err := s.UserRepo.FindByEmail(email)
    if err != nil {
        return nil, err
    }

    // パスワード比較
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, err
    }

    return user, nil
}