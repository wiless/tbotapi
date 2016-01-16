package tbotapi

import (
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

// ForwardMessage forwards a message with ID messageID from the fromChatID to the toChatID chat.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ForwardMessage(of *OutgoingForward) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.postJSON(forwardMessage, resp, of)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendPhoto resends a photo that is already on the Telegram servers by fileID.
// Use NewOutgoingPhoto to construct the outgoing photo message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendPhoto(op *OutgoingPhoto, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingPhoto
		Photo string `json:"photo"`
	}{
		OutgoingPhoto: *op,
		Photo:         fileID,
	}
	_, err := api.c.postJSON(sendPhoto, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendPhoto sends a photo message with a photo that is not yet on the Telegram servers.
// Use NewOutgoingPhoto to construct the outgoing photo message and specify the path to the file.
// Note, that the Telegram API will check the filename for its extension and will reject non-image files.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendPhoto(op *OutgoingPhoto, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendPhoto, resp, file{fieldName: "photo", path: filePath}, op)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendVoice resends a voice message that is already on the Telegram servers by fileID.
// Use NewOutgoingVoice to construct the voice message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendVoice(ov *OutgoingVoice, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingVoice
		Audio string `json:"audio"`
	}{
		OutgoingVoice: *ov,
		Audio:         fileID,
	}
	_, err := api.c.postJSON(sendVoice, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendVoice sends a voice message with the contents not already on the Telegram servers.
// Use NewOutgoingVoice to construct the voice message and specify the path to the file.
// Note that the Telegram servers check the extension of the file name and will reject non-audio files.
// Check the current API documentation for the file types accepted.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendVoice(ov *OutgoingVoice, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendVoice, resp, file{fieldName: "audio", path: filePath}, ov)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendAudio resends audio that is already on the Telegram servers by fileID.
// Use NewOutgoingAudio to construct the audio message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendAudio(oa *OutgoingAudio, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingAudio
		Audio string `json:"audio"`
	}{
		OutgoingAudio: *oa,
		Audio:         fileID,
	}
	_, err := api.c.postJSON(sendAudio, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendAudio sends an audio message with the contents not already on the Telegram servers.
// Use NewOutgoingAudio to construct the audio message and specify the path to the file.
// Note that the Telegram servers check the extension of the file name and will reject non-audio files.
// Check the current API documentation for the file types accepted.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendAudio(oa *OutgoingAudio, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendAudio, resp, file{fieldName: "audio", path: filePath}, oa)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendDocument resends a general file that is already on the Telegram servers by fileID.
// Use NewOutgoingDocument to construct the message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendDocument(od *OutgoingDocument, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingDocument
		Document string `json:"document"`
	}{
		OutgoingDocument: *od,
		Document:         fileID,
	}
	_, err := api.c.postJSON(sendDocument, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendDocument sends a general file that is not already on the Telegram servers.
// Use NewOutgoingDocument to construct the message and specify the path to the file.
// For current limitations on what a bot can send, check the bot API documentation.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendDocument(od *OutgoingDocument, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendDocument, resp, file{fieldName: "document", path: filePath}, od)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendSticker resends a sticker that is already on the Telegram servers by fileID.
// Use NewOutgoingSticker to construct the message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendSticker(os *OutgoingSticker, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingSticker
		Sticker string `json:"sticker"`
	}{
		OutgoingSticker: *os,
		Sticker:         fileID,
	}
	_, err := api.c.postJSON(sendSticker, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendSticker sends a sticker that is not already on the Telegram server.
// Use NewOutgoingSticker to construct the message and specify the path to the file.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what a bot can send, check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendSticker(os *OutgoingSticker, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendSticker, resp, file{fieldName: "sticker", path: filePath}, os)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResendVideo resends a video that is already on the Telegram servers by fileID.
// Use NewOutgoingVideo to construct the video message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) ResendVideo(ov *OutgoingVideo, fileID string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	toSend := struct {
		OutgoingVideo
		Video string `json:"video"`
	}{
		OutgoingVideo: *ov,
		Video:         fileID,
	}
	_, err := api.c.postJSON(sendVideo, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendVideo sends a video that is not already on the Telegram servers.
// Use OutgoingVideo to construct the message and specify the path to the file.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendVideo(ov *OutgoingVideo, filePath string) (*MessageResponse, error) {
	resp := &MessageResponse{}
	_, err := api.c.uploadFile(sendVideo, resp, file{fieldName: "video", path: filePath}, ov)

	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendChatAction sends a chat action to the specified chatID.
// Use the ChatAction constants to specify the action.
// On success, a BaseResponse is returned.
func (api *TelegramBotAPI) SendChatAction(recipient Recipient, action ChatAction) (*BaseResponse, error) {
	resp := &BaseResponse{}
	toSend := struct {
		OutgoingBase
		Action string `json:"action"`
	}{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		Action: string(action),
	}
	_, err := api.c.postJSON(sendChatAction, resp, toSend)

	if err != nil {
		return nil, err
	}
	err = check(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetProfilePhotos gets a users profile pictures.
// Use NewOutgoingUserProfilePhotosRequest to create the request.
// On success, the photos are returned as a UserProfilePhotosResponse.
func (api *TelegramBotAPI) GetProfilePhotos(op *OutgoingUserProfilePhotosRequest) (*UserProfilePhotosResponse, error) {
	resp := &UserProfilePhotosResponse{}
	_, err := api.c.postJSON(getUserProfilePhotos, resp, op)

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

func (api *TelegramBotAPI) send(s sendable) (resp *MessageResponse, err error) {
	resp = &MessageResponse{}

	switch s := s.(type) {
	case *OutgoingMessage:
		_, err = api.c.postJSON(sendMessage, resp, s)
	case *OutgoingLocation:
		_, err = api.c.postJSON(sendLocation, resp, s)
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
