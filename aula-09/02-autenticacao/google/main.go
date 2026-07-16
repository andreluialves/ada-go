// Login com Google usando Goth.
//
// Documentação:
//   - Goth: https://github.com/markbates/goth
//   - OAuth 2.0 do Google: https://developers.google.com/identity/protocols/oauth2/web-server
//   - OpenID Connect: https://developers.google.com/identity/openid-connect/openid-connect
package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)


func main() {	
	godotenv.Load()

	provider := google.New(
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		os.Getenv("GOOGLE_CALLBACK_URL"),
		"openid",
		"email",
		"profile",
	)
	
	provider.SetAccessType("online")
	goth.UseProviders(provider)

	
	sessionKey := make([]byte, 32)
	if _, err := rand.Read(sessionKey); err != nil {
		log.Fatal(err)
	}

	gothic.Store = newSessionStore(sessionKey)

	log.Println("acesse http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", routes()))
}


func newSessionStore(sessionKey []byte) *sessions.CookieStore {
	store := sessions.NewCookieStore(sessionKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   10 * 60,
		HttpOnly: true,
		Secure:   false, 
		SameSite: http.SameSiteLaxMode,
	}
	return store
}


func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /auth/{provider}", gothic.BeginAuthHandler)
	mux.HandleFunc("GET /auth/{provider}/callback", authCallback)
	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}


func authCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "não foi possível concluir o login com Google", http.StatusUnauthorized)
		return
	}

	fmt.Println("Usuário retornado pelo Google")
	fmt.Println("Provider:", user.Provider)
	fmt.Println("ID:", user.UserID)
	fmt.Println("Nome:", user.Name)
	fmt.Println("E-mail:", user.Email)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Login concluído!\n\nProvider: %s\nID: %s\nNome: %s\nE-mail: %s\n",
		user.Provider, user.UserID, user.Name, user.Email)
}
