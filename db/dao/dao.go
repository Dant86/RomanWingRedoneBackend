package dao

import (
    "RomanWingBackend/db/utils"
    "RomanWingBackend/db/models"
    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
)

func CreateUser(fName, lName, email, pword string) (int, error) {
    hashByte, _ := bcrypt.GenerateFromPassword([]byte(pword),
                                               bcrypt.DefaultCost)
    hashStr := string(hashByte)
    db := utils.OpenMySQL("root", "dbpassword")
    cmd1 := "INSERT INTO users (first_name, last_name, email, is_admin) " +
            "VALUES (?, ?, ?, FALSE)"
    cmd2 := "INSERT INTO user_auth (hash, user_id) VALUES (?, ?)"
    stmt1, _ := db.Prepare(cmd1)
    stmt2, _ := db.Prepare(cmd2)
    res, err := stmt1.Exec(fName, lName, email)
    if res == nil || err != nil { return -1, err }
    id, _ := res.LastInsertId()
    if err != nil { return -1, err }
    _, err = stmt2.Exec(hashStr, id)
    if err != nil { return -1, err }
    return int(id), nil
}

func GetUser(id int) (models.User, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, first_name, last_name, email, is_admin FROM users " +
           "WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    var u models.User
    err := stmt.QueryRow(id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email,
                                     &u.IsAdmin)
    if err != nil { return u, err }
    return u, nil
}

func GetUserByEmail(email string) (models.User, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, first_name, last_name, email, is_admin FROM users " +
           "WHERE email=?"
    stmt, _ := db.Prepare(cmd)
    var u models.User
    err := stmt.QueryRow(email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email,
                                     &u.IsAdmin)
    if err != nil { return u, err }
    return u, nil
}

func GetHash(userId int) (string, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT hash from user_auth WHERE user_id=?"
    stmt, _ := db.Prepare(cmd)
    var hash string
    err := stmt.QueryRow(userId).Scan(&hash)
    if err != nil { return "", err }
    return hash, nil
}

func ValidateUser(email string, pword string) error {
    u, err := GetUserByEmail(email)
    if err != nil { return err }
    id := u.ID
    h, err := GetHash(id)
    if err != nil { return err }
    err = bcrypt.CompareHashAndPassword([]byte(h), []byte(pword))
    if err != nil { return err }
    return nil
}

func UpdatePassword(userId int, currPword, newPword string) error {
    db := utils.OpenMySQL("root", "dbpassword")
    usr, err := GetUser(userId)
    if err != nil { return err }
    err := ValidateUser(usr.Email, currPword)
    if err != nil { return err }
    hashByte, _ := bcrypt.GenerateFromPassword([]byte(newPword),
                                               bcrypt.DefaultCost)
    hashStr := string(hashByte)
    cmd := "UPDATE user_auth SET hash=? WHERE user_id=?"
    stmt, _ := db.Prepare(cmd)
    _, err = stmt.Exec(hashStr, userId)
    if err != nil { return err }
    return nil
}

func CreateArticle(creatorId int, title, description, body,
                   thumbnailUrl string) (int, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := GetUser(creatorId)
    if err != nil { return -1, err }
    cmd := "INSERT INTO articles (title, description, creator_id, " +
           "body, thumbnail_url, date_created, is_authorized) " +
           "VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, FALSE)"
    stmt, _ := db.Prepare(cmd)
    res, err := stmt.Exec(title, description, creatorId, body, thumbnailUrl)
    if res == nil || err != nil { return -1, err }
    id, _ := res.LastInsertId()
    return int(id), nil
}

func GetArticle(id int) (models.Article, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, title, description, creator_id, body, thumbnail_url, " +
           "date_created, is_authorized from ARTICLES WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    var a models.Article
    err := stmt.QueryRow(id).Scan(&a.ID, &a.Title, &a.Description, &a.CreatorID,
                                  &a.Body, &a.ThumbnailUrl, &a.DateCreated,
                                  &a.IsAuthorized)
    if err != nil { return a, err }
    return a, nil
}

func DeleteArticle(id int) error {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "DELETE FROM articles WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    _, err := stmt.Exec(id)
    if err != nil { return err }
    return nil
}

func UpdateArticleBody(id int, body string) error {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "UPDATE articles SET body=? WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    _, err := stmt.Exec(body, id)
    if err != nil { return err }
    return nil
}

func ApproveArticle(id int) error {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "UPDATE articles SET is_authorized=TRUE WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    _, err := stmt.Exec(id)
    if err != nil { return err }
    return nil
}

