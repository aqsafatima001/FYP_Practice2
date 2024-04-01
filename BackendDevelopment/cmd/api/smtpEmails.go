package main

import (
	"fmt"
	"net/smtp"
)

func (app *application) sendEmailWithOTP(email, otp string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)
	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: OTP Verification\r\n" +
		"\r\n" +
		"Your OTP for registration is: " + otp + "\r\n")
	to := email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
	fmt.Println("Email sent successfully")
	return nil
}

func (app *application) sendEmailAccountCreationRequest(email string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)
	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Account Creation Request Received \r\n" +
		"\r\n" +
		"Dear User, \r\n\r\n" +
		"Thank you for submitting your request for account creation.\r\n\r\n" +
		"We have received your request and it is currently being processed by our administrative team. We will review the information provided and notify you once a decision has been made.\r\n\r\n" +
		"Thank you for your patience.\r\n\r\n" +
		"Best regards")
	to := email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{to},
		msg,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Email sent successfully")
	return nil
}

func (app *application) sendEmailAcceptRequest(email, username string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)
	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Account Creation Request Accepted \r\n" +
		"\r\n" +
		"Dear " + username + ",\r\n" +
		"We are pleased to inform you that your request for account creation has been accepted. Your account has been successfully created, and you are now a valued member of our community.\r\n\r\n" +
		"You can now access all the features and services available to registered users. Please find below your account details:\r\n\r\n" +
		"Username: " + username + "\r\n" +
		"Email Address: " + email + "\r\n\r\n" +
		"We encourage you to explore our platform and take advantage of the resources and opportunities it offers. Should you have any questions or require assistance, please do not hesitate to contact our support team at [Support Email Address].\r\n\r\n" +
		"Thank you for choosing to be a part of our community. We look forward to serving you and providing you with an exceptional experience.\r\n\r\n" +
		"Best regards")
	to := email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{to},
		msg,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Email sent successfully")
	return nil
}

func (app *application) sendEmailRejectRequest(email string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth(
		"",
		"aqsafatima0202@gmail.com",
		"bgyn xcsk yfeb ajkz",
		"smtp.gmail.com",
	)
	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Account Creation Request Rejected \r\n" +
		"\r\n" +
		"Dear User, \r\n\r\n" +
		"We regret to inform you that your request for account creation has been rejected due to incomplete information provided.\r\n\r\n" +
		"Best regards")
	to := email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aqsafatima0202@gmail.com",
		[]string{to},
		msg,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Email sent successfully")
	return nil
}
