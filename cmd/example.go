package main

import (
	"bitbucket.org/mrd0ll4r/tbotapi"
	"bitbucket.org/mrd0ll4r/tbotapi/model"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	api, err := tbotapi.New("YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	// just to show its working
	fmt.Printf("User ID: %d\n", api.ID)
	fmt.Printf("Bot Name: %s\n", api.Name)
	fmt.Printf("Bot Username: %s\n", api.Username)

	closed := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			select {
			case <-closed:
				wg.Done()
				return
			case val := <-api.Updates:
				typ := val.Message.Type()
				if typ != model.TEXT {
					//ignore non-text messages for now
					continue
				}

				// -> simple echo bot
				msg, err := api.SendMessage(model.NewRecipientFromChat(val.Message.Chat), *val.Message.Text)

				//or
				//msg, err := api.SendMessage(model.NewChatRecipient(val.Message.Chat.Id), *val.Message.Text)

				// -> simple echo bot with disabled web page preview
				//msg, err := api.SendMessageExtended(model.NewOutgoingMessage(model.NewChatRecipient(val.Message.Chat.Id), val.Message.Text).SetDisableWebPagePreview(true))

				// or:
				//msg, err := api.SendMessageExtended(model.NewOutgoingMessage(model.NewRecipientFromChat(val.Message.Chat), val.Message.Text).SetDisableWebPagePreview(true))

				// -> simple echo bot via forwarding
				//msg, err = api.ForwardMessage(model.NewRecipientFromChat(val.Message.Chat), val.Message.Chat, val.Message.Id)

				// -> bot that always sends an image as response
				//	file, err := os.Open("F:/image.jpg")
				//	if err != nil {
				//		fmt.Printf("Err: %s\n", err)
				//		continue
				//	}
				//	msg, err := api.SendPhoto(model.NewOutgoingPhoto(model.NewChatRecipient(val.Message.Chat.Id)), file, "image.jpg")

				if err != nil {
					fmt.Printf("Err: %s\n", err)
					continue
				}
				fmt.Printf("MessageID: %d, Text: %s, IsGroupChat:%t\n", msg.Message.Id, *msg.Message.Text, msg.Message.Chat.IsGroupChat())
			case val := <-api.Errors:
				fmt.Printf("Err: %s\n", val)
			}
		}
	}()

	// let it run for five minutes
	time.Sleep(time.Duration(5) * time.Minute)

	fmt.Println("Closing...")

	api.Close()
	close(closed)
	wg.Wait()
}
