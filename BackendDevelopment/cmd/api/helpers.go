package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
)

// -------------- READ JSON FUNCTION ------------
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	maxBytes := 1048576 //1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

// -------------- WRITE JSON FUNCTION ------------
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {

	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// -------------- ERROR FUNCTION ------------
func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	app.writeJSON(w, statusCode, payload)
}

// -------------- ERROR FUNCTION ------------
func (app *application) sendMailSimple() {
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)

	msg := "Subject : My special subject\nThis is the body of my email."

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{"aqsafatima0202@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func (app *application) generateOTP() (string, error) {
	// Generate random bytes (should be enough for a 6-digit OTP)
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// Encode the random bytes to a 6-digit numeric OTP
	otp := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	return otp[:6], nil
}

func (app *application) sendEmailWithOTP(email, otp string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)

	// Compose the email message
	// to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: OTP Verification\r\n" +
		"\r\n" +
		"Your OTP for registration is: " + otp + "\r\n")

	// Send the email
	// err := smtp.SendMail(
	// 	"smtp.example.com:587",
	// 	auth,
	// 	"aqsafatima0202@gmail.com",
	// 	to,
	// 	msg,
	// )
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{"aqsafatima0202@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
	fmt.Println("Email sent successfully")
	return nil
}
