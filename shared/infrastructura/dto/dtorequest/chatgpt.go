package dtorequest

type ChatGptDtorequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
