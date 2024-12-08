package workers

import (
	db "ametory-crud/database"
	"ametory-crud/objects"
	srv "ametory-crud/services"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func SendRegMail() {
	var emailData objects.UserReg
	// fmt.Println("SendRegMail")
	data, err := db.REDIS.LPop(objects.QueueSendMail).Result()
	if err != nil {
		// log.Println("ERROR SEND EMAIL TO ", emailData.Email, err)
		return
	}
	json.Unmarshal([]byte(data), &emailData)

	log.Println("READY SEND EMAIL TO ", emailData.Email)
	// go func() {
	srv.MAIL.SetAddress(emailData.Name, emailData.Email)
	subject := "Pendaftaran User"
	if emailData.Subject != "" {
		subject = emailData.Subject
	}
	srv.MAIL.SetTemplate("template/layout.html", "template/new_user.html")
	err = srv.MAIL.SendEmail(subject, gin.H{
		"Email":    emailData.Email,
		"Name":     emailData.Name,
		"Link":     emailData.Link,
		"Password": emailData.Password,
	}, []string{})
	if err != nil {
		log.Println("ERROR SEND EMAIL TO ", emailData.Email, err)
	} else {
		log.Println("EMAIL TO ", emailData.Email, "SENT")
	}

}
