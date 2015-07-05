package main

import (
	"bitbucket.org/mrd0ll4r/tbotapi"
	"fmt"
	"github.com/syncthing/syncthing/internal/sync"
	"log"
	"time"
)

func main() {
	api, err := tbotapi.New("YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User ID: %d\n", api.Id)
	fmt.Printf("Bot Name: %s\n", api.Name)
	fmt.Printf("Bot Username: %s\n", api.Username)

	close := make(chan struct{})
	wg := sync.NewWaitGroup()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-close:
				wg.Done()
				return
			case val := <-api.Updates:
				// -> simple echo bot
				msg, err := api.SendMessage(val.Message.Chat.Id, val.Message.Text)

				// -> simple echo bot with disabled web page preview
				//msg, err := api.SendMessageExtended(model.NewOutgoingMessage(val.Message.Chat.Id, val.Message.Text).SetDisableWebPagePreview(true))

				// -> simple echo bot via forwarding
				//msg, err := api.ForwardMessage(val.Message.Chat.Id, val.Message.Chat.Id, val.Message.Id)

				// -> bot that always sends an image as response
				//	file, err := os.Open("F:/image.jpg")
				//	if err != nil {
				//		fmt.Printf("Err: %s\n", err)
				//		continue
				//	}
				//	msg, err := api.SendPhoto(model.NewOutgoingPhoto(val.Message.Chat.Id), file, "image.jpg")

				if err != nil {
					fmt.Printf("Err: %s\n", err)
					continue
				}
				fmt.Printf("MessageID: %d, Text: %s, IsGroupChat:%t\n", msg.Message.Id, msg.Message.Text, msg.Message.Chat.IsGroupChat)
			case val := <-api.Errors:
				fmt.Printf("Err: %s\n", val)
			}
		}
	}()

	timer := time.NewTimer(time.Duration(5) * time.Minute)

	<-timer.C

	fmt.Println("Closing...")

	api.Close()
}
