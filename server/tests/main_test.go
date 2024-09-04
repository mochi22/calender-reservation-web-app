package calendar

import (
    "github.com/mochi22/calender-reservation-web-app/server/src"
	// "calender-reservation-web-app/server/src"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // ルーティングの設定
    r.GET("/events", src.GetEvents)
    r.POST("/events", src.CreateEvent)
    r.PUT("/events/:id", src.UpdateEvent)
    r.DELETE("/events/:id", src.DeleteEvent)

    r.Run(":8080")
}

func testGetEvents(c *gin.Context) {
    date, _ := time.Parse("2006-01-02", "2023-06-01")
    events, err := src.GetEvents(c, date)
    if err != nil {
        log.Println(err)
        return
    }

    for _, event := range events {
        fmt.Printf("ID: %d, Title: %s, User: %s, Date: %s, Hour: %s\n", event.ID, event.Title, event.User, event.Date.Format("2006-01-02"), event.Hour)
    }
}

func testCreateEvent(c *gin.Context) {
    event := src.Event{
        Title: "テストイベント",
        User:  "テストユーザー",
        Date:  time.Now(),
        Hour:  "10:00",
    }

    if err := src.CreateEvent(c, &event); err != nil {
        log.Println(err)
        return
    }

    fmt.Printf("イベントを作成しました: %+v\n", event)
}

func testUpdateEvent(c *gin.Context) {
    event := src.Event{
        ID:    1, // 更新する予定のID
        Title: "更新後のタイトル",
        User:  "更新後のユーザー",
        Date:  time.Now(),
        Hour:  "14:00",
    }

    if err := src.UpdateEvent(c, &event); err != nil {
        log.Println(err)
        return
    }

    fmt.Printf("イベントを更新しました: %+v\n", event)
}

func testDeleteEvent(c *gin.Context) {
    id := 1 // 削除する予定のID

    if err := src.DeleteEvent(c, id); err != nil {
        log.Println(err)
        return
    }

    fmt.Println("イベントを削除しました")
}