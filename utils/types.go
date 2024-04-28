package utils

import (
	"encoding/json"
	"net/http"
)

type Key struct {
    Username string `json:"username"`
    Name string `json:"name"`
    Key string `json:"key"`
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func DecodeUser(r *http.Request) (user User) {
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&user);
    return;
}

func DecodeKey(r *http.Request) (key Key) {
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&key);
    return;
}
