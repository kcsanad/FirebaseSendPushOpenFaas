package logic

import (
	"errors"
	"log"
	"net/http"

	"github.com/kcsanad/FirebaseSendPushOpenFaas/pkg/commons"
	"github.com/kcsanad/FirebaseSendPushOpenFaas/pkg/helpers"

	"github.com/gorilla/mux"
)

func SendPush(w http.ResponseWriter, r *http.Request) {

}

func SendPushWithCategory(w http.ResponseWriter, r *http.Request) {
	var p commons.GeneralRequestPushNotifStruct

	err := helpers.DecodeJSONBody(w, r, &p)
	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	vars := mux.Vars(r)
	category := vars["category"]

	app, err := initializeAppWithServiceAccount()
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
	}

	err = sendToTokenWithCategory(app, &category, &p)
	if err != nil {
		log.Printf("error sendToTokenWithCategory: %v\n", err)
	}
}
