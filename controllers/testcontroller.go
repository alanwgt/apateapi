package controllers

import (
	"net/http"
	"strconv"

	"github.com/alanwgt/apateapi/cache"
	"github.com/alanwgt/apateapi/messages"
	"github.com/alanwgt/apateapi/services"
)

func TestMe(w http.ResponseWriter, r *http.Request) {
	uc, _ := cache.GetUser("alanwgt")

	md := map[string]string{
		"sender":       uc.Model.Username,
		"id":           strconv.FormatInt(23, 10),
		"message_type": "text",
		"alguma coisa": "nada",
	}

	// log.Println("Message sent to FCMID:", uc.Model.FcmToken)
	services.SendFCMMessage(
		"fX4A5m_1PY4:APA91bGJaCRI1aD7sMD_FvBDJBNO_TqDtZ2fxj9k4wh0oNDof5Wnw0XK6Fi_ngLxElHyFIQf3S0Eh2g0I3QQgcrgMVIxVlQSPoSbjtAs9kBDBwo5SY2Ra2asvS1-MxeO8JrihEAivuOz",
		"Newmessage",
		"fromranco",
		"message_test",
		"Text",
		"eastward",
		"12")
	messages.RawJSON(w, md)
}
