package tbotapi

import (
	"bitbucket.org/mrd0ll4r/tbotapi/model"
	"fmt"
	"github.com/jmcvetta/napping"
	"io"
	"menteslibres.net/gosexy/rest"
	"net/url"
	"sync"
	"time"
)

// A TelegramBotAPI is an API Client for one Telegram bot.
// Create a new client by calling the New() function.
type TelegramBotAPI struct {
	ID       int                // the bots ID
	Name     string             // the bots Name as seen by users
	Username string             // the bots username
	Updates  chan *model.Update // a channel providing updates this bot receives
	Errors   chan error         // a channel providing errors that occur during the retrieval of updates
	baseURI  string
	closed   chan struct{}
	wg       sync.WaitGroup
	session  napping.Session
}

const apiBaseURI string = "https://api.telegram.org/bot%s"

// New creates a new API Client for a Telegram bot using the apiKey provided.
// It will call the GetMe method to retrieve the bots id, name and username.
// Additionally, an update loop is started, pumping updates into the Updates channel.
func New(apiKey string) (*TelegramBotAPI, error) {
	toReturn := TelegramBotAPI{
		baseURI: fmt.Sprintf(apiBaseURI, apiKey),
		Updates: make(chan *model.Update),
		Errors:  make(chan error),
		closed:  make(chan struct{}),
		session: napping.Session{},
	}
	user, err := toReturn.GetMe()
	if err != nil {
		return nil, err
	}
	toReturn.ID = user.User.Id
	toReturn.Name = user.User.FirstName
	toReturn.Username = user.User.Username

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
	var offset int

	for {
		select {
		case <-api.closed:
			api.wg.Done()
			return
		default:
		}

		if err != nil {
			if err.Error() != "Could not convert response []uint8 to model.UpdateResponse." {
				api.Errors <- err
			} else {
				time.Sleep(time.Duration(10) * time.Second)
			}
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

func putUpdatesInChannel(channel chan *model.Update, updates []model.Update) int {
	highestOffset := -1
	for _, update := range updates {
		highestOffset = update.Id
		channel <- &update
	}

	return highestOffset
}

func (api *TelegramBotAPI) getUpdates() (*model.UpdateResponse, error) {
	resp := &model.UpdateResponse{}
	querystring := url.Values{}
	querystring.Set("timeout", fmt.Sprint(60))
	_, err := api.session.Get(fmt.Sprint(api.baseURI, "/GetUpdates"), &querystring, resp, resp)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
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
	_, err := api.session.Get(fmt.Sprint(api.baseURI, "/GetUpdates"), &querystring, resp, resp)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMe returns basic information about the bot in form of a UserResponse.
func (api *TelegramBotAPI) GetMe() (*model.UserResponse, error) {
	resp := &model.UserResponse{}
	_, err := api.session.Get(fmt.Sprint(api.baseURI, "/GetMe"), nil, resp, resp)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendMessage sends a text message to the chatID specified, with the given text.
// For more options, use the SendMessageExtended function.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendMessage(chatID int, text string) (*model.MessageResponse, error) {
	return api.SendMessageExtended(model.NewOutgoingMessage(chatID, text))
}

// SendMessageExtended sends a text message with additional options.
// Use NewOutgoingMessage to construct the outgoing message.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendMessageExtended(om *model.OutgoingMessage) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendMessage"), om.GetPub(), resp, resp)
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
func (api *TelegramBotAPI) ForwardMessage(toChatID, fromChatID, messageID int) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingBasePub
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}{
		OutgoingBasePub: model.OutgoingBasePub{
			ChatId: toChatID,
		},
		FromChatId: fromChatID,
		MessageId:  messageID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/ForwardMessage"), toSend, resp, resp)
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
func (api *TelegramBotAPI) ResendPhoto(op *model.OutgoingPhoto, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingPhotoPub
		Photo string `json:"photo"`
	}{
		OutgoingPhotoPub: op.GetPub(),
		Photo:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendPhoto"), toSend, resp, resp)
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
// Use NewOutgoingPhoto to construct the outgoing photo message, specify an io.Reader and a fileName for the file.
// Note, that the Telegram API will check the filename for its extension and will reject non-image files.
// On success, the sent message is returned as a MessageResponse.
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendPhoto"), message)
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
func (api *TelegramBotAPI) ResendVoice(ov *model.OutgoingVoice, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingVoicePub
		Audio string `json:"audio"`
	}{
		OutgoingVoicePub: ov.GetPub(),
		Audio:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendVoice"), toSend, resp, resp)
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
// Use NewOutgoingVoice to construct the voice message, specify an io.Reader and a fileName.
// Note that the Telegram servers check the extension of the file name and will reject non-audio files.
// Check the current API documentation for the file types accepted.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendVoice(ov *model.OutgoingVoice, file io.Reader, fileName string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	files := rest.FileMap{
		"audio": []rest.File{
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendVoice"), message)
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
func (api *TelegramBotAPI) ResendAudio(oa *model.OutgoingAudio, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingAudioPub
		Audio string `json:"audio"`
	}{
		OutgoingAudioPub: oa.GetPub(),
		Audio:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendAudio"), toSend, resp, resp)
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
// Use NewOutgoingAudio to construct the audio message, specify an io.Reader and a fileName.
// Note that the Telegram servers check the extension of the file name and will reject non-audio files.
// Check the current API documentation for the file types accepted.
// On success, the sent message is returned as a MessageResponse.
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendAudio"), message)
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
func (api *TelegramBotAPI) ResendDocument(od *model.OutgoingDocument, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingDocumentPub
		Document string `json:"document"`
	}{
		OutgoingDocumentPub: od.GetPub(),
		Document:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendDocument"), toSend, resp, resp)
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
// Use NewOutgoingDocument to construct the message, specify an io.Reader and a fileName.
// For current limitations on what a bot can send, check the bot API documentation.
// On success, the sent message is returned as a MessageResponse.
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendDocument"), message)
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
func (api *TelegramBotAPI) ResendSticker(os *model.OutgoingSticker, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingStickerPub
		Sticker string `json:"sticker"`
	}{
		OutgoingStickerPub: os.GetPub(),
		Sticker:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendSticker"), toSend, resp, resp)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendSticker sends a sticker that is not already on the Telegram servesr.
// Use NewOutgoingSticker to construct the message, specify an io.Reader and a fileName.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what a bot can send, check the API documentation.
// On success, the sent message is returned as a MessageResponse.
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendSticker"), message)
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
func (api *TelegramBotAPI) ResendVideo(ov *model.OutgoingVideo, fileID string) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}
	toSend := struct {
		model.OutgoingVideoPub
		Video string `json:"video"`
	}{
		OutgoingVideoPub: ov.GetPub(),
		Video:            fileID,
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendVideo"), toSend, resp, resp)
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
// Use OutgoingVideo to construct the message, specify an io.Reader and a fileName.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
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

	err = rest.PostMultipart(resp, fmt.Sprint(api.baseURI, "/SendVideo"), message)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendLocation sends a location.
// Use NewOutgoingLocation to construct the message to send.
// On success, the sent message is returned as a MessageResponse.
func (api *TelegramBotAPI) SendLocation(ol *model.OutgoingLocation) (*model.MessageResponse, error) {
	resp := &model.MessageResponse{}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendLocation"), ol.GetPub(), resp, resp)
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
func (api *TelegramBotAPI) SendChatAction(chatID int, action model.ChatAction) (*model.BaseResponse, error) {
	resp := &model.BaseResponse{}
	toSend := struct {
		model.OutgoingBasePub
		Action string `json:"action"`
	}{
		OutgoingBasePub: model.OutgoingBasePub{
			ChatId: chatID,
		},
		Action: string(action),
	}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/SendChatAction"), toSend, resp, resp)
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
func (api *TelegramBotAPI) GetProfilePhotos(op *model.OutgoingUserProfilePhotosRequest) (*model.UserProfilePhotosResponse, error) {
	resp := &model.UserProfilePhotosResponse{}

	_, err := api.session.Post(fmt.Sprint(api.baseURI, "/GetUserProfilePhotos"), op.GetPub(), resp, resp)
	if err != nil {
		return nil, err
	}
	err = check(&resp.BaseResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func check(br *model.BaseResponse) error {
	if br.Ok {
		return nil
	}

	return fmt.Errorf("tbotapi: API error: %d - %s", br.ErrorCode, br.Description)
}
