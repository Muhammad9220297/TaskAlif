package Senders

import (
	"fmt"
	"net/smtp"
)

const (
	userName = "Muhammad"
	password = "1234"
	hostName = "Alifjob.com"
)

func SendEmail(Email string, Name string, p Pakupka) {
	auth := smtp.PlainAuth("", userName, password, hostName)
	to := []string{"alif@mail.com"}
	createmassage := fmt.Sprintf("Салом! %s  Шумо %s-ро  ба  Maблаги  %f (сомони)  бо alif харид  кардед.Бо Эхтиром Alif:)", Name, p.Tovar, p.Sena_pokupka)
	massage := []byte(createmassage)
	//fmt.Println("Send  test  email:sending  Email")
	err := smtp.SendMail(hostName, auth, "muhammad9220297@mail.ru", to, massage)
	if err != nil {
		fmt.Println("xatogi ! ", err)
	} else {

		fmt.Println("Send test email: success !")
	}

	//Давомаш хай  холи
}
