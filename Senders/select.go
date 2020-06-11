package Senders

import "fmt"

func Selector(p Pakupkas) {
	p1 := Pakupka{}
	c := Client{}
	for _, val := range p {
		if err := p1.ScanPakupkaBY_Id_client(val.Id_client); err != nil {
			fmt.Println("ERR-ScanPakupkaBYP", err.Error)
		} else {
			if err := c.GetClients(p1.Id_client); err == nil {
				//оператор интихоб кунад  агар логикаш хай давом  дорад!
				// Email ба
				SendEmail(c.Email, c.Name, p1)
				//Sms номери телфонба
				SendSms(c.Tel, c.Name, p1)
			} else {
				fmt.Println("ERROR-GetClients")
			}

		}

	}

}
