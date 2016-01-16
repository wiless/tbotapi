package main

import (
	"bitbucket.org/mrd0ll4r/tbotapi"
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
			case update := <-api.Updates:
				if update.Error() != nil {
					fmt.Printf("Update error: %s\n", update.Error())
					continue
				}

				val := update.Update()
				typ := val.Message.Type()
				if typ != tbotapi.TextType {
					//ignore non-text messages for now
					continue
				}

				// -> simple echo bot
				msg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(val.Message.Chat), *val.Message.Text).Send()

				//or
				//msg, err := api.NewOutgoingMessage(tbotapi.NewChatRecipient(val.Message.Chat.ID), *val.Message.Text).Send()

				// -> simple echo bot with disabled web page preview
				//msg, err := api.NewOutgoingMessage(tbotapi.NewChatRecipient(val.Message.Chat.ID), val.Message.Text).SetDisableWebPagePreview(true).Send()

				// or:
				//msg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(val.Message.Chat), val.Message.Text).SetDisableWebPagePreview(true).Send()

				// -> simple echo bot via forwarding
				//msg, err = api.NewOutgoingForward(tbotapi.NewRecipientFromChat(val.Message.Chat), val.Message.Chat, val.Message.ID).Send()

				// -> bot that always sends an image as response
				//msg, err := api.NewOutgoingPhoto(tbotapi.NewChatRecipient(val.Message.Chat.ID), "/path/to/your/image.jpg").Send()

				if err != nil {
					fmt.Printf("Err: %s\n", err)
					continue
				}
				fmt.Printf("MessageID: %d, Text: %s, IsGroupChat:%t\n", msg.Message.ID, *msg.Message.Text, msg.Message.Chat.IsGroupChat())
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
