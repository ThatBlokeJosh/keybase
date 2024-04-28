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

func Register(w http.ResponseWriter, r *http.Request) {
    user := utils.DecodeUser(r);
    utils.Register(user);
    JSON(w, user);
}

func Login(w http.ResponseWriter, r *http.Request) {
    user := utils.DecodeUser(r);
    JSON(w, user);
}

func Remove(w http.ResponseWriter, r *http.Request) {
    user := utils.DecodeUser(r);
    utils.Remove(user);
    JSON(w, user);
}

func Create(w http.ResponseWriter, r *http.Request) {
    key := utils.DecodeKey(r);
    utils.Set(key)
    JSON(w, key);
}

func Read(w http.ResponseWriter, r *http.Request) {
    key := utils.DecodeKey(r);
    JSON(w, utils.Get(key));
}

func Update(w http.ResponseWriter, r *http.Request) {
    key := utils.DecodeKey(r);
    utils.Set(key)
    JSON(w, key);
}

func Delete(w http.ResponseWriter, r *http.Request) {
    key := utils.DecodeKey(r);
    utils.Delete(key);
    JSON(w, key);
}
