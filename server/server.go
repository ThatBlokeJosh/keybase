package server

import (
    "net/http"
)

const PORT = ":3000";

func Serve() {
    app := http.NewServeMux();

    // Root
    app.HandleFunc("/", Default);

    // Keys
    app.HandleFunc("POST /keys", Create);
    app.HandleFunc("GET /keys", Read);
    app.HandleFunc("PUT /keys", Update);
    app.HandleFunc("DELETE /keys", Update);

    // Users
    app.HandleFunc("POST /user", Register);
    app.HandleFunc("PUT /user", Login);
    app.HandleFunc("DELETE /user", Remove);

    http.ListenAndServe(PORT, app);
}
