package main

import (
	"banker/handler"
	"banker/logic"
)

func main() {
	handler.Register("/transfer", logic.Transfer)
	handler.Serve()
}
