package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// FUNCTION FOR PASSWORD HASHING
func (app *application) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (app *application) OTP_verfication(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Response Content-Type:", w.Header().Get("Content-Type"))

	// Parse the JSON data sent from the frontend
	var formData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&formData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Generate OTP
	otp, errOTP := app.generateOTP()
	if errOTP != nil {
		fmt.Println("Error generating OTP:", errOTP)
		http.Error(w, "Error generating OTP", http.StatusInternalServerError)
		return
	}
	fmt.Println("Generated OTP:", otp)

	// Send email with OTP
	email := formData.Email // Update with recipient's email
	if errEmail := app.sendEmailWithOTP(email, otp); errEmail != nil {
		fmt.Println("Error sending email:", errEmail)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully")

	// Save OTP and Email to OTPTable
	query := `
    INSERT INTO OTPTable (Email, otp, ExpirationTime)
    VALUES (@Email, @otp , DATEADD(MINUTE, 5, GETDATE()))
`
	_, err := app.db.Exec(query, sql.Named("Email", formData.Email), sql.Named("otp", otp))
	if err != nil {
		fmt.Println("Error saving OTP and email:", err)
		http.Error(w, "Error saving OTP and email", http.StatusInternalServerError)
		return
	}

	// Respond with a success JSON message
	response := Response{Message: "Registration successful"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (app *application) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data sent from the frontend
	var formData struct {
		Email    string `json:"email"`
		OTP      string `json:"otp"`
		Password string `json:"password"`
		Username string `json:"username"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&formData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check OTP and email in the database
	expirationTime, err := app.verifyOTP(formData.Email, formData.OTP)
	if err != nil {
		fmt.Println("Error verifying OTP:", err)
		http.Error(w, "Invalid OTP or OTP expired", http.StatusUnauthorized)
		return
	}
	fmt.Println(expirationTime)
	// OTP is valid, proceed to user registration
	if err := app.registerUser(formData.Username, formData.Email, formData.Password); err != nil {
		fmt.Println("Error registering the user:", err)
		http.Error(w, "Error registering the user", http.StatusInternalServerError)
		return
	}

	// You can add further logic here like deleting the used OTP record or redirecting to the next page.
}

// Separate function for OTP verification
func (app *application) verifyOTP(email, otp string) (time.Time, error) {
	query := `
        SELECT ExpirationTime
        FROM OTPTable
        WHERE Email = @Email AND OTP = @otp AND ExpirationTime >= GETDATE()
    `

	var expirationTime time.Time
	err := app.db.QueryRow(query, sql.Named("Email", email), sql.Named("otp", otp)).Scan(&expirationTime)
	if err != nil {
		return time.Time{}, err
	}

	return expirationTime, nil
}

// Separate function for user registration
func (app *application) registerUser(username, email, password string) error {

	hashedPassword, err_hashp := app.hashPassword(password)

	if err_hashp != nil {
		return err_hashp
	}

	// hashedPassword := password

	insertUserQuery := `
		INSERT INTO User_Registration_pending (Username, Email, PasswordHash)
		VALUES (@Username, @Email, @PasswordHash)
	`

	_, err := app.db.Exec(insertUserQuery, sql.Named("Username", username), sql.Named("Email", email), sql.Named("PasswordHash", hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
