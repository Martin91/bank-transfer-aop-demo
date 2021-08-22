// Package logic maintains the main business logic
package logic

import "net/http"

func Transfer(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Transfer balance is done successfully!"))
}
