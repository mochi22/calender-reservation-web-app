package main

import (
	"log"
	// "fmt"
	"time"

    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "github.com/mochi22/calender-reservation-web-app/server/src"
)

func main() {
	// connect to DB
    db, err := NewDB()
    if err != nil {
		log.Print("connecting DB error!!")
		log.Fatal(err)
    }
    defer db.Close()

	// drop table
	// rows, err := db.Query("DROP TABLE events")
	// if err != nil {
	// 	log.Print("miss create table!!")
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// create table
	rows, err = db.Query(`CREATE TABLE IF NOT EXISTS events (
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			title VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL,
			date DATE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		log.Print("miss create table!!")
		log.Fatal(err)
	}
	defer rows.Close()


	// // usersテーブルからデータを取得
	// rows, err := db.Query("SELECT * FROM events")
	// if err != nil {
	// 	log.Print("select DB error!!")
	// 	log.Fatal(err)
	// }
	// log.Print("select * from events:", rows)
	// defer rows.Close()

    r := gin.Default()

	// // CORSの設定
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"} // TypeScriptアプリケーションのオリジンを指定
	// r.Use(cors.New(config))
	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
	// アクセスを許可したいアクセス元
	AllowOrigins: []string{
		"http://localhost:3000",
		"http://localhost:3000/calender",
	},
	// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
	AllowMethods: []string{
		"POST",
		"GET",
		"OPTIONS",
	},
	// 許可したいHTTPリクエストヘッダ
	AllowHeaders: []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	},
	// cookieなどの情報を必要とするかどうか
	AllowCredentials: true,
	// preflightリクエストの結果をキャッシュする時間
	MaxAge: 24 * time.Hour,
	}))

    // routing
    r.GET("/events", GetEvents)
    r.POST("/events", CreateEvent)
    // r.PUT("/events/:id", UpdateEvent)
    r.DELETE("/events/:id", DeleteEvent)

    r.Run(":8080")
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type user struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Password string `json:"password"`
// 	Calender_ID int `json:"calender_id"`
// }

// var users = []user{
// 	{ID: 1, Name: "apple", Password: "pass1", Calender_ID: 10},
// 	{ID: 2, Name: "banana", Password: "pass22", Calender_ID: 20},
// 	{ID: 3, Name: "grape", Password: "pass333", Calender_ID: 10},
// }



// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
//     // CORS対応コード, react側がport3000を使用している。
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // 追加
// 	json.NewEncoder(w).Encode(users)
// }

// // ※Goではコードの記述順序は関係ないので、上に書いても下に書いても構いません。
// func main() {
// 	http.HandleFunc("/", getUsers)
// 	fmt.Println("Starting server at port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }