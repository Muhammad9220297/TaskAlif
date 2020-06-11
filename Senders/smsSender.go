package Senders

import (
	"fmt"
	"net/http"

	"github.com/nexmo-community/nexmo-go"
)

func SendSms(tel string, Name string, p Pakupka) {
	Auth := nexmo.NewAuthSet()
	Auth.SetAPISecret("API_Key", "Api_Scret")
	client := nexmo.NewClient(http.DefaultClient, Auth)
	text := fmt.Sprintf("Салом! %s  Шумо %s-ро  ба  Maблаги  %f (сомони)  бо alif харид  кардед.Бо Эхтиром Alif:)", Name, p.Tovar, p.Sena_pokupka)
	smsContext := nexmo.SendSMSRequest{
		From: "Alif",
		To:   tel,
		Text: text,
	}
	smsResponse, _, err := client.SMS.SendSMS(smsContext)
	if err != nil {
		fmt.Println("Err_sms_Response:", err)
	}
	fmt.Println(smsResponse.Messages[0].Status)
	//Davom dorad..
}
