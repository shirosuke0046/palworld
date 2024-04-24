package palworld

type BanRequestBody struct {
	UserID  string `json:"userid"`
	Message string `json:"message"`
}

func (c *Client) Ban(userId string, message string) error {
	apiPath := "/v1/api/ban"

	body := &BanRequestBody{
		UserID:  userId,
		Message: message,
	}

	if err := c.post(apiPath, body); err != nil {
		return err
	}

	return nil
}
