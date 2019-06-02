//this package contains structs for client and server communication
package rest

//LoginRequest is sent by the client for a login
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//LoginResponse is the servers response on a successful login
type LoginResponse struct {
	Token string `json:"token"`
}

//AddEntryRequest is sent by the client for adding an entry
type AddEntryRequest struct {
	URI      string `json:"url"`
	Category string `json:"category"`
}

//AddEntryResponse is the servers response when an entry is successfully added
type AddEntryResponse struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

//FeedListResponse is the servers answer when asked from the client which feeds exist
type FeedListResponse struct {
	Feeds []*Feed `json:"feeds"`
}

//Feed is a representation of a RSS Feed for the client
type Feed struct {
	Name string   `json:"name"`
	URIs []string `json:"uris"`
}
