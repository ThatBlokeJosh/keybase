package server

import (
    "encoding/json"
    "net/http"
    "keybase/utils"	
)

func JSON(w http.ResponseWriter, v any) {
    w.Header().Add("Content-Type", "application/json; charset=utf-8")
    message, _ := json.Marshal(v);
    w.Write(message);
}

func Default(w http.ResponseWriter, r *http.Request) {
    JSON(w, "Welcome!")
}

func Create(w http.ResponseWriter, r *http.Request) {
    var key utils.Key; 
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&key);
    utils.Set(key)
    JSON(w, key);
}

func Delete(w http.ResponseWriter, r *http.Request) {
    var key utils.Key; 
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&key);
    JSON(w, key);
}

func Update(w http.ResponseWriter, r *http.Request) {
    var key utils.Key; 
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&key);
    JSON(w, key);
}

func Read(w http.ResponseWriter, r *http.Request) {
    var key utils.Key; 
    decoder := json.NewDecoder(r.Body);
    decoder.Decode(&key); 
    JSON(w, utils.Get(key));
}
