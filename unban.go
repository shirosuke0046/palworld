package palworld

type UnBanRequestBody struct {
	UserID string `json:"userid"`
}

func (c *Client) UnBan(userId string) error {
	apiPath := "/v1/api/unban"

	body := &UnBanRequestBody{
		UserID: userId,
	}

	if err := c.post(apiPath, body); err != nil {
		return err
	}

	return nil
}
