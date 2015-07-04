package tbotapi

import (
	"bitbucket.org/mrd0ll4r/tbotapi/model"
	"errors"
	"fmt"
	"github.com/syncthing/syncthing/internal/sync"
	"menteslibres.net/gosexy/rest"
	"net/url"
)

type TelegramBotAPI struct {
	baseUri  string
	Id       int
	Name     string
	Username string
	Updates  chan model.Update
	Errors   chan error
	closed   chan struct{}
	wg       sync.WaitGroup
}

const apiBaseUri string = "https://api.telegram.org/bot%s"

func New(apiKey string) (*TelegramBotAPI, error) {
	toReturn := TelegramBotAPI{
		baseUri: fmt.Sprintf(apiBaseUri, apiKey),
		Updates: make(chan model.Update),
		Errors:  make(chan error),
		closed:  make(chan struct{}),
		wg:      sync.NewWaitGroup(),
	}
	user, err := toReturn.GetMe()
	if err != nil {
		return nil, err
	}
	toReturn.Id = user.User.Id
	toReturn.Name = user.User.FirstName
	toReturn.Username = user.User.Username

	toReturn.wg.Add(1)
	go toReturn.updateLoop()

	return &toReturn, nil
}

func (api *TelegramBotAPI) Close() {
	select {
	case <-api.closed:
		return
	default:
	}
	close(api.closed)
	api.wg.Wait()
}

func (api *TelegramBotAPI) updateLoop() {
	updates, err := api.getUpdates()
	var offset int

	for {
		select {
		case <-api.closed:
			api.wg.Done()
			return
		default:
		}

		if err != nil {
			api.Errors <- err
		} else if !updates.Ok {
			api.Errors <- errors.New(fmt.Sprintf("TBotAPI: GetUpdates:%d - %s", updates.ErrorCode, updates.Description))
		} else {
			updates.Sort()
			offset = putUpdatesInChannel(api.Updates, updates.Update)
		}

		if offset == -1 {
			updates, err = api.getUpdates()
		} else {
			updates, err = api.getUpdatesByOffset(offset + 1)
		}
	}
}

func putUpdatesInChannel(channel chan model.Update, updates []model.Update) int {
	highestOffset := -1
	for _, update := range updates {
		highestOffset = update.Id
		channel <- update
	}

	return highestOffset
}

func (api *TelegramBotAPI) getUpdates() (*model.UpdateResponse, error) {
	resp := &model.UpdateResponse{}
	querystring := url.Values{}
	querystring.Set("timeout", fmt.Sprint(60))
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/GetUpdates"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) getUpdatesByOffset(offset int) (*model.UpdateResponse, error) {
	resp := &model.UpdateResponse{}
	querystring := url.Values{}
	querystring.Set("timeout", fmt.Sprint(60))
	querystring.Set("offset", fmt.Sprint(offset))
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/GetUpdates"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) GetMe() (*model.UserResponse, error) {
	resp := &model.UserResponse{}
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/GetMe"), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendMessage(chatId int, text string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values{}
	querystring.Set("chat_id", fmt.Sprint(chatId))
	querystring.Set("text", text)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendMessage"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendMessageExtended(querystring model.Querystring) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendMessage"), url.Values(querystring))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) ForwardMessage(chatId, fromChatId, messageId int) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values{}
	querystring.Set("chat_id", fmt.Sprint(chatId))
	querystring.Set("from_chat_id", fmt.Sprint(fromChatId))
	querystring.Set("message_id", fmt.Sprint(messageId))
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/ForwardMessage"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
