package helpers

import "net/http"

func SetSessionCookie(w http.ResponseWriter, userID string) {
	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    userID,
		MaxAge:   3600 * 24,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, coo)
}
