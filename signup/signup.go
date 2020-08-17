package signup

import "encoding/json"


// Request represents a signup request
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	GToken   string `json:"g_token"`
}

// UnmarshalSignupRequest unmarshal json into SignupRequest
func UnmarshalSignupRequest(data []byte) (Request, error) {
	var r Request
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal marshal signup request to json
func (r *Request) Marshal() ([]byte, error) {
	return json.Marshal(r)
}