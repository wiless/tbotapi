package tbotapi

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// A TelegramBotAPI is an API Client for one Telegram bot.
// Create a new client by calling the New() function.
type TelegramBotAPI struct {
	ID       int                 // the bots ID
	Name     string              // the bots Name as seen by users
	Username string              // the bots username
	Updates  chan InternalUpdate // a channel providing updates this bot receives
	baseURIs map[method]string
	closed   chan struct{}
	c        *client
	wg       sync.WaitGroup
}

type InternalUpdate struct {
	update Update
	err    error
}

func (u *InternalUpdate) Update() Update {
	return u.update
}

func (u *InternalUpdate) Error() error {
	return u.err
}

const apiBaseURI string = "https://api.telegram.org/bot%s"

// New creates a new API Client for a Telegram bot using the apiKey provided.
// It will call the GetMe method to retrieve the bots id, name and username.
// Additionally, an update loop is started, pumping updates into the Updates channel.
func New(apiKey string) (*TelegramBotAPI, error) {
	toReturn := TelegramBotAPI{
		Updates:  make(chan InternalUpdate),
		baseURIs: createEndpoints(fmt.Sprintf(apiBaseURI, apiKey)),
		closed:   make(chan struct{}),
		c:        newClient(fmt.Sprintf(apiBaseURI, apiKey)),
	}
	user, err := toReturn.GetMe()
	if err != nil {
		return nil, err
	}
	toReturn.ID = user.User.ID
	toReturn.Name = user.User.FirstName
	toReturn.Username = *user.User.Username

	toReturn.wg.Add(1)
	go toReturn.updateLoop()

	return &toReturn, nil
}

// Close shuts down this client.
// Until Close returns, new updates and errors will be put into the respective channels.
// Note that, if no updates are received, this function may block for up to one minute, which is the time interval
// for long polling.
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
	offset := -1

	for {
		select {
		case <-api.closed:
			api.wg.Done()
			return
		default:
		}

		if err != nil {
			api.Updates <- InternalUpdate{err: err}
		} else {
			updates.Sort()
			for _, update := range updates.Update {
				api.Updates <- InternalUpdate{update: update}
				offset = update.ID
			}
		}

		if offset == -1 {
			updates, err = api.getUpdates()
		} else {
			updates, err = api.getUpdatesByOffset(offset + 1)
		}
	}
}

func (api *TelegramBotAPI) getUpdates() (*UpdateResponse, error) {
	resp := &UpdateResponse{}
	response, err := api.c.getQuerystring(getUpdates, resp, map[string]string{"timeout": fmt.Sprint(60)})

	if err != nil {
		if response != nil {
			if response.StatusCode() < 500 {
				return nil, err
			}
			//Telegram server problems, retry later...
			time.Sleep(time.Duration(5) * time.Second)
			return api.getUpdates()
		}
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *TelegramBotAPI) getUpdatesByOffset(offset int) (*UpdateResponse, error) {
	resp := &UpdateResponse{}
	response, err := api.c.getQuerystring(getUpdates, resp, map[string]string{
		"timeout": fmt.Sprint(60),
		"offset":  fmt.Sprint(offset),
	})

	if err != nil {
		if response != nil {
			if response.StatusCode() < 500 {
				return nil, err
			}
			//Telegram server problems, retry later...
			time.Sleep(time.Duration(5) * time.Second)
			return api.getUpdatesByOffset(offset)
		}
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMe returns basic information about the bot in form of a UserResponse.
func (api *TelegramBotAPI) GetMe() (*UserResponse, error) {
	resp := &UserResponse{}
	_, err := api.c.get(getMe, resp)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFile returns a FileResponse containing a Path string needed to download a file.
// You will have to construct the download link manually like
// https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response.
func (api *TelegramBotAPI) GetFile(fileID string) (*FileResponse, error) {
	resp := &FileResponse{}
	_, err := api.c.getQuerystring(getFile, resp, map[string]string{"file_id": fileID})

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func check(br *BaseResponse) error {
	if br.Ok {
		return nil
	}

	return fmt.Errorf("tbotapi: API error: %d - %s", br.ErrorCode, br.Description)
}

var ErrNoFileSpecified = errors.New("tbotapi: Neither a fileID nor a filePath were specified")

func (api *TelegramBotAPI) send(s sendable) (resp *MessageResponse, err error) {
	resp = &MessageResponse{}

	switch s := s.(type) {
	case *OutgoingMessage:
		_, err = api.c.postJSON(sendMessage, resp, s)
	case *OutgoingLocation:
		_, err = api.c.postJSON(sendLocation, resp, s)
	case *OutgoingForward:
		_, err = api.c.postJSON(forwardMessage, resp, s)
	case *OutgoingVideo:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendVideo, resp, file{fieldName: "video", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingVideo
				Video string `json:"video"`
			}{
				OutgoingVideo: *s,
				Video:         s.fileID,
			}
			_, err = api.c.postJSON(sendVideo, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	case *OutgoingPhoto:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendPhoto, resp, file{fieldName: "photo", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingPhoto
				Photo string `json:"photo"`
			}{
				OutgoingPhoto: *s,
				Photo:         s.fileID,
			}
			_, err = api.c.postJSON(sendPhoto, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	case *OutgoingVoice:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendVoice, resp, file{fieldName: "audio", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingVoice
				Audio string `json:"audio"`
			}{
				OutgoingVoice: *s,
				Audio:         s.fileID,
			}
			_, err = api.c.postJSON(sendVoice, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	case *OutgoingAudio:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendAudio, resp, file{fieldName: "audio", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingAudio
				Audio string `json:"audio"`
			}{
				OutgoingAudio: *s,
				Audio:         s.fileID,
			}
			_, err = api.c.postJSON(sendAudio, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	case *OutgoingDocument:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendDocument, resp, file{fieldName: "document", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingDocument
				Document string `json:"document"`
			}{
				OutgoingDocument: *s,
				Document:         s.fileID,
			}
			_, err = api.c.postJSON(sendDocument, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	case *OutgoingSticker:
		if s.filePath != "" {
			_, err = api.c.uploadFile(sendSticker, resp, file{fieldName: "sticker", path: s.filePath}, s)
		} else if s.fileID != "" {
			toSend := struct {
				OutgoingSticker
				Sticker string `json:"sticker"`
			}{
				OutgoingSticker: *s,
				Sticker:         s.fileID,
			}
			_, err = api.c.postJSON(sendSticker, resp, toSend)
		} else {
			return nil, ErrNoFileSpecified
		}
	default:
		panic(fmt.Sprintf("tbotapi: internal: unexpected type for send(): %T", s))
	}

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
