package rest

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type AddEntryRequest struct {
	URI      string `json:"url"`
	Category string `json:"category"`
}

type AddEntryResponse struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type FeedListResponse struct {
	Feeds []*Feed `json:"feeds"`
}
type Feed struct {
	Name string   `json:"name"`
	URIs []string `json:"uris"`
}
