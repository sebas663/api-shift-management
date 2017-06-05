package handlers

import (
	"fmt"
	"net/http"
)

//Index root of the api.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "There is no resource here!!!\n")
}
