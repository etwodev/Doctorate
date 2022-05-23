package account

import (
	// "fmt"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/Etwodev/Doctorate/server/helpers"
)

func EmailVerificationPostRoute(w http.ResponseWriter, r *http.Request) {
	account := 	r.FormValue("account")
	platform := r.FormValue("platform")

	host := os.Getenv("HOST_SMTP_SERVER")
	address := host + ":587"
	pass := os.Getenv("HOST_SMTP_PASS")
	from := os.Getenv("HOST_SMTP_EMAIL")

	if account == "" || platform == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "No account or platform attribute")
		return
	}

	opt, err := helpers.GenerateOTP(6)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	to := []string{account}
	subject := "Verify your email!"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	top, err := helpers.Serialization("./static/assets/account/emailhead.html")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	bot, err := helpers.Serialization("./static/assets/account/emailbody.html")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	body := top + fmt.Sprintf(bot, string(opt))
	payload := []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n%s\n\n%s", from, account, subject, mime, body))

	auth := smtp.PlainAuth("", from, pass, host)
	err = smtp.SendMail(address, auth, from, to, payload)
	
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	// STORE VERIFICATION CODE IN DB FOR 30 MINS

	helpers.RespondWithCode(w, http.StatusOK, "0")
}
