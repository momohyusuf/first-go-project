package main

import "net/http"

func handleErrorREquest(w http.ResponseWriter, r *http.Request) {

	responseWithError(w, 404, "Sorry route does not exist")

}
