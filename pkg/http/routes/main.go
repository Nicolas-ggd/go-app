package routes

import "net/http"

func ServeApp() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    "8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
