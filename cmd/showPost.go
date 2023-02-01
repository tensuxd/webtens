package cmd

import (
	"fmt"
	"net/http"
)

func ShowPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "dads")
}
