package main

import (
    "RomanWingBackend/db/dao"
    "RomanWingBackend/db/models"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "C"
)

func serializeError(err error) *C.char {
    err_struct := models.Error{Message: err.Error()}
    serialized, _ := json.Marshal(err_struct)
    return C.CString(string(serialized))
}

//export CreateUser
func CreateUser(fName, lName, email, pword string) *C.char {
    id, err := dao.CreateUser(fName, lName, email, pword)
    if err != nil {
        return serializeError(err)
    }
    id_struct := models.ID{ID: id}
    serialized, _ := json.Marshal(id_struct)
    return C.CString(string(serialized))
}

//export GetUser
func GetUser(id int) *C.char {
    u, err := dao.GetUser(id)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(u)
    return C.CString(string(serialized))
}

//export GetUserByEmail
func GetUserByEmail(email string) *C.char {
    u, err := dao.GetUserByEmail(email)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(u)
    return C.CString(string(serialized))
}

//export ValidateUser
func ValidateUser(email string, pword string) *C.char {
    err := dao.ValidateUser(email, pword)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export CreateArticle
func CreateArticle(creatorId int, title, description, body,
                   thumbnailUrl string) *C.char {
    id, err := dao.CreateArticle(creatorId, title, description,
                                 body, thumbnailUrl)
    if err != nil { return serializeError(err) }
    id_struct := models.ID{ID: id}
    serialized, _ := json.Marshal(id_struct)
    return C.CString(string(serialized))
}

//export GetArticle
func GetArticle(id int) *C.char {
    a, err := dao.GetArticle(id)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(a)
    return C.CString(string(serialized))
}

//export DeleteArticle
func DeleteArticle(id int) *C.char {
    err := dao.DeleteArticle(id)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export ApproveArticle
func ApproveArticle(id int) *C.char {
    err := dao.ApproveArticle(id)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export GetArticlesFromUser
func GetArticlesFromUser(userId int) *C.char {
    result := "["
    articles, err := dao.GetArticlesFromUser(userId)
    if err != nil { return serializeError(err) }
    for ix, article := range articles {
        serialized, _ := json.Marshal(article)
        result += string(serialized)
        if ix < len(articles) - 1 {
            result += ", "
        }
    }
    return C.CString(result + "]")
}

//export GetApprovedArticles
func GetApprovedArticles() *C.char {
    result := "["
    articles, err := dao.GetApprovedArticles()
    if err != nil { return serializeError(err) }
    for ix, article := range articles {
        serialized, _ := json.Marshal(article)
        result += string(serialized)
        if ix < len(articles) - 1 {
            result += ", "
        }
    }
    return C.CString(result + "]")
}

//export Get10MostRecentArticles
func Get10MostRecentArticles() string {
    result := "["
    articles, err := dao.Get10Get10MostRecentArticles()
    if err != nil { return serializeError(err) }
    for ix, article := range articles {
        serialized, _ := json.Marshal(article)
        result += string(serialized)
        if ix < len(articles) - 1 {
            result += ", "
        }
    }
    return C.CString(result + "]")
}

//export GetArticleAuthor
func GetArticleAuthor(articleId int) *C.char {
    u, err := dao.GetArticleAuthor(articleId)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(u)
    return C.CString(string(serialized))
}

//export SaveArticle
func SaveArticle(userId, articleId int) *C.char {
    err := dao.SaveArticle(userId, articleId)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export GetSavedArticles
func GetSavedArticles(userId int) *C.char {
    result := "["
    articles, err := dao.GetSavedArticles(userId)
    if err != nil { return serializeError(err) }
    for ix, article := range articles {
        serialized, _ := json.Marshal(article)
        result += string(serialized)
        if ix < len(articles) - 1 {
            result += ", "
        }
    }
    return C.CString(result + "]")
}

//export CreateEvent
func CreateEvent(name, description, date, location string) *C.char {
    err := dao.CreateEvent(name, description, date, location)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export GetEvent
func GetEvent(id int) *C.char {
    e, err := dao.GetEvent(id)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(e)
    return C.CString(string(serialized))
}

//export GetFutureEvents
func GetFutureEvents() *C.char {
    events, err := dao.GetFutureEvents()
    if err != nil { return serializeError(err) }
    result := "["
    for ix, event := range events {
        serialized, _ := json.Marshal(event)
        result += string(serialized)
        if ix < len(events) - 1 {
            result += ", "
        }
    }
    return C.CString(result + "]")
}

func main() {}
