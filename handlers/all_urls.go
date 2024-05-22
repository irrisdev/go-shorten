package handlers

import (
	"fmt"
	"github.com/irrisdev/go-shorten/utils"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	urls := utils.QueryAll()
	fmt.Println(urls)

}
