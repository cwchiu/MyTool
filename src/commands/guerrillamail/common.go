package guerrillamail

import (
	"fmt"
	"github.com/qianlnk/guerrillamail"
)

func getEmailList(token string) {
	mailClient := guerrillamail.NewGuerrillamailClient(nil)
	mailClient.SidToken = token //
	resp, err := mailClient.GetEmailList(guerrillamail.Argument{
		"offset": "0",
		"seq":    "0",
	})

	if err != nil {
		panic(err)
	}
	// fmt.Println(resp.Count)
	for _, mail := range resp.List {
		fmt.Printf("%v|%s|%s|%v\n", mail.MailID, mail.MailSubject, mail.MailFrom, mail.MailRead)
	}

}

func checkMail(token string) {
	mailClient := guerrillamail.NewGuerrillamailClient(nil)
	mailClient.SidToken = token //
	resp, err := mailClient.CheckEmail(guerrillamail.Argument{
		"seq": "0",
	})

	if err != nil {
		panic(err)
	}
	// fmt.Println(resp.Count)
	for _, mail := range resp.List {
		fmt.Printf("%v|%s|%s|%v\n", mail.MailID, mail.MailSubject, mail.MailFrom, mail.MailRead)
	}
}

func fetchMail(token string, email_id string) {
	mailClient := guerrillamail.NewGuerrillamailClient(nil)
	mailClient.SidToken = token //
	resp, err := mailClient.FetchEmail(guerrillamail.Argument{
		"email_id": email_id,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.MailBody)
}

func delMail(token string, email_id string) {
	mailClient := guerrillamail.NewGuerrillamailClient(nil)
	mailClient.SidToken = token //
	resp, err := mailClient.DelEmail(guerrillamail.Argument{
		"email_ids[]": email_id,
	})

	if err != nil {
		panic(err)
	}
	if len(resp.DeletedIDs) > 0 {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

func getEmail() {
	mailClient := guerrillamail.NewGuerrillamailClient(nil)
	resp, err := mailClient.GetEmailAddress(guerrillamail.Argument{})
	if err != nil {
		panic(err)
	}
	// fmt.Println(resp)
	fmt.Printf("Email: %v\n", resp.EmailAddr)
	fmt.Printf("Timestamp: %v\n", resp.EmailTimestamp)
	fmt.Printf("Alias: %v\n", resp.Alias)
	fmt.Printf("Token: %v\n", resp.SidToken) // session token
}
