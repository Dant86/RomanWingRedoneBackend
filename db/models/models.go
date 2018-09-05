package models

import "time"

type User struct {
    ID int
    FirstName string
    LastName string
    Email string
    IsAdmin bool
}

type Article struct {
    ID int
    Title string
    Description string
    CreatorID string
    Body string
    ThumbnailUrl string
    DateCreated time.Time
    IsAuthorized bool
}

type Event struct {
    ID int
    EventName string
    EventDescription string
    Date time.Time
    Location string
}

type Error struct {
    Message string
}

type ID struct {
    ID int
}
