package alerts

type Type string

const (
	TypeImage  Type = "image"
	TypeGhost  Type = "ghost"
	TypeFriend Type = "friend"
)

type PayloadImage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	ImageUrl string `json:"imageUrl"`
	AudioUrl string `json:"audioUrl"`
}

type PayloadGhost struct {
	Username    string `json:"username"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type PayloadFriend struct {
	Username        string `json:"username"`
	BackgroundColor string `json:"backgroundColor"`
	FriendName      string `json:"friendName"`
	ImageUrl        string `json:"imageUrl"`
}
