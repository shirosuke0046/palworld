package palworld

import "fmt"

type ShutdownRequestBody struct {
	WaitTime int    `json:"waittime"`
	Message  string `json:"message"`
}

func (c *Client) Shutdown(waitTime int, message string) error {
	if waitTime < 1 {
		return fmt.Errorf("invalid value for 'waitTime' - must be a positive integer")
	}

	apiPath := "/v1/api/shutdown"

	body := &ShutdownRequestBody{
		WaitTime: waitTime,
		Message:  message,
	}

	if err := c.post(apiPath, body); err != nil {
		return err
	}

	return nil
}
