package cmd

import (
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "dads")
}
