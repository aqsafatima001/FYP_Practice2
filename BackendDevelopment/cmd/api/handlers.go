package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	rows, err := app.db.Query("SELECT email, username FROM User_Registration_pending_2")
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

func (app *application) getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// Prepare the SQL statement with a placeholder for the email parameter
	stmt, err := app.db.Prepare("SELECT * FROM User_Registration_pending_2 WHERE Email = @Email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the prepared statement with the email parameter
	row := stmt.QueryRow(sql.Named("Email", email))

	// Define variables to store column values
	var (
		UserID            int
		Username          string
		Email             string
		PasswordHash      string
		NCP_ID            sql.NullInt64
		DateOfBirth       sql.NullString
		Nationality       sql.NullString
		NIC_Number        sql.NullString
		Passport_Number   sql.NullString
		Title             sql.NullString
		FirstName         sql.NullString
		MiddleName        sql.NullString
		LastName          sql.NullString
		University_Centre sql.NullString
		ThesisTitle       sql.NullString
		StartDate         sql.NullString
		EndDate           sql.NullString
		PictureProof      []byte // Assuming PictureProof is of type []byte
	)

	// Scan the row into variables
	err = row.Scan(&UserID, &Username, &Email, &PasswordHash, &NCP_ID, &DateOfBirth, &Nationality, &NIC_Number,
		&Passport_Number, &Title, &FirstName, &MiddleName, &LastName, &University_Centre, &ThesisTitle, &StartDate,
		&EndDate, &PictureProof)
	if err != nil {
		http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a map containing user information
	userInfo := map[string]interface{}{
		"UserID":            UserID,
		"Username":          Username,
		"Email":             Email,
		"PasswordHash":      PasswordHash,
		"NCP_ID":            NCP_ID.Int64,       // Use NCP_ID.Int64 to get the int64 value
		"DateOfBirth":       DateOfBirth.String, // Use DateOfBirth.String to get the string value
		"Nationality":       Nationality.String,
		"NIC_Number":        NIC_Number.String,
		"Passport_Number":   Passport_Number.String,
		"Title":             Title.String,
		"FirstName":         FirstName.String,
		"MiddleName":        MiddleName.String,
		"LastName":          LastName.String,
		"University_Centre": University_Centre.String,
		"ThesisTitle":       ThesisTitle.String,
		"StartDate":         StartDate.String,
		"EndDate":           EndDate.String,
		// Convert PictureProof to a base64 encoded string or any other appropriate representation
		// You may also want to handle nil values in nullable columns if necessary
	}

	// Convert userInfo to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userInfo); err != nil {
		http.Error(w, "Error encoding JSON: "+err.Error(), http.StatusInternalServerError)
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

	// ------------------------------------------------------------------------------------------
	// Retrieve user information from pending registration table
	getUserQuery := `
        SELECT Username FROM User_Registration_pending_2 WHERE Email = @Email;
    `
	row := app.db.QueryRow(getUserQuery, sql.Named("Email", requestData.Email))

	var userInfo struct {
		Username string
	}
	if err := row.Scan(&userInfo.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call function to create user on CentOS machine using the retrieved user information
	err_createuser := app.createUserOnCentOS(userInfo.Username, "aqsafatima")
	if err_createuser != nil {
		http.Error(w, err_createuser.Error(), http.StatusInternalServerError)
		return
	}

	// Send EMail to User that his request for Account Creation has been accepted
	email := requestData.Email // Update with recipient's email
	username := userInfo.Username
	if errEmail := app.sendEmailAcceptRequest(email, username); errEmail != nil {
		fmt.Println("Error sending email:", errEmail)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully")

	// ------------------------------------------------------------------------------------------

	acceptUserQuery := `
		INSERT INTO User_Registration_test_2 (Username, Email, PasswordHash, NCP_ID, DateOfBirth, Nationality, NIC_Number, Passport_Number, Title, FirstName, MiddleName, LastName, University_Centre, ThesisTitle, StartDate, EndDate, PictureProof)
        SELECT Username, Email, PasswordHash, NCP_ID, DateOfBirth, Nationality, NIC_Number, Passport_Number, Title, FirstName, MiddleName, LastName, University_Centre, ThesisTitle, StartDate, EndDate, PictureProof FROM User_Registration_pending_2 WHERE Email = @Email;

        DELETE FROM User_Registration_pending_2 WHERE Email = @Email;
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

	// Send EMail to User that his request for Account Creation has been REJECTED
	email := requestData.Email // Update with recipient's email
	if errEmail := app.sendEmailRejectRequest(email); errEmail != nil {
		fmt.Println("Error sending email:", errEmail)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully")

	declineUserQuery := `
        DELETE FROM User_Registration_pending_2 WHERE Email = @Email;
    `

	_, err := app.db.Exec(declineUserQuery, sql.Named("Email", requestData.Email))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusOK)
}
