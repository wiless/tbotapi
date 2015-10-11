package tbotapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v0"
	"net/http"
)

type method string

const (
	m_GetMe                = method("GetMe")
	m_SendMessage          = method("SendMessage")
	m_ForwardMessage       = method("ForwardMessage")
	m_SendPhoto            = method("SendPhoto")
	m_SendAudio            = method("SendAudio")
	m_SendDocument         = method("SendDocument")
	m_SendSticker          = method("SendSticker")
	m_SendVideo            = method("SendVideo")
	m_SendVoice            = method("SendVoice")
	m_SendLocation         = method("SendLocation")
	m_SendChatAction       = method("SendChatAction")
	m_GetUserProfilePhotos = method("GetUserProfilePhotos")
	m_GetUpdates           = method("GetUpdates")
	m_SetWebhook           = method("SetWebhook")
	m_GetFile              = method("GetFile")
)

type client struct {
	c         *resty.Client
	endpoints map[method]string
}

func newClient(baseURI string) *client {
	toReturn := &client{
		c:         resty.New().SetHTTPMode().OnAfterResponse(parseResponseBody).OnAfterResponse(checkHTTPStatus).SetDebug(true),
		endpoints: createEndpoints(baseURI),
	}

	return toReturn
}

func (c *client) get(m method, result interface{}) (*resty.Response, error) {
	return c.c.R().SetResult(result).Get(c.getEndpoint(m))
}

func (c *client) getQuerystring(m method, result interface{}, querystring map[string]string) (*resty.Response, error) {
	return c.c.R().SetQueryParams(querystring).SetResult(result).Get(c.getEndpoint(m))
}

func (c *client) postJSON(m method, result interface{}, data interface{}) (*resty.Response, error) {
	return c.c.R().SetBody(data).SetResult(result).Post(c.getEndpoint(m))
}

func (c *client) uploadFile(m method, result interface{}, data file, fields encodable) (*resty.Response, error) {
	return c.c.R().SetFile(data.fieldName, data.path).SetResult(result).SetFormData(map[string]string(fields.GetQueryString())).Post(c.getEndpoint(m))
}

func parseResponseBody(c *resty.Client, res *resty.Response) (err error) {
	// Handles only JSON
	ct := res.Header().Get(http.CanonicalHeaderKey("Content-Type"))
	if resty.IsJSONType(ct) {
		// Considered as Result
		if res.StatusCode() > 199 && res.StatusCode() < 500 {
			if res.Request.Result != nil {
				err = json.Unmarshal(res.Body, res.Request.Result)
			}
		}
	}

	return
}

func checkHTTPStatus(_ *resty.Client, res *resty.Response) error {
	if res.StatusCode() >= 500 {
		return fmt.Errorf("API: Server error: returned %s when requesting %s", res.Status(), res.Request.URL)
	}
	return nil
}

func (c *client) getEndpoint(method method) string {
	endpoint, ok := c.endpoints[method]
	if !ok {
		panic(fmt.Errorf("tbotapi: internal: Endpoint for method %s not found", string(method)))
	}
	return endpoint
}

func createEndpoints(baseURI string) map[method]string {
	toReturn := map[method]string{}

	toReturn[m_GetMe] = fmt.Sprint(baseURI, "/", string(m_GetMe))
	toReturn[m_SendMessage] = fmt.Sprint(baseURI, "/", string(m_SendMessage))
	toReturn[m_ForwardMessage] = fmt.Sprint(baseURI, "/", string(m_ForwardMessage))
	toReturn[m_SendPhoto] = fmt.Sprint(baseURI, "/", string(m_SendPhoto))
	toReturn[m_SendAudio] = fmt.Sprint(baseURI, "/", string(m_SendAudio))
	toReturn[m_SendDocument] = fmt.Sprint(baseURI, "/", string(m_SendDocument))
	toReturn[m_SendSticker] = fmt.Sprint(baseURI, "/", string(m_SendSticker))
	toReturn[m_SendVideo] = fmt.Sprint(baseURI, "/", string(m_SendVideo))
	toReturn[m_SendVoice] = fmt.Sprint(baseURI, "/", string(m_SendVoice))
	toReturn[m_SendLocation] = fmt.Sprint(baseURI, "/", string(m_SendLocation))
	toReturn[m_SendChatAction] = fmt.Sprint(baseURI, "/", string(m_SendChatAction))
	toReturn[m_GetUserProfilePhotos] = fmt.Sprint(baseURI, "/", string(m_GetUserProfilePhotos))
	toReturn[m_GetUpdates] = fmt.Sprint(baseURI, "/", string(m_GetUpdates))
	toReturn[m_SetWebhook] = fmt.Sprint(baseURI, "/", string(m_SetWebhook))
	toReturn[m_GetFile] = fmt.Sprint(baseURI, "/", string(m_GetFile))

	return toReturn
}
