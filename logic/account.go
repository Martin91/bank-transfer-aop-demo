package logic

import "net/http"

func Transfer(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Transfer balance is done successfully!"))
}
