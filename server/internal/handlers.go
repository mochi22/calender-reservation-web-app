// /home/ec2-user/calendar-reservation-web-app/server/internal/handlers.go
package internal

import (
    "encoding/json"
    "net/http"
)

type AuthHandler struct {
    AuthService *AuthService
}

func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
    // リクエストボディからユーザー情報を取得
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // サービスレイヤーのSignUpメソッドを呼び出す
    err = h.AuthService.SignUp(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
    // リクエストボディからメールアドレスとパスワードを取得
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // サービスレイヤーのSignInメソッドを呼び出す
    user, err := h.AuthService.SignIn(credentials.Email, credentials.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    // ユーザー情報をレスポンスとして返す
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}