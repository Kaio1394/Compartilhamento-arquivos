package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "kaio" {
		return "$1$FnAB9N4Z$B2y4R8P1EyrOyQBE5rFpV0"
	}
	return ""
}
func main() {
	// validação de parâmetros
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <dir> <port>")
		os.Exit(1)
	}

	// Parâmetros via terminal
	httpDIr := os.Args[1]
	port := os.Args[2]

	// auth
	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDIr)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("Subindo servidor na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
