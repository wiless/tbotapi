package tbotapi

import (
	"bitbucket.org/mrd0ll4r/tbotapi/model"
	"errors"
	"fmt"
	"io"
	"menteslibres.net/gosexy/rest"
	"net/url"
	"sync"
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

func (api *TelegramBotAPI) SendMessageExtended(om *model.OutgoingMessage) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendMessage"), url.Values(om.GetQueryString()))
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

func (api *TelegramBotAPI) ResendPhoto(op *model.OutgoingPhoto, fileId string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values(op.GetQueryString())
	querystring.Set("photo", fileId)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendPhoto"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendPhoto(op *model.OutgoingPhoto, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"photo": []rest.File{
			{
				Name:   fileName,
				Reader: file,
			},
		},
	}

	message, err := rest.NewMultipartMessage(url.Values(op.GetQueryString()), files)
	if err != nil {
		return nil, err
	}

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseUri, "/SendPhoto"), message)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) ResendAudio(oa *model.OutgoingAudio, fileId string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values(oa.GetQueryString())
	querystring.Set("audio", fileId)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendAudio"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendAudio(oa *model.OutgoingAudio, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"audio": []rest.File{
			{
				Name:   fileName,
				Reader: file,
			},
		},
	}

	message, err := rest.NewMultipartMessage(url.Values(oa.GetQueryString()), files)
	if err != nil {
		return nil, err
	}

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseUri, "/SendAudio"), message)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) ResendDocument(od *model.OutgoingDocument, fileId string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values(od.GetQueryString())
	querystring.Set("document", fileId)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendDocument"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendDocument(od *model.OutgoingDocument, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"document": []rest.File{
			{
				Name:   fileName,
				Reader: file,
			},
		},
	}

	message, err := rest.NewMultipartMessage(url.Values(od.GetQueryString()), files)
	if err != nil {
		return nil, err
	}

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseUri, "/SendDocument"), message)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) ResendSticker(os *model.OutgoingSticker, fileId string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values(os.GetQueryString())
	querystring.Set("sticker", fileId)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendSticker"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendSticker(os *model.OutgoingSticker, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"sticker": []rest.File{
			{
				Name:   fileName,
				Reader: file,
			},
		},
	}

	message, err := rest.NewMultipartMessage(url.Values(os.GetQueryString()), files)
	if err != nil {
		return nil, err
	}

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseUri, "/SendSticker"), message)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) ResendVideo(ov *model.OutgoingVideo, fileId string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	querystring := url.Values(ov.GetQueryString())
	querystring.Set("video", fileId)
	err := rest.Get(resp, fmt.Sprint(api.baseUri, "/SendSticker"), querystring)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) SendVideo(ov *model.OutgoingVideo, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"video": []rest.File{
			{
				Name:   fileName,
				Reader: file,
			},
		},
	}

	message, err := rest.NewMultipartMessage(url.Values(ov.GetQueryString()), files)
	if err != nil {
		return nil, err
	}

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseUri, "/SendSticker"), message)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
