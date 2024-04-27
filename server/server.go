package server

import (
    "net/http"
)

const PORT = ":3000";

func Serve() {
    app := http.NewServeMux();
    app.HandleFunc("/", Default);

    app.HandleFunc("POST /keys", Create);
    app.HandleFunc("GET /keys", Read);
    app.HandleFunc("PUT /keys", Update);
    app.HandleFunc("DELETE /keys", Update);

    http.ListenAndServe(PORT, app);
}
