package commons

// Request struct
type GeneralRequestPushNotifStruct struct {
	Title   string            `json:"title"`
	Message string            `json:"message"`
	Token   string            `json:"token"`
	Data    map[string]string `json:"data"`
}
