package models

type Message struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Text     string `db:"text"`
	Date     string `db:"date"`
	ChatID   string `db:"chat_id"`
}
