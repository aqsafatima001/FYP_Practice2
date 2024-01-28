package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) OTP_verfication(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Response Content-Type:", w.Header().Get("Content-Type"))

	// Generate OTP
	otp, errOTP := app.generateOTP()
	if errOTP != nil {
		fmt.Println("Error generating OTP:", errOTP)
		http.Error(w, "Error generating OTP", http.StatusInternalServerError)
		return
	}
	fmt.Println("Generated OTP:", otp)

	// Send email with OTP
	email := "aqsafatima0202@gmail.com" // Update with recipient's email
	if errEmail := app.sendEmailWithOTP(email, otp); errEmail != nil {
		fmt.Println("Error sending email:", errEmail)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully")

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
