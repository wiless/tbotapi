package tbotapi

import (
	"bitbucket.org/mrd0ll4r/tbotapi/model"
	"errors"
	"fmt"
	"github.com/bndr/gopencils"
	"github.com/syncthing/syncthing/internal/sync"
)

type TelegramBotAPI struct {
	baseApi  *gopencils.Resource
	Id       int
	Name     string
	Username string
	Updates  chan model.Update
	Errors   chan error
	closed   chan struct{}
	wg       sync.WaitGroup
}

const baseUri string = "https://api.telegram.org/bot%s"

func New(apiKey string) (*TelegramBotAPI, error) {
	toReturn := TelegramBotAPI{
		baseApi: gopencils.Api(fmt.Sprintf(baseUri, apiKey)),
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
	querystring := map[string]string{"timeout": fmt.Sprint(60)}
	_, err := api.baseApi.Res("GetUpdates", resp).Get(querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) getUpdatesByOffset(offset int) (*model.UpdateResponse, error) {
	resp := &model.UpdateResponse{}
	querystring := map[string]string{"offset": fmt.Sprint(offset), "timeout": fmt.Sprint(60)}
	_, err := api.baseApi.Res("GetUpdates", resp).Get(querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) GetMe() (*model.UserResponse, error) {
	resp := &model.UserResponse{}
	_, err := api.baseApi.Res("GetMe", resp).Get()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
