// main.go
package main

import "github.com/DulatMedApp/UsmAstana/backend/cmd/routers"

func main() {
	r := routers.SetupRouter()

	// Запустите сервер
	r.Run(":3000")
}
