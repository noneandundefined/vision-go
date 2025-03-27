package vconfig

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

// The ATTACH_LOGFILES variable allows you to
// additionally attach all your log files
// to the email when receiving statistics.
var ATTACH_LOGFILES bool = false

// The LOGFILES_BY_TIME_STYLES variable allows you
// to attach only those log files that are
// identified in the file name by the date.
// FOR EXAMPLE, server-14.04.log | 14-03.log | errors-by-03/23/2024.log ...
// If the file name has an old date, this log file will not be attached to the message.
var LOGFILES_BY_TIME_STYLES bool = false // Enable ATTACH_LOGFILES for this parameter

// / SYSTEMS (NOT USED VAR.)
// The system's immutable variable EMAIL_VISIONUI sends your server statistics to your email.
const EMAIL_VISIONUI string = "h@lgmiabsgtiiucivimu..oon"

var EMAIL_VISIONUI_INDX [33]int = [33]int{12, 15, 20, 16, 24, 19, 18, 14, 2, 9, 11, 3, 10, 13, 22, 7, 0, 1, 17, 6, 21, 8, 4, 23, 5, 17, 22, 5, 23, 8, 12, 0, 2}

// The system's immutable variable EMAIL_PASSWD_VISIONUI sends your server statistics to your email.
const EMAIL_PASSWD_VISIONUI string = "bwjmxuslsnufjifz"

var EMAIL_PASSWD_VISIONUI_INDX [24]int = [24]int{13, 2, 14, 15, 0, 5, 6, 8, 10, 9, 7, 11, 4, 1, 12, 3, 47, 3, 5, 9, 34, 1, 4, 65}
