package calendar

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/joho/godotenv" // .envファイルの読み込みに必要
    "os"
	"strconv"
)

func NewDB() (*sql.DB, error) {
    err := godotenv.Load() // .envファイルの読み込み
    if err != nil {
        return nil, err
    }

    host := os.Getenv("DB_HOST")
    port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}