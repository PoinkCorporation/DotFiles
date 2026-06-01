package chat

type ChatService interface {
}

type chat struct {
}

func New() *chat {
	return &chat{}
}
