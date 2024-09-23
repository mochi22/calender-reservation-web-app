package main

import (
	"log"
	// "fmt"
	// "time"

    "github.com/gin-gonic/gin"
	"github.com/mochi22/calender-reservation-web-app/server/src/db"
	// "github.com/mochi22/calender-reservation-web-app/server/src"
)

// func main() {
//     r := gin.Default()

//     // ルーティングの設定
//     r.GET("/events", GetEvents)
//     r.POST("/events", CreateEvent)
//     r.PUT("/events/:id", UpdateEvent)
//     r.DELETE("/events/:id", DeleteEvent)

//     r.Run(":8080")
// }

// func testGetEvents(c *gin.Context) {
//     date, _ := time.Parse("2006-01-02", "2023-06-01")
//     events, err := GetEvents(c, date)
//     if err != nil {
//         log.Println(err)
//         return
//     }

//     for _, event := range events {
//         fmt.Printf("ID: %d, Title: %s, User: %s, Date: %s, Hour: %s\n", event.ID, event.Title, event.User, event.Date.Format("2006-01-02"), event.Hour)
//     }
// }

// func testCreateEvent(c *gin.Context) {
//     event := Event{
//         Title: "テストイベント",
//         User:  "テストユーザー",
//         Date:  time.Now(),
//         Hour:  "10:00",
//     }

//     if err := CreateEvent(c, &event); err != nil {
//         log.Println(err)
//         return
//     }

//     fmt.Printf("イベントを作成しました: %+v\n", event)
// }

// func testUpdateEvent(c *gin.Context) {
//     event := Event{
//         ID:    1, // 更新する予定のID
//         Title: "更新後のタイトル",
//         User:  "更新後のユーザー",
//         Date:  time.Now(),
//         Hour:  "14:00",
//     }

//     if err := UpdateEvent(c, &event); err != nil {
//         log.Println(err)
//         return
//     }

//     fmt.Printf("イベントを更新しました: %+v\n", event)
// }

// func testDeleteEvent(c *gin.Context) {
//     id := 1 // 削除する予定のID

//     if err := DeleteEvent(c, id); err != nil {
//         log.Println(err)
//         return
//     }

//     fmt.Println("イベントを削除しました")
// }

func main() {
	// connect to DB
    db, err := db.NewDB()
    if err != nil {
		log.Print("connecting DB error!!")
		log.Fatal(err)
    }
    defer db.Close()

	// usersテーブルからデータを取得
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		log.Print("select DB error!!")
		log.Fatal(err)
	}
	log.Print("select * from events:", rows)
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