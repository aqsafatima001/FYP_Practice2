package main

import (
	"database/sql"
	"net/http"
)

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) authenticateUser(email string) (string, error) {
	var storedPassword string

	// Log the username before querying the database
	app.infoLog.Printf("Authenticating user: %s", email)

	// err := app.db.QueryRow("SELECT Password FROM UserLogin2 WHERE Username = @username", sql.Named("username", username)).Scan(&storedPassword)
	err := app.db.QueryRow("SELECT Password FROM UserLogin2 WHERE Email = @email", sql.Named("email", email)).Scan(&storedPassword)
	return storedPassword, err
}

func (app *application) authenticateAdmin(email string) (string, error) {
	var storedPassword string

	// Log the username before querying the database
	app.infoLog.Printf("Authenticating Admin: %s", email)

	// err := app.db.QueryRow("SELECT Password FROM UserLogin2 WHERE Username = @username", sql.Named("username", username)).Scan(&storedPassword)
	err := app.db.QueryRow("SELECT Password FROM AdminLogin WHERE Email = @email", sql.Named("email", email)).Scan(&storedPassword)
	return storedPassword, err
}

func (app *application) UserloginAPI(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		// UserName string `json:"username"`
		Email    string `json:"email`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"

		_ = app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	// Log the received credentials
	app.infoLog.Printf("Received login request. Username: %s, Password: %s", creds.Email, creds.Password)

	// Authenticate with the database
	storedPassword, err := app.authenticateUser(creds.Email)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "User not found"

		_ = app.writeJSON(w, http.StatusUnauthorized, payload)
		return
	}

	// Compare the provided password with the stored password
	if creds.Password == storedPassword {
		// Log a message indicating successful login
		app.infoLog.Printf("User %s logged in successfully", creds.Email)

		// send back a response
		payload.Error = false
		payload.Message = "Signed in"

		err := app.writeJSON(w, http.StatusOK, payload)
		if err != nil {
			app.errorLog.Println(err)
		}
	} else {
		payload.Error = true
		payload.Message = "Login failed"

		_ = app.writeJSON(w, http.StatusUnauthorized, payload)
	}
}

func (app *application) AdminloginAPI(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		// UserName string `json:"username"`
		Email    string `json:"email`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"

		_ = app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	// Log the received credentials
	app.infoLog.Printf("Received login request. Username: %s, Password: %s", creds.Email, creds.Password)

	// Authenticate with the database
	storedPassword, err := app.authenticateAdmin(creds.Email)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "User not found"

		_ = app.writeJSON(w, http.StatusUnauthorized, payload)
		return
	}

	// Compare the provided password with the stored password
	if creds.Password == storedPassword {
		// Log a message indicating successful login
		app.infoLog.Printf("User %s logged in successfully", creds.Email)

		// send back a response
		payload.Error = false
		payload.Message = "Signed in"

		err := app.writeJSON(w, http.StatusOK, payload)
		if err != nil {
			app.errorLog.Println(err)
		}
	} else {
		payload.Error = true
		payload.Message = "Login failed"

		_ = app.writeJSON(w, http.StatusUnauthorized, payload)
	}
}
func (app *application) serveLogin(w http.ResponseWriter, r *http.Request) {

}
