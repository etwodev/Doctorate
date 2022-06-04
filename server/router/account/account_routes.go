package account

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
	"github.com/rs/zerolog/log"
)

func EmailVerificationPostRoute(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("account")
	platform := r.FormValue("platform")
	
	if email == "" || platform == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "No account or platform attribute")
		return
	}

	otp, err := helpers.GenerateOTP(6)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}


	var old OTPVerify
	old.Email = email

	var new OTPVerify
	new.CurrentTime = time.Now().UnixMicro()
	new.OTP = otp
	new.Email = email
	
	if !helpers.ValidateEmail(email) {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid email")
		return
	}
	
	exists, err := helpers.Engine.Exist(&old)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	} else {
		if exists {
			_, err = helpers.Engine.Update(&new)
			if err != nil {
				helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
				return
			}
		} else {
			_, err = helpers.Engine.Insert(&new)
			if err != nil {
				helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
				return
			}
		}
	}

	payload := []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n%s\n\n%s", static.EmailAddress, email, static.VerifyEmailSubject, static.VerifyEmailMime, strings.Replace(static.VerifyEmail, static.VerifyEmailCode, otp, 1)))
	auth := smtp.PlainAuth("", static.EmailAddress, static.EmailPassword, static.EmailHost)

	err = smtp.SendMail(static.EmailIP, auth, static.EmailAddress, []string{email}, payload)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	helpers.RespondWithCode(w, http.StatusOK, "0")
}

func EmailSubmitPostRoute(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("account")
	code := r.FormValue("code")
	var response AuthSubmit

	verify := OTPVerify{Email: email, OTP: code}
	exists, err := helpers.Engine.Get(&verify)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	if exists {
		if (time.Now().UnixMicro() - verify.CurrentTime > 1800000000) {
			helpers.RespondWithError(w, http.StatusBadRequest, "Please verify your account again")
			return
		}
	} else {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid code or Email")
		return
	}


	var account GeneralAccount
	account.Email = email
	exists, err = helpers.Engine.Get(&account)
	if err != nil {
		log.Debug().Msg(err.Error())
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	if !exists {
		Snowflake, err := helpers.GenerateSnowflake(1)
		if err != nil {
			log.Debug().Msg(err.Error())
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}

		account.MasterUID = Snowflake.String()
		account.MasterToken = fmt.Sprintf("%s-%s", account.MasterUID, helpers.GenerateSecureToken(32))
		log.Debug().Msg(account.MasterToken)

		_, err = helpers.Engine.Insert(&account)
		if err != nil {
			log.Debug().Msg(err.Error())
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}
	}

	response.YostarToken = account.MasterToken 
	response.YostarUID = account.MasterUID
	response.Result = 0
	response.YostarAccount = account.Email

	helpers.RespondWithJSON(w, http.StatusOK, response, "application/json")
}