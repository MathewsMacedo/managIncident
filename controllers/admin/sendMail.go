package admin

import (
	"fmt"
	"strconv"

	"gopkg.in/gomail.v1"
)

const (
	ip = "sokys.synology.me"
)

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func SendMailAdmin() error {
	smtpHost := "smtp.gmail.com"         // change to your SMTP provider address
	smtpPort := 587                      // change to your SMTP provider port number
	smtpPass := "72683564sm"             // change here
	smtpUser := "sokys.golang@gmail.com" // change here

	sender := "sokys.golang@gmail.com" // change here

	subject := "Demande de connexion\n"
	body := `<html><body><h1>Demande de Connexion</h1></body></html>`

	msg := gomail.NewMessage()
	msg.SetHeader("From", smtpUser)
	msg.SetHeader("To", sender)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	var err error
	go func() {
		mailer := gomail.NewMailer(smtpHost, smtpUser, smtpPass, smtpPort)
		err := mailer.Send(msg)
		if err != nil {
			panic(err)
			fmt.Println("Erreur : Email non envoyé")
		} else {
			fmt.Println("Email envoyé")
		}
	}()
	return err
}

func sendMail(mail string, md5 string) error {

	// authentication configuration
	smtpHost := "smtp.gmail.com"         // change to your SMTP provider address
	smtpPort := 587                      // change to your SMTP provider port number
	smtpPass := "72683564sm"             // change here
	smtpUser := "sokys.golang@gmail.com" // change here

	sender := "sokys.golang@gmail.com" // change here

	urlConfirm := "http://" + ip + ":8080/mail/confirmation/" + md5

	subject := "Confirmation d'inscription!\n"
	body := `<!DOCTYPE html>
            <html>
            <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <meta charset="utf-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <title>Confirmation de Mail</title>
            </head>
            <body style="background-color: #2B3E50; color: rgba(255,255,255,0.8); text-align: center; margin: 0 auto;" bgcolor="#2B3E50">
            <div style="width: auto; margin: 0 auto;">

            	<h1 style="position: static; width: 100%; background-color: #4E5D6C; height: auto; margin: 0 auto; padding: 10px;">
            <b>Incident Manager : </b>Demande de connexion</h1>
            	<br><p style="font-size: 20px; margin: 20px;">Suite à votre demande je vous fait part d'un lien.</p>
            	<p style="font-size: 20px; margin: 20px;">Celui vous permettra d'enregistrer votre mot de passe</p>
            	<p style="font-size: 20px; margin: 20px;">Voici le lien ci-dessous</p>
            	<p style="font-size: 20px; margin: 20px;"><a href="` + urlConfirm + `" style="color: rgba(255,255,255,1); text-decoration: underline;">Cliques ici pour valider ton inscription</a></p>
            <br><br><h3>A tout de suite. ;-)</h3>
            	<br><h1 style="position: static; width: 100%; background-color: #4E5D6C; height: auto; margin: 0 auto; padding: 10px;"><a href="mailto:` + sender + `" style="color: rgba(255,255,255,1); text-decoration: underline;">Si vous avez un soucis sur le site, faites nous un retour en cliquant sur ce lien</a>
            	</h1>
            </div>
            </body>
            </html>`

	msg := gomail.NewMessage()
	msg.SetHeader("From", smtpUser)
	msg.SetHeader("To", mail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	// Send the email to Bob, Cora and Dan
	var err error
	go func() {
		mailer := gomail.NewMailer(smtpHost, smtpUser, smtpPass, smtpPort)
		err := mailer.Send(msg)
		if err != nil {
			panic(err)
			fmt.Println("Erreur : Email non envoyé")
		} else {
			fmt.Println("Email envoyé")
		}
	}()
	return err
}

func SendMailUpdate(mail string, id int, titre string) error {
	smtpHost := "smtp.gmail.com"         // change to your SMTP provider address
	smtpPort := 587                      // change to your SMTP provider port number
	smtpPass := "72683564sm"             // change here
	smtpUser := "sokys.golang@gmail.com" // change here

	sender := "sokys.golang@gmail.com" // change here

	urlConfirm := "http://" + ip + ":8080//incident-manager/user/incident/" + strconv.Itoa(id)
	subject := "Mise à jour de votre incident!\n"
	body := `<!DOCTYPE html>
            <html>
            <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <meta charset="utf-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <title>Mise à jour de l'incident ` + titre + `</title>
            </head>
            <body style="background-color: #2B3E50; color: rgba(255,255,255,0.8); text-align: center; margin: 0 auto;" bgcolor="#2B3E50">
            <div style="width: auto; margin: 0 auto;">

            	<h1 style="position: static; width: 100%; background-color: #4E5D6C; height: auto; margin: 0 auto; padding: 10px;">
            <b>Incident Manager : </b>Ton incident ` + titre + ` a été mise à jour</h1>
            	<br>
            	<p style="font-size: 20px; margin: 20px;"><a href="` + urlConfirm + `" style="color: rgba(255,255,255,1); text-decoration: underline;">Retrouve ton incident en cliquant ici</a></p>
            <br><br><h3>A tout de suite. ;-)</h3>
            	<br><h1 style="position: static; width: 100%; background-color: #4E5D6C; height: auto; margin: 0 auto; padding: 10px;"><a href="mailto:` + sender + `" style="color: rgba(255,255,255,1); text-decoration: underline;">Si vous avez un soucis sur le site, faites nous un retour en cliquant sur ce lien</a>
            	</h1>
            </div>
            </body>
            </html>`

	msg := gomail.NewMessage()
	msg.SetHeader("From", smtpUser)
	msg.SetHeader("To", mail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	var err error
	go func() {
		mailer := gomail.NewMailer(smtpHost, smtpUser, smtpPass, smtpPort)
		err := mailer.Send(msg)
		if err != nil {
			panic(err)
			fmt.Println("Erreur : Email de modification non envoyé! :'(")
		} else {
			fmt.Println("Email de modification envoyé! :-)")
		}
	}()
	return err
}
