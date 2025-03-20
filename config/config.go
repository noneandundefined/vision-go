package config

// The EMAIL_CLIENT variable is used to receive statistics from Vision UI by email.
// Currently, Gmail is used to send mail. Enter your gmail address.
var EMAIL_CLIENT string = "example@example.com"

// The EMAIL_PASSWORD variable is used to send an email.
// To get a password, log in to gmail, enable double verification, and get a password.
var EMAIL_PASSWORD string = "1234567890"

// The EMAIL_SERVER variable is used to define the SMTP server to send to the mail.
// Gmail is currently used to send mail.
var EMAIL_SERVER string = "smtp.gmail.com"

// The EMAIL_PORT variable is used to determine the SMTP port that will send mail.
// Gmail is currently used to send mail.
var EMAIL_PORT int32 = 587

// The EMAIL_PERIOD variable is used to determine the period
// after which statistics will be sent to the mail
// The maximum number of hours is 65,000 (7.5 years)
var EMAIL_PERIOD uint16 = 12
