// /home/ec2-user/calendar-reservation-web-app/server/internal/models/models.go
package models

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
    Role     string
}

type ReservationType struct {
    ID        int
    Name      string
    CreatedBy int
}

type Reservation struct {
    ID          int
    Title       string
    Description string
    StartTime   time.Time
    EndTime     time.Time
    TypeID      string
    UserID      int
	Fee 		int
}