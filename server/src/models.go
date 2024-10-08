package main

import "time"

type Event struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Username  string    `json:"username"`
    Date      string    `json:"date"`
    Hour      string    `json:"hour"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


// CREATE TABLE events (
//     id SERIAL PRIMARY KEY,
//     title VARCHAR(255) NOT NULL,
//     username VARCHAR(255) NOT NULL,
//     date DATE NOT NULL,
//     hour VARCHAR(10) NOT NULL,
//     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
// );