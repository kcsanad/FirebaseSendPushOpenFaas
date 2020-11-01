package handlers

import (
	"log"
	"net/http"

	"github.com/kcsanad/FirebaseSendPushOpenFaas/pkg/logic"

	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	//var input []byte

	mr := mux.NewRouter()
	log.Println("Firebase PushNotification server api starting...")

	api := mr.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/send", logic.SendPush).Methods(http.MethodPost)
	api.HandleFunc("/send/category/{category}", logic.SendPushWithCategory).Methods(http.MethodPost)

	api.ServeHTTP(w, r)
}
