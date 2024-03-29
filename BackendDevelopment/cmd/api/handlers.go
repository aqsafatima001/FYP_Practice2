package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type RegistrationData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Define a struct for the response
type RegistrationResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func (app *application) authenticateUser(email, enteredPassword string) error {
	var storedPasswordHash string

	// Log the username before querying the database
	app.infoLog.Printf("Authenticating user: %s", email)

	// Retrieve stored password hash from the database
	err := app.db.QueryRow("SELECT PasswordHash FROM User_Registration_test WHERE Email = @email", sql.Named("email", email)).Scan(&storedPasswordHash)
	if err != nil {
		return err // User not found or other database error
	}

	// Compare the entered password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(enteredPassword))
	return err
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
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "Invalid JSON supplied, or JSON missing entirely"

		_ = app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	// Log the received credentials
	app.infoLog.Printf("Received login request. Email: %s, Password: %s", creds.Email, creds.Password)

	// Authenticate with the database
	err = app.authenticateUser(creds.Email, creds.Password)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "Login failed"

		_ = app.writeJSON(w, http.StatusUnauthorized, payload)
		return
	}

	// Authentication successful
	// send back a response
	payload.Error = false
	payload.Message = "Signed in"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
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

func (app *application) serveUserRegistrationPage(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to serve the user registration page.
	// This could be an HTML page or a Vue.js frontend.
}

func (app *application) getPendingRequestsHandler(w http.ResponseWriter, r *http.Request) {
	// Query pending user requests from the database
	rows, err := app.db.Query("SELECT email, username FROM User_Registration_pending")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pendingRequests []map[string]interface{}

	for rows.Next() {
		var email, username string
		if err := rows.Scan(&email, &username); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pendingRequests = append(pendingRequests, map[string]interface{}{
			"email":    email,
			"username": username,
		})
	}

	// Convert pendingRequests to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pendingRequests); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) acceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acceptUserQuery := `
        INSERT INTO User_Registration_test (Username, Email, PasswordHash)
        SELECT Username, Email, PasswordHash FROM User_Registration_pending WHERE Email = @Email;

        DELETE FROM User_Registration_pending WHERE Email = @Email;
    `

	_, err := app.db.Exec(acceptUserQuery, sql.Named("Email", requestData.Email))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusOK)
}

func (app *application) declineRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	declineUserQuery := `
        DELETE FROM User_Registration_pending WHERE Email = @Email;
    `

	_, err := app.db.Exec(declineUserQuery, sql.Named("Email", requestData.Email))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusOK)
}
