package database

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/variables"
	"time"
)

func InsertUser(user *variables.User) {

	InsertData :=
		`
	INSERT INTO users
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := DB.Exec(InsertData, user.ID, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByEmail(email string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE email = ?;
	`
	Rows, err := DB.Query(GetData, email)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func GetUserByNickname(nickname string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE nickname = ?;
	`
	Rows, err := DB.Query(GetData, nickname)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func InsertPost(post *variables.Post) {

	InsertData :=
		`
	INSERT INTO posts (
		created_at, 
		title, 
		content, 
		category, 
		user_id
		)
	VALUES (?, ?, ?, ?, ?);
	`

	_, err := DB.Exec(InsertData, time.Now(), post.Title, post.Content, post.Category, post.User.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPostByID(id int) *variables.Post {
	var post variables.Post

	GetData :=
		`
	SELECT * FROM posts
	WHERE id = ?;
	`
	Rows, err := DB.Query(GetData, id)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		var userID string
		var date time.Time
		err = Rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &userID, &date)
		if err != nil {
			log.Fatal(err)
		}
		post.User = GetUserByID(userID)
		post.Date = date.Format("Mon 2 Jan 15:04")
	}
	return &post
}

func GetCurrentUser(r *http.Request) *variables.User {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error getting cookie:", err)
		return nil
	}

	var userID string

	err = DB.QueryRow("SELECT user_id FROM sessions WHERE id = ? AND expiration > ?", cookie.Value, time.Now()).Scan(&userID)
	if err != nil {
		fmt.Println("Error getting user ID from session:", err)
		return nil
	}

	user := GetUserByID(userID)
	return user

}

func GetAllUsers(r *http.Request) []map[string]any {
	user := GetCurrentUser(r)
	if user == nil {
		return nil
	}
	rows, err := DB.Query("SELECT * FROM users WHERE id != ?", user.ID)
	if err != nil {
		log.Println("Error getting users:", err)
		return nil
	}
	defer rows.Close()

	var users []map[string]any

	for rows.Next() {
		var user variables.User
		err := rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Println("Erreur lors du scan d’un utilisateur :", err)
			continue
		}
		users = append(users, map[string]any{
			"user":      user,
			"connected": checkIfUserConnected(user.ID)})
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over users:", err)
		return nil
	}
	return users
}

func checkIfUserConnected(user_id string) bool {
	var count int
	GetData :=
		`
	SELECT COUNT(*) FROM sessions
	WHERE user_id = ?
	AND expiration > ?;
	`
	err := DB.QueryRow(GetData, user_id, time.Now()).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return true
	}
	return false

}

func GetUserByID(id string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE id = ?;
	`
	Rows, err := DB.Query(GetData, id)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		password := []byte{}
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func GetpostHome() []*variables.Post {
	var posts []*variables.Post

	GetData :=
		`
	SELECT * FROM posts
	ORDER BY created_at DESC;
	`
	Rows, err := DB.Query(GetData)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		var post variables.Post
		var userID string
		var date time.Time
		err = Rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &userID, &date)
		if err != nil {
			log.Fatal(err)
		}
		post.User = GetUserByID(userID)
		post.Date = date.Format("Mon 2 Jan 15:04")
		posts = append(posts, &post)
	}
	return posts
}

func InsertSession(session_token string, user *variables.User) {
	DeleteSessionFromUserID(user.ID)
	InsertData :=
		`
	INSERT INTO sessions
	VALUES (?, ?, ?);
	`

	expiration := time.Now().Add(time.Hour * 1)
	_, err := DB.Exec(InsertData, session_token, user.ID, expiration)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session inserted")

}

func DeleteSessionFromUserID(userID string) {
	DeleteData :=
		`
	DELETE FROM sessions
	WHERE user_id = ?;
	`

	_, err := DB.Exec(DeleteData, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session deleted")

}

func DeleteSession(session_token string) {
	DeleteData :=
		`
	DELETE FROM sessions
	WHERE id = ?;
	`

	_, err := DB.Exec(DeleteData, session_token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session deleted")

}

func GetNicknameByUserId(userID string) string {
	var nickname string

	GetData :=
		`
	SELECT nickname FROM users
	WHERE id = ?;
	`
	err := DB.QueryRow(GetData, userID).Scan(&nickname)
	if err != nil {
	}
	return nickname

}

func GetUserIdBySession(id string) string {
	var userID string

	GetData :=
		`
	SELECT user_id FROM sessions
	WHERE id = ?;
	`
	err := DB.QueryRow(GetData, id).Scan(&userID)
	if err != nil {
	}
	return userID
}

func InsertComment(comment *variables.Comment) {
	query := `
	INSERT INTO comments (content, post_id, user_id,created_at)
	VALUES (?, ?, ?,?);`
	_, err := DB.Exec(query, comment.Content, comment.PostID, comment.UserID, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

// Récupérer les commentaires d’un post
func GetCommentsByPostID(postID int) []variables.Comment {
	query := `
	SELECT id, content, post_id, user_id, created_at
	FROM comments
	WHERE post_id = ?
	ORDER BY created_at ASC;
	`

	rows, err := DB.Query(query, postID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var comments []variables.Comment

	for rows.Next() {
		var c variables.Comment
		err := rows.Scan(&c.ID, &c.Content, &c.PostID, &c.UserID, &c.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, c)
	}
	return comments
}
