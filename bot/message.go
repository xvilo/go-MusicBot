package bot

type Sender struct {
	Name     string
	NickName string
}

type Message struct {
	Message   string
	Sender    Sender
	Target    string
	IsPrivate bool
}
