package palworld

type KickRequestBody struct {
	UserID  string `json:"userid"`
	Message string `json:"message"`
}

func (c *Client) Kick(userId string, message string) error {
	apiPath := "/v1/api/kick"

	body := &KickRequestBody{
		UserID:  userId,
		Message: message,
	}

	if err := c.post(apiPath, body); err != nil {
		return err
	}

	return nil
}
