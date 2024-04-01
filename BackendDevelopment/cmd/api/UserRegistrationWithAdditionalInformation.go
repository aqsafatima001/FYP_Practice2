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
func (app *application) hashPassword_User_additional_Info(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (app *application) OTP_verfication_User_additional_Info(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Response Content-Type:", w.Header().Get("Content-Type"))

	// Parse the JSON data sent from the frontend
	var formData struct {
		Username         string `json:"username"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		DateOfBirth      string `json:"dateOfBirth"`
		Nationality      string `json:"nationality"`
		NICNumber        string `json:"nicNumber"`
		PassportNumber   string `json:"passportNumber"`
		Title            string `json:"title"`
		FirstName        string `json:"firstName"`
		MiddleName       string `json:"middleName"`
		LastName         string `json:"lastName"`
		UniversityCentre string `json:"universityCentre"`
		ThesisTitle      string `json:"thesisTitle"`
		StartDate        string `json:"startDate"`
		EndDate          string `json:"endDate"`
		PictureProof     string `json:"pictureProof"`
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

func (app *application) VerifyOTP_User_additional_Info(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data sent from the frontend
	var formData struct {
		Email            string `json:"email"`
		OTP              string `json:"otp"`
		Password         string `json:"password"`
		Username         string `json:"username"`
		DateOfBirth      string `json:"dateOfBirth"`
		Nationality      string `json:"nationality"`
		NICNumber        string `json:"nicNumber"`
		PassportNumber   string `json:"passportNumber"`
		Title            string `json:"title"`
		FirstName        string `json:"firstName"`
		MiddleName       string `json:"middleName"`
		LastName         string `json:"lastName"`
		UniversityCentre string `json:"universityCentre"`
		ThesisTitle      string `json:"thesisTitle"`
		StartDate        string `json:"startDate"`
		EndDate          string `json:"endDate"`
		PictureProof     string `json:"pictureProof"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&formData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check OTP and email in the database
	expirationTime, err := app.verifyOTP_User_additional_Info(formData.Email, formData.OTP)
	if err != nil {
		fmt.Println("Error verifying OTP:", err)
		http.Error(w, "Invalid OTP or OTP expired", http.StatusUnauthorized)
		return
	}
	fmt.Println(expirationTime)
	// OTP is valid, proceed to user registration
	if err := app.registerUser_User_additional_Info(
		formData.Username,
		formData.Email,
		formData.Password,
		formData.DateOfBirth,
		formData.Nationality,
		formData.NICNumber,
		formData.PassportNumber,
		formData.Title,
		formData.FirstName,
		formData.MiddleName,
		formData.LastName,
		formData.UniversityCentre,
		formData.ThesisTitle,
		formData.StartDate,
		formData.EndDate,
		formData.PictureProof,
	); err != nil {
		fmt.Println("Error registering the user:", err)
		http.Error(w, "Error registering the user", http.StatusInternalServerError)
		return
	}

	if errEmail := app.sendEmailAccountCreationRequest(formData.Email); errEmail != nil {
		fmt.Println("Error sending email:", errEmail)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully")

	// You can add further logic here like deleting the used OTP record or redirecting to the next page.
}

// Separate function for OTP verification
func (app *application) verifyOTP_User_additional_Info(email, otp string) (time.Time, error) {
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
func (app *application) registerUser_User_additional_Info(username, email, password, dateOfBirth, nationality, nicNumber, passportNumber, title, firstName, middleName, lastName, universityCentre, thesisTitle, startDate, endDate, pictureProof string) error {
	hashedPassword, errHash := app.hashPassword_User_additional_Info(password)
	if errHash != nil {
		return errHash
	}

	insertUserQuery := `
        INSERT INTO User_Registration_pending_2 
            (Username, Email, PasswordHash, DateOfBirth, Nationality, NIC_Number, Passport_Number, Title, FirstName, MiddleName, LastName, University_Centre, ThesisTitle, StartDate, EndDate)
            VALUES 
            (@Username, @Email, @PasswordHash, @DateOfBirth, @Nationality, @NIC_Number, @Passport_Number, @Title, @FirstName, @MiddleName, @LastName, @University_Centre, @ThesisTitle, @StartDate, @EndDate)
    `

	_, err := app.db.Exec(insertUserQuery,
		sql.Named("Username", username),
		sql.Named("Email", email),
		sql.Named("PasswordHash", hashedPassword),
		sql.Named("DateOfBirth", dateOfBirth),
		sql.Named("Nationality", nationality),
		sql.Named("NIC_Number", nicNumber),
		sql.Named("Passport_Number", passportNumber),
		sql.Named("Title", title),
		sql.Named("FirstName", firstName),
		sql.Named("MiddleName", middleName),
		sql.Named("LastName", lastName),
		sql.Named("University_Centre", universityCentre),
		sql.Named("ThesisTitle", thesisTitle),
		sql.Named("StartDate", startDate),
		sql.Named("EndDate", endDate),
		sql.Named("PictureProof", pictureProof),
	)
	if err != nil {
		return err
	}

	return nil
}

// // Separate function for user registration
// func (app *application) registerUser_User_additional_Info(username, email, password string) error {

// 	hashedPassword, err_hashp := app.hashPassword_User_additional_Info(password)

// 	if err_hashp != nil {
// 		return err_hashp
// 	}

// 	// hashedPassword := password

// 	insertUserQuery := `
// 	INSERT INTO User_Registration_pending_2
// 		(Username, Email, PasswordHash,DateOfBirth, Nationality,
// 		 NIC_Number, Passport_Number, Title, FirstName, MiddleName, LastName,
// 		 University_Centre, ThesisTitle, StartDate, EndDate, PictureProof)

// 		 VALUES
// 		(@Username, @Email, @PasswordHash,@DateOfBirth, @Nationality,
// 		 @NIC_Number, @Passport_Number, @Title, @FirstName, @MiddleName, @LastName,
// 		 @University_Centre, @ThesisTitle, @StartDate, @EndDate, @PictureProof);
// 	`

// 	_, err := app.db.Exec(insertUserQuery, sql.Named("Username", username), sql.Named("Email", email), sql.Named("PasswordHash", hashedPassword))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
