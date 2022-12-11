package message

type Message struct {
	Sender    User
	Recipient User
	Text      string
}

type User struct {
	DisplayName string
}
