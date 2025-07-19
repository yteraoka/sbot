package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	apiBaseURL = "https://api.switch-bot.com"
)

// Client is a client for the SwitchBot API.
type Client struct {
	token      string
	secret     string
	httpClient *http.Client
}

// NewClient creates a new SwitchBot API client.
func NewClient(token, secret string) *Client {
	return &Client{
		token:      token,
		secret:     secret,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// AuthHeaders represents the authentication headers for the SwitchBot API.
type AuthHeaders struct {
	Token string
	Sign  string
	Nonce string
	T     string
}

func (c *Client) newAuthHeaders() AuthHeaders {
	t := strconv.FormatInt(time.Now().UnixMilli(), 10)
	nonce := uuid.New().String()
	data := c.token + t + nonce
	h := hmac.New(sha256.New, []byte(c.secret))
	h.Write([]byte(data))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return AuthHeaders{
		Token: c.token,
		Sign:  sign,
		Nonce: nonce,
		T:     t,
	}
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := apiBaseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	headers := c.newAuthHeaders()
	req.Header.Set("Authorization", headers.Token)
	req.Header.Set("sign", headers.Sign)
	req.Header.Set("nonce", headers.Nonce)
	req.Header.Set("t", headers.T)
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	return req, nil
}

// APIResponseBody is the generic response body from the SwitchBot API.
type APIResponseBody struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Body       json.RawMessage `json:"body"`
}

// Device represents a single device.
type Device struct {
	ID           string `json:"deviceId"`
	Name         string `json:"deviceName"`
	Type         string `json:"deviceType"`
	CloudService bool   `json:"enableCloudService"`
	HubDeviceID  string `json:"hubDeviceId"`
}

// Scene represents a single scene.
type Scene struct {
	ID   string `json:"sceneId"`
	Name string `json:"sceneName"`
}

// ListDevicesResponse is the response body for the list devices endpoint.
type ListDevicesResponse struct {
	DeviceList         []Device `json:"deviceList"`
	InfraredRemoteList []Device `json:"infraredRemoteList"`
}

// ListDevices fetches the list of devices.
func (c *Client) ListDevices() (*ListDevicesResponse, error) {
	req, err := c.newRequest("GET", "/v1.1/devices", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponseBody
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if apiResp.StatusCode != 100 {
		return nil, fmt.Errorf("API error: %s (status code: %d)", apiResp.Message, apiResp.StatusCode)
	}

	var listResp ListDevicesResponse
	if err := json.Unmarshal(apiResp.Body, &listResp); err != nil {
		return nil, fmt.Errorf("failed to parse device list: %w", err)
	}

	return &listResp, nil
}

// ListScenes fetches the list of scenes.
func (c *Client) ListScenes() ([]Scene, error) {
	req, err := c.newRequest("GET", "/v1.1/scenes", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponseBody
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if apiResp.StatusCode != 100 {
		return nil, fmt.Errorf("API error: %s (status code: %d)", apiResp.Message, apiResp.StatusCode)
	}

	var scenes []Scene
	if err := json.Unmarshal(apiResp.Body, &scenes); err != nil {
		return nil, fmt.Errorf("failed to parse scene list: %w", err)
	}

	return scenes, nil
}

// ExecuteScene executes a specific scene.
func (c *Client) ExecuteScene(sceneID string) error {
	path := fmt.Sprintf("/v1.1/scenes/%s/execute", sceneID)
	req, err := c.newRequest("POST", path, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResp APIResponseBody
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		return fmt.Errorf("failed to parse API response: %w. Response: %s", err, string(respBody))
	}

	if apiResp.StatusCode != 100 {
		return fmt.Errorf("API error: %s (status code: %d)", apiResp.Message, apiResp.StatusCode)
	}

	return nil
}

// GetDeviceStatus fetches the status of a specific device.
func (c *Client) GetDeviceStatus(deviceID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v1.1/devices/%s/status", deviceID)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponseBody
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if apiResp.StatusCode != 100 {
		return nil, fmt.Errorf("API error: %s (status code: %d)", apiResp.Message, apiResp.StatusCode)
	}

	return apiResp.Body, nil
}

// CommandRequestBody is the body for sending a command.
type CommandRequestBody struct {
	Command     string `json:"command"`
	Parameter   string `json:"parameter"`
	CommandType string `json:"commandType"`
}

// SendCommand sends a command to a specific device.
func (c *Client) SendCommand(deviceID, command, parameter string) error {
	path := fmt.Sprintf("/v1.1/devices/%s/commands", deviceID)
	cmdBody := &CommandRequestBody{
		Command:     command,
		Parameter:   parameter,
		CommandType: "command",
	}

	jsonBody, err := json.Marshal(cmdBody)
	if err != nil {
		return err
	}

	req, err := c.newRequest("POST", path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResp APIResponseBody
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		return fmt.Errorf("failed to parse API response: %w. Response: %s", err, string(respBody))
	}

	if apiResp.StatusCode != 100 {
		return fmt.Errorf("API error: %s (status code: %d)", apiResp.Message, apiResp.StatusCode)
	}

	return nil
}

// GetDeviceID resolves a device name or ID to a device ID.
// It first tries to match by ID, then by name.
// If a name matches multiple devices, an error is returned.
func (c *Client) GetDeviceID(nameOrID string) (string, error) {
	devices, err := c.ListDevices()
	if err != nil {
		return "", fmt.Errorf("could not list devices to find ID: %w", err)
	}

	allDevices := append(devices.DeviceList, devices.InfraredRemoteList...)

	// First pass: check for exact ID match
	for _, device := range allDevices {
		if device.ID == nameOrID {
			return device.ID, nil
		}
	}

	// Second pass: check for name match
	var foundDeviceID string
	var foundCount int
	for _, device := range allDevices {
		if device.Name == nameOrID {
			foundDeviceID = device.ID
			foundCount++
		}
	}

	if foundCount == 1 {
		return foundDeviceID, nil
	}
	if foundCount > 1 {
		return "", fmt.Errorf("multiple devices found with name '%s', please use device ID instead", nameOrID)
	}

	return "", fmt.Errorf("no device found with name or ID '%s'", nameOrID)
}

// GetSceneID resolves a scene name or ID to a scene ID.
func (c *Client) GetSceneID(nameOrID string) (string, error) {
	scenes, err := c.ListScenes()
	if err != nil {
		return "", fmt.Errorf("could not list scenes to find ID: %w", err)
	}

	// First pass: check for exact ID match
	for _, scene := range scenes {
		if scene.ID == nameOrID {
			return scene.ID, nil
		}
	}

	// Second pass: check for name match
	var foundSceneID string
	var foundCount int
	for _, scene := range scenes {
		if scene.Name == nameOrID {
			foundSceneID = scene.ID
			foundCount++
		}
	}

	if foundCount == 1 {
		return foundSceneID, nil
	}
	if foundCount > 1 {
		return "", fmt.Errorf("multiple scenes found with name '%s', please use scene ID instead", nameOrID)
	}

	return "", fmt.Errorf("no scene found with name or ID '%s'", nameOrID)
}
