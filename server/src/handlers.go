package calendar

import (
    "database/sql"
    "net/http"
    "time"
	"strconv"

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

    var event Event
    if err := c.BindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := createEvent(db, &event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, event)
}

func getEvents(db *sql.DB, date time.Time) ([]Event, error) {
    rows, err := db.Query("SELECT id, title, user, date, hour, created_at, updated_at FROM events WHERE date = $1", date)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var events []Event
    for rows.Next() {
        var event Event
        if err := rows.Scan(&event.ID, &event.Title, &event.User, &event.Date, &event.Hour, &event.CreatedAt, &event.UpdatedAt); err != nil {
            return nil, err
        }
        events = append(events, event)
    }

    return events, nil
}

func createEvent(db *sql.DB, event *Event) error {
    event.CreatedAt = time.Now()
    event.UpdatedAt = time.Now()

    _, err := db.Exec("INSERT INTO events (title, user, date, hour, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
        event.Title, event.User, event.Date, event.Hour, event.CreatedAt, event.UpdatedAt)

    return err
}

func UpdateEvent(c *gin.Context) {
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

    var event Event
    if err := c.BindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    event.ID = id

    if err := updateEvent(db, &event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, event)
}

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

func updateEvent(db *sql.DB, event *Event) error {
    event.UpdatedAt = time.Now()

    _, err := db.Exec("UPDATE events SET title = $1, user = $2, date = $3, hour = $4, updated_at = $5 WHERE id = $6",
        event.Title, event.User, event.Date, event.Hour, event.UpdatedAt, event.ID)

    return err
}

func deleteEvent(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM events WHERE id = $1", id)
    return err
}