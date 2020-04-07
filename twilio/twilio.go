package twilio

import (
	"github.com/NavenduDuari/gomessenger/utils"
	"github.com/sfreiberg/gotwilio"
)

func SendWhatsappMsg(recipient, message string) {
	accountSid := utils.TwilioAccountSid
	authToken := utils.TwilioAuthToken
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "whatsapp:+14155238886"
	to := "whatsapp:+91" + recipient //
	twilio.SendSMS(from, to, message, "", "")
}
