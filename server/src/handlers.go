package main

import (
    "database/sql"
    "net/http"
    "time"
	"strconv"
    "log"

    "github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
    db, err := NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()

    date, err := time.Parse("2006-01-02", c.Query("date"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
        return
    }

    events, err := getEvents(db, date)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
    db, err := NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()
    log.Print("ssss!!")
    var event Event

    if err := c.BindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"bad request error": err.Error()})
        log.Fatal(err)
        return
    }
    // イベントデータの処理
    log.Print("Received event data:", event)

    if err := createEvent(db, &event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error createEvent": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, event)
}

func getEvents(db *sql.DB, date time.Time) ([]Event, error) {
    rows, err := db.Query("SELECT id, title, username, date, created_at, updated_at FROM events WHERE date = $1", date)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var events []Event
    for rows.Next() {
        var event Event
        if err := rows.Scan(&event.ID, &event.Title, &event.Username, &event.Date, &event.CreatedAt, &event.UpdatedAt); err != nil {
            return nil, err
        }
        events = append(events, event)
    }

    return events, nil
}

func createEvent(db *sql.DB, event *Event) error {
    event.CreatedAt = time.Now()
    event.UpdatedAt = time.Now()

    _, err := db.Exec("INSERT INTO events (id, title, username, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
        event.ID, event.Title, event.Username, event.Date, event.CreatedAt, event.UpdatedAt)

    if err != nil {
        // c.JSON(http.StatusInternalServerError, gin.H{"error when insert": err.Error()})
        log.Fatal(err)
        return err
    }
    return err
}

// func UpdateEvent(c *gin.Context) {
//     db, err := NewDB()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     defer db.Close()

//     id, err := strconv.Atoi(c.Param("id"))
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
//         return
//     }

//     var event Event
//     if err := c.BindJSON(&event); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     // event.ID = id

//     if err := updateEvent(db, &event); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, event)
// }

func DeleteEvent(c *gin.Context) {
    db, err := NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    if err := deleteEvent(db, id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// func updateEvent(db *sql.DB, event *Event) error {
//     event.UpdatedAt = time.Now()

//     _, err := db.Exec("UPDATE events SET title = $1, username = $2, date = $3, updated_at = $4 WHERE id = $5",
//         event.Title, event.Username, event.Date, event.UpdatedAt, event.ID)

//     return err
// }

func deleteEvent(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM events WHERE id = $1", id)
    return err
}