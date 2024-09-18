package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    "log"
    "fmt"

    "github.com/gin-gonic/gin"
)




func TestCreateEvent(t *testing.T) {
    // テストデータの作成
    event := Event{
        ID: 1,
        Title: "Test Event",
        User: "Test User",
        Date: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
        Hour: "10:00",
        CreatedAt: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
        UpdatedAt: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
    }

    // リクエストボディの作成
    jsonData, err := json.Marshal(event)
    if err != nil {
        t.Fatalf("Failed to marshal event data: %v", err)
    }

    // テストリクエストの作成
    req, err := http.NewRequest("POST", "/events", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // log.Print("bytes.NewBuffer(jsonData):", bytes.NewBuffer(jsonData)) //{"id":1,"title":"Test Event","username":"Test User","date":"2023-06-01T00:00:00Z","hour":"10:00","created_at":"2023-06-01T00:00:00Z","updated_at":"2023-06-01T00:00:00Z"}

    // Ginルーターの初期化とハンドラの設定
    r := gin.New()
    r.POST("/events", CreateEvent)

    // テストリクエストの実行
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // レスポンスの検証
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
    }

    var createdEvent Event
    err = json.Unmarshal(w.Body.Bytes(), &createdEvent)
    if err != nil {
        t.Errorf("Failed to unmarshal response body: %v", err)
    }
}

func TestGetEvents(t *testing.T) {
    // テストデータの作成
    event := Event{
        ID: 1,
        Title: "Test Event",
        User:  "Test User",
        Date:  time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
        Hour:  "10:00",
    }

    // データベースに予定を挿入する

    // テストリクエストの作成
    req, err := http.NewRequest("GET", "/events?date=2023-06-01", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // Ginルーターの初期化とハンドラの設定
    r := gin.New()
    r.GET("/events", GetEvents)

    // テストリクエストの実行
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // レスポンスの検証
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    var events []Event
    err = json.Unmarshal(w.Body.Bytes(), &events)
    if err != nil {
        t.Errorf("Failed to unmarshal response body: %v", err)
    }

    log.Print("event:", event.User)
    log.Print(events)

    if len(events) != 1 || events[0].Title != event.Title || events[0].User != event.User || events[0].Date != event.Date || events[0].Hour != event.Hour {
        t.Errorf("Retrieved events do not match the expected events")
    }

    // // レコードの削除
    // _, err = db.Exec("DELETE FROM events WHERE id = $1", createdEvent.ID)
    // if err != nil {
    //     t.Errorf("Failed to delete test event: %v", err)
    // }
}


func TestDeleteEvents(t *testing.T) {
    // テストデータの作成
    event := Event{
        ID: 1,
        Title: "Test Event",
        User:  "Test User",
        Date:  time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
        Hour:  "10:00",
    }

    // テストリクエストの作成
    req, err := http.NewRequest("DELETE", fmt.Sprintf("/events/%d", event.ID), nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // Ginルーターの初期化とハンドラの設定
    r := gin.New()
    r.DELETE("/events/:id", DeleteEvent)

    // テストリクエストの実行
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    fmt.Print("req:", req)

    // レスポンスの検証
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    var response struct {
        Message string `json:"message"`
    }
    err = json.Unmarshal(w.Body.Bytes(), &response)
    if err != nil {
        t.Errorf("Failed to unmarshal response body: %v", err)
    }

    if response.Message != "Event deleted successfully" {
        t.Errorf("Unexpected response message: %s", response.Message)
    }

    // レコードが削除されたことを確認
    // var count int
    // err = db.QueryRow("SELECT COUNT(*) FROM events WHERE id = $1", eventID).Scan(&count)
    // if err != nil {
    //     t.Errorf("Failed to check event count: %v", err)
    // }
    // if count != 0 {
    //     t.Errorf("Event was not deleted")
    // }
}

// package main

// import (
//     // "bytes"
//     "encoding/json"
//     "net/http"
//     "net/http/httptest"
//     "testing"
//     "time"
//     "database/sql"

//     "github.com/gin-gonic/gin"
//     "github.com/stretchr/testify/assert"
// )

// // モックデータベース
// type mockDB struct {
//     events []Event
// }

// // func (m *mockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
// //     // クエリに応じてモックデータを返すロジックを実装
// //     return nil, nil
// // }

// // func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
// //     // クエリに応じてモックデータを更新するロジックを実装
// //     return nil, nil
// // }

// func TestGetEvents(t *testing.T) {
//     // モックデータベースの初期化
//     mockDB := &mockDB{
//         events: []Event{
//             {ID: 1, Title: "Meeting", User: "John", Date: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), Hour: "10:00"},
//             {ID: 2, Title: "Break", User: "Jane", Date: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), Hour: "15:00"},
//         },
//     }

//     // モックデータベースを使ってNewDBを置き換える
//     // oldNewDB := NewDB
//     // NewDB := func() (*sql.DB, error) {
//     //     return mockDB, nil
//     // }
//     // defer func() { NewDB = oldNewDB }()

//     // テストケース
//     w := httptest.NewRecorder()
//     c, _ := gin.CreateTestContext(w)
//     c.Request, _ = http.NewRequest("GET", "/events?date=2023-06-01", nil)

//     GetEvents(c)

//     assert.Equal(t, http.StatusOK, w.Code)

//     var events []Event
//     err := json.Unmarshal(w.Body.Bytes(), &events)
//     assert.NoError(t, err)
//     assert.Len(t, events, 2)
// }

// // 他のテストケースも追加