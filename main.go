package main

import (
    "fmt"
    "RomanWingBackend/db/dao"
    "RomanWingBackend/db/models"
    "RomanWingBackend/db/utils"
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
func CreateUser(fName, lName, email, pword *C.char) *C.char {
    newFName := C.GoString(fName)
    newLName := C.GoString(lName)
    newEmail := C.GoString(email)
    newPword := C.GoString(pword)
    id, err := dao.CreateUser(newFName, newLName, newEmail, newPword)
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
func GetUserByEmail(email *C.char) *C.char {
    newEmail := C.GoString(email)
    u, err := dao.GetUserByEmail(newEmail)
    if err != nil { return serializeError(err) }
    serialized, _ := json.Marshal(u)
    return C.CString(string(serialized))
}

//export ValidateUser
func ValidateUser(email, pword *C.char) *C.char {
    fmt.Println("I'm here")
    newEmail := C.GoString(email)
    newPword := C.GoString(pword)
    fmt.Println(newEmail + " " + newPword)
    err := dao.ValidateUser(newEmail, newPword)
    if err != nil { return serializeError(err) }
    return C.CString("{}")
}

//export CreateArticle
func CreateArticle(creatorId int, title, description, body,
                   thumbnailUrl *C.char) *C.char {
    newTitle := C.GoString(title)
    newDesc := C.GoString(description)
    newBody := C.GoString(body)
    newURL := C.GoString(thumbnailUrl)
    id, err := dao.CreateArticle(creatorId, newTitle, newDesc,
                                 newBody, newURL)
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

//export UpdateArticleBody
func UpdateArticleBody(id int, body *C.char) *C.char {
    newBody := C.GoString(body)
    err := dao.UpdateArticleBody(id, newBody)
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

//export Get12MostRecentArticles
func Get12MostRecentArticles() *C.char {
    result := "["
    articles, err := dao.Get12MostRecentArticles()
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
func CreateEvent(name, description, date, location *C.char) *C.char {
    newName := C.GoString(name)
    newDesc := C.GoString(description)
    newDate := C.GoString(date)
    newLoc := C.GoString(location)
    err := dao.CreateEvent(newName, newDesc, newDate, newLoc)
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

//export Migrate
func Migrate() {
    tNames := []string{"users-table", "user-auth-table", "articles-table",
                       "saved-articles-table", "events-table"}
    utils.Migrate("root", "dbpassword", "db/migrations.sql", tNames)
}

func main() {}
