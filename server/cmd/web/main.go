// /home/ec2-user/calendar-reservation-web-app/server/cmd/web/main.go
package main

import (
    "log"
    "net/http"

    "github.com/your-username/calendar-reservation-web-app/server/internal"
)

func main() {
    // データベース接続の初期化
    db, err := internal.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // リポジトリの初期化
    userRepo := &internal.UserRepository{DB: db}

    // サービスの初期化
    authService := &internal.AuthService{UserRepo: userRepo}

    // ハンドラの初期化
    authHandler := &internal.AuthHandler{AuthService: authService}

    // ルーティングの設定
    http.HandleFunc("/signup", authHandler.SignUpHandler)
    http.HandleFunc("/signin", authHandler.SignInHandler)

    // Webサーバーの起動
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}