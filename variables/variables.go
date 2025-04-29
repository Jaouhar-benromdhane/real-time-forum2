package variables

import "time"

type User struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
}

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Date     string `json:"date"`
	User     *User  `json:"user"`
}

type UserStatus struct {
	Nickname      string `json:"nickname"`
	Online        bool   `json:"online"`
	HasNewMessage bool   `json:"hasNewMessage"`
}

type Message struct {
	Type     string   `json:"type"`
	Sender   string   `json:"sender"`
	Receiver string   `json:"receiver"`
	Content  string   `json:"content"`
	Userlist []string `json:"userlist"`
}

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    string    `json:"user_id"` // <- modifié ici pour que ça soit un string
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
