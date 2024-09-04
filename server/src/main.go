package calendar

import (
	"log"

    "github.com/gin-gonic/gin"
)

func main() {
    db, err := NewDB()
    if err != nil {
		log.Fatal(err)
    }
    defer db.Close()

	// usersテーブルからデータを取得
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

    r := gin.Default()

    // routing
    r.GET("/events", GetEvents)
    r.POST("/events", CreateEvent)
    r.PUT("/events/:id", UpdateEvent)
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