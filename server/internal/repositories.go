// /home/ec2-user/calendar-reservation-web-app/server/internal/repositories.go
package internal

import (
    "database/sql"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) Create(user *User) error {
    // ユーザー作成ロジック
    // DBに接続してINSERTクエリを実行
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
    // メールアドレスからユーザーを取得するロジック
    // DBに接続してSELECTクエリを実行
}

// 他のメソッドも実装