func GetArticlesFromUser(userId int) ([]models.Article, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := GetUser(userId)
    var artArr []models.Article
    if err != nil { return artArr, err }
    cmd := "SELECT id, title, description, creator_id, body, thumbnail_url, " +
           "date_created, is_authorized from ARTICLES WHERE creator_id=?"
    stmt, _ := db.Prepare(cmd)
    rows, err := stmt.Query(userId)
    if err != nil { return artArr, err }
    for rows.Next() {
        var a models.Article
        err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.CreatorID,
                         &a.Body, &a.ThumbnailUrl, &a.DateCreated,
                         &a.IsAuthorized)
        if err != nil { return artArr, err }
        artArr = append(artArr, a)
    }
    return artArr, nil
}

func GetApprovedArticles() ([]models.Article, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    var artArr []models.Article
    cmd := "SELECT id, title, description, creator_id, body, thumbnail_url, " +
           "date_created, is_authorized from ARTICLES WHERE is_authorized=TRUE"
    rows, err := db.Query(cmd)
    if err != nil { return artArr, err }
    for rows.Next() {
        var a models.Article
        err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.CreatorID,
                         &a.Body, &a.ThumbnailUrl, &a.DateCreated,
                         &a.IsAuthorized)
        if err != nil { return artArr, err }
        artArr = append(artArr, a)
    }
    return artArr, nil
}

func Get10MostRecentArticles() ([]models.Article, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    var artArr []models.Article
    cmd := "SELECT id, title, description, creator_id, body, thumbnail_url, " +
           "date_created, is_authorized from ARTICLES WHERE is_authorized= " +
           "TRUE LIMIT 10 ORDER BY date_created"
    rows, err := db.Query(cmd)
    if err != nil { return artArr, err }
    for rows.Next() {
        var a models.Article
        err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.CreatorID,
                         &a.Body, &a.ThumbnailUrl, &a.DateCreated,
                         &a.IsAuthorized)
        if err != nil { return artArr, err }
        artArr = append(artArr, a)
    }
    return artArr, nil
}

func GetArticleAuthor(articleId int) (models.User, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := GetArticle(articleId)
    var u models.User
    if err != nil { return u, err }
    cmd := "SELECT users.id, first_name, last_name, email, is_admin FROM " +
           "users LEFT JOIN articles ON users.id=articles.creator_id WHERE " +
           "articles.id=?"
    stmt, err := db.Prepare(cmd)
    err = stmt.QueryRow(articleId).Scan(&u.ID, &u.FirstName, &u.LastName,
                                         &u.Email, &u.IsAdmin)
    if err != nil { return u, err }
    return u, nil
}

func SaveArticle(userId, articleId int) error {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := GetArticle(articleId)
    if err != nil { return err }
    _, err = GetUser(userId)
    if err != nil { return err }
    cmd := "INSERT INTO saved_articles (article_id, user_id) VALUES (?, ?)"
    stmt, _ := db.Prepare(cmd)
    _, err = stmt.Exec(articleId, userId)
    if err != nil { return err }
    return nil
}

func GetSavedArticles(userId int) ([]models.Article, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := GetUser(userId)
    var articles []models.Article
    if err != nil { return articles, err }
    cmd := "SELECT articles.id, title, description, creator_id, body, " +
           "thumbnail_url, date_created, is_authorized from ARTICLES LEFT " +
           "JOIN saved_articles on articles.id=saved_articles.article_id " +
           "WHERE saved_articles.user_id=?"
    stmt, err := db.Prepare(cmd)
    if err != nil { return articles, err }
    rows, err := stmt.Query(userId)
    if err != nil { return articles, err }
    for rows.Next() {
        var a models.Article
        err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.CreatorID,
                         &a.Body, &a.ThumbnailUrl, &a.DateCreated,
                         &a.IsAuthorized)
        if err != nil { return articles, err }
        articles = append(articles, a)
    }
    return articles, nil
}

func CreateEvent(name, description, date, location string) error {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "INSERT INTO events (event_name, event_description, date, " +
           "location) VALUES (?, ?, STR_TO_DATE(?, '%Y-%m-%d'), ?)"
    stmt, _ := db.Prepare(cmd)
    _, err := stmt.Exec(name, description, date, location)
    if err != nil { return err }
    return nil
}

func GetEvent(id int) (models.Event, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, event_name, event_description, date, location FROM " +
           "events WHERE id=?"
    stmt, _ := db.Prepare(cmd)
    var e models.Event
    err := stmt.QueryRow(id).Scan(&e.ID, &e.EventName, &e.EventDescription,
                                  &e.Date, &e.Location)
    if err != nil { return e, err }
    return e, nil
}

func GetFutureEvents() ([]models.Event, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, event_name, event_description, date, location FROM " +
           "events WHERE date>=CURDATE()"
    stmt, _ := db.Prepare(cmd)
    var events []models.Event
    rows, err := stmt.Query()
    if err != nil { return events, err }
    for rows.Next() {
        var e models.Event
        err := rows.Scan(&e.ID, &e.EventName, &e.EventDescription, &e.Date,
                         &e.Location)
        if err != nil { return events, err }
        events = append(events, e)
    }
    return events, nil
}
