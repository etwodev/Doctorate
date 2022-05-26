package user

import (

	"net/http"
	"time"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/rs/zerolog/log"
)


func AgreementPostRoute(w http.ResponseWriter, r *http.Request) {
	bin, err := helpers.OpenFile("./static/config/Agreements.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	helpers.RespondWithRaw(w, http.StatusOK, bin, "application/json")
}

func LoginPostRoute(w http.ResponseWriter, r *http.Request) {

}

func CreatePostRoute(w http.ResponseWriter, r *http.Request) {

	newValue			:= 	0
	resultValue 		:= 	0
	deviceValue 		:= 	r.FormValue("deviceId")
	storeValue  		:= 	r.FormValue("channelId")

	if deviceValue == "" || storeValue == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "Internal error")
		return
	}
	
	accessUID, err		:=  helpers.GenerateSnowflake(1)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	snowflakeUID, err	:=  helpers.GenerateSnowflake(2)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	
	auth := new(AuthData)
	auth.DeviceID = deviceValue
	auth.StoreID = storeValue
	auth.TimeNow = time.Now().UnixMicro()
	auth.EntryToken = helpers.HashWithMD5(deviceValue)
	auth.AccessToken = helpers.GenerateSecureToken(16)
	auth.CryptoToken = helpers.GenerateSecureToken(12)
	auth.AccessUID = accessUID.String()
	auth.SnowflakeUID = snowflakeUID.String()

	_, err = helpers.Engine.Insert(auth)
	if err != nil {
		log.Debug().Err(err).Msg("An error has occured...")
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	resp := &AuthyJSON {
		Result: &resultValue,
		UID: accessUID.Int64(),
		Token: helpers.HashWithMD5(deviceValue),
		IsNew: &newValue,
	}


	log.Debug().Msgf("Data: ", resp)

	helpers.RespondWithJSON(w, http.StatusOK, resp, "application/octet-stream")
}