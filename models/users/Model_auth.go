package users

import (
	"../../models"
	"crypto/md5"
	"encoding/hex"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

type Models_init_Users models.DB_init

func (model Models_init_Users) ForgotPassword(email string, pin string) (responseEmail int, responsePin int, err error) {

	var checkEmailInt int
	var checkPinInt int
	checkEmail := model.DB.QueryRow(`select count(*) from user where email = ?`, email).Scan(&checkEmailInt)

	if checkEmail != nil {
		log.Println(checkEmail.Error())
	}
	checkPin := model.DB.QueryRow(`select count(*) from user where recovery_pin = ?`, pin).Scan(&checkPinInt)

	if checkPin != nil {
		log.Println(checkPin.Error())
	}
	if checkEmailInt > 0 {
		responseEmail = 1
	} else {
		responseEmail = 0
	}

	if checkPinInt > 0 {
		responsePin = 1
	} else {
		responsePin = 0
	}
	//send email
	if checkEmailInt > 0 && checkPinInt > 0 {

		hasher := md5.New()
		hasher.Write([]byte(email))
		tokenString := hex.EncodeToString(hasher.Sum(nil))

		linkForgotPassword := `link/` + email + `/` + tokenString
		from := mail.NewEmail("Kasumi", "siapasayachannel@gmail.com")
		subject := "Your Kasumi password reset request"
		to := mail.NewEmail("Example User", email)
		plainTextContent := `hello` + email
		htmlContent := `A request has been received to change the password for your Kasumi account.
					<br><br><a href="` + linkForgotPassword + `">Reset Password</a> </strong>`
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		_, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			//fmt.Println(response.StatusCode)
			//fmt.Println(response.Body)
			//fmt.Println(response.Headers)
			return responseEmail, responsePin, nil
		}
	}
	return responseEmail, responsePin, nil
}

func (model Models_init_Users) ForgotPasswordAction(email string, token string, NewPassword string) (responseForgot string, err error) {

	hasher := md5.New()
	hasher.Write([]byte(email))
	tokenString := hex.EncodeToString(hasher.Sum(nil))

	if tokenString == token {

		hasher := md5.New()
		hasher.Write([]byte(NewPassword))
		NewPasswordString := hex.EncodeToString(hasher.Sum(nil))

		ForgotAction, errForgotAction := model.DB.Exec(`update user set password = ? where email =?`, NewPasswordString, email)

		if errForgotAction != nil {
			log.Println(errForgotAction)
		}

		rowsAffected, _ := ForgotAction.RowsAffected()

		if rowsAffected == 1 {
			return "Success Response", nil
		}
	}

	return "Bad Request", nil
}
