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

	// // drop table if drop table, rows is defined
	// rows, err := db.Query("DROP TABLE events")
	// if err != nil {
	// 	log.Print("miss drop table!!")
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// create table //TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	rows, err := db.Query(`CREATE TABLE IF NOT EXISTS events (
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			title VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL,
			date VARCHAR(16) NOT NULL,
			hour VARCHAR(16) NOT NULL,
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
		"PUT",
		"DELETE",
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