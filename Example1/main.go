package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Account struct {
	UserName string
	Email    string
	Tel      int
}

type AccountNofitier interface {
	NotifyAccountCreated(context.Context, Account) error
}

// Class
type SmsAccountNotifier struct {
}

// Metotd
func (n SmsAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("New account Created send by sms:", "username", account.UserName, "email", account.Email, "tel", account.Tel)
	return nil
}

// Class
type EmailAccountNotifier struct {
}

// Metotd
func (n EmailAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("New account Created send by email:", "username", account.UserName, "email", account.Email, "tel", account.Tel)
	return nil
}

// Class
type AccountHandler struct {
	// Interfaz injection
	AccountNofitier AccountNofitier
}

// Method
func (h AccountHandler) handlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		fmt.Print("Failed to decode the response body")
		return
	}
	// Interfaz injection
	if err := h.AccountNofitier.NotifyAccountCreated(r.Context(), account); err != nil {
		fmt.Print("failed to notify account created")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)

}

// the problem here is that we want to send a notification in diferents ways and we dont want to hard code here
// func notifyAccountCreated(account Account) error {
// 	time.Sleep(time.Millisecond * 500)

// 	slog.Info("New account Created: ", "username: ", account.UserName, "email: ", account.Email)
// 	return nil
// }

func main() {
	mux := http.NewServeMux()

	accountHandler := &AccountHandler{
		// Denpendency injection
		AccountNofitier: SmsAccountNotifier{},
		// AccountNofitier: EmailAccountNotifier{},
	}
	// mux.HandleFunc("/account", handlerCreateAccount)
	mux.HandleFunc("/account", accountHandler.handlerCreateAccount)
	http.ListenAndServe(":3000", mux)
}
