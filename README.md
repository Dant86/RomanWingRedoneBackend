# Roman wing back-end library

## Build instructions:

__MacOS__: `go build -o libbackend.dylib -buildmode=c-shared main.go`

__Linux__: `go build -o libbackend.so -buildmode=c-shared main.go`

__Windows__: `go build -o libbackend.dll -buildmode=c-shared main.go`

## Library Function Documentation

Each library function returns a JSON string that still has to be parsed
wherever you import it.

### `CreateUser(fName, lName, email, pword)`

__Parameters__:

_fName: string_

    User's first name

_lName: string_

    User's last name

_email: string_

    User's email

_pword: string_

    User's password

__Returns__:

On success: [ID](#id)

On failure: [Error](#error)

### `GetUser(id)`

__Parameters__:

_id: int_

    User's ID within its table

__Returns__:

On success: [User](#user)

On failure: [Error](#error)

### `GetUserByEmail(email)`

_email: string_

    User's email address

__Returns__:

On success: [User](#user)

On failure: [Error](#error)

### `ValidateUser(email, pword)`

__Parameters__:

_email: string_

    User's ID within its table

_pword: string_

    Provided password to be checked against hash

__Returns__:

On success: __nothing__

On failure: [Error](#error)

### `CreateArticle(creatorId, title, description, body, thumbnailUrl)`

__Parameters__:

_creatorId: int_

    ID of the person creating the article

_title: string_

    Title of the article

_description: string_

    Description of the article

_body: string_

    Article body

_thumbnailUrl: string_

    Link to thumbnail image of the article. Note that thumbnail is not nullable, so if no form data is passed client-side,
    choose a default thumbnail image and pass that through.

__Returns__:

On success: [ID](#id)

On failure: [Error](#error)

### `GetArticle(id)`

__Parameters__:

_id: int_

    The article's ID within its respective table

__Returns__:

On success: [Article](#article)

On failure: [Error](#error)

### `DeleteArticle(id)`

__Parameters__:

_id: int_

    The article's ID within its respective table

__Returns__:

On success: __nothing__

On failure: [Error](#error)

### `ApproveArticle(id)`

__Parameters__:

_id: int_

    The article's ID within its respective table

__Returns__:

On success: __nothing__

On failure: [Error](#error)

### `UpdateArticleBody(id, body)`

__Parameters__:

_id: int_

    The ID of the article whose body is to be updated

_body: string_

    The new body of the article

__Returns__:

On success: _none_

On failure: [Error](#error)

### `GetArticlesFromUser(userId)`

__Parameters__:

_userId: int_

    The ID of the user whose articles you want to view

__Returns__:

On success: Array of [Article](#article)s

On failure: [Error](#error)

### `GetApprovedArticles()`

__Parameters__:

_none_

__Returns__:

On success: Array of [Article](#article)s

on failure: [Error](#error)

### `Get10MostRecentArticles()`

__Parameters__:

_none_

__Returns__:

On success: Array of [Article](#article)s

on failure: [Error](#error)

### `GetArticleAuthor(articleId)`

__Parameters__:

_articleId: int_

    The ID of the article whose author you're looking for

__Returns__:

On success: [User](#user)

On failure: [Error](#error)

### `SaveArticle(userId, articleId)`

__Parameters__:

_userId: int_

    ID of the user who is saving the article

_articleId: int_

    ID of the article being saved

__Returns__:

On success: _none_

On failure: [Error](#error)

### `GetSavedArticles(userId)`

__Parameters__:

_userId: int_

    The ID of the user whose saved articles you're looking for

__Returns__:

On success: Array of [Article](#article)s

on failure: [Error](#error)

### `CreateEvent(name, description, date, location)`

__Parameters__:

_name: string_

    The name of the event

_description: string:_

    A quick description of the event

_date: string_

    The event's date. THIS MUST BE A STRING IN THE FORM: YYYY-MM-DD

_location: string_

    The event's location

__Returns__

On success: [ID](#id)

On failure: [Error](#error)

### `GetEvent(id)`

__Parameters__:

_id: int_

    The ID of the event you're looking for

__Returns__:

On success: [Event](#event)

On failure: [Error](#error)

### `GetFutureEvents()`

__Parameters__:

_none_

__Returns__:

On success: Array of [Event](#event)s

On failure: [Error](#error)

--------
## Definitions


### ID

<a name="id"></a>

|Name|Description|Type|
|---|---|---|
|ID|Newly created item's ID within its respective table|int|

### Error

<a name="error"></a>

|Name|Description|Type|
|---|---|---|
|Message|Golang's error message in the event of an error|string|

### User

<a name="user"></a>

|Name|Description|Type|
|---|---|---|
|ID|User's ID within its table|int|
|FirstName|User's first name|string|
|LastName|User's last name|string|
|Email|User's email address|string|
|IsAdmin|Whether a user is an admin|boolean|

### Article

<a name="article"></a>

|Name|Description|Type|
|---|---|---|
|ID|Article's ID within its table|int|
|Title|Title of the article|string|
|Description|Description of the article|string|
|CreatorID|ID of the article's creator|string|
|Body|The article's body|string|
|ThumbnailUrl|The URL for the article's thumbnail image|string|
|DateCreated|The date and time at which the article was created|datetime|
|IsAuthorized|Whether an admin has authorized the article|boolean|

### Event

<a name="event"></a>

|Name|Description|Type|
|---|---|---|
|ID|The event's ID within its table|int|
|EventName|The event's name|string|
|EventDescription|A quick description of the event|string|
|Date|The event's date|datetime|
|Location|The event's location|string|
