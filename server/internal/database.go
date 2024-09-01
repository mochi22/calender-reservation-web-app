// /home/ec2-user/calendar-reservation-web-app/server/internal/repositories/database.go
package repositories

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func connectDB() (*sql.DB, error) {
    // Amazon RDSのエンドポイントとポート番号を設定
    // database-1.cluster-chyjw5kivnix.ap-northeast-1.rds.amazonaws.com	ライター
    // database-1.cluster-ro-chyjw5kivnix.ap-northeast-1.rds.amazonaws.com	 リーダー
    endpoint := "database-1.cluster-chyjw5kivnix.ap-northeast-1.rds.amazonaws.com"
    port := 5432

    // データソース名(DSN)を作成
    dsn := fmt.Sprintf("host=%s port=%d user=postgres password=Saqwedcxz22! dbname=calendarapp sslmode=disable", endpoint, port)

    // データベースに接続
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    return db, nil
}