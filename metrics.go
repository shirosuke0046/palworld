package palworld

type Metrics struct {
	CurrentPlayerNum int     `json:"currentplayernum"`
	MaxPlayerNum     int     `json:"maxplayernum"`
	ServerFPS        int     `json:"serverfps"`
	ServerFrameTime  float64 `json:"serverframetime"`
	Uptime           int     `json:"uptime"`
}

func (c *Client) Metrics() (*Metrics, error) {
	apiPath := "/v1/api/metrics"

	var v Metrics
	if err := c.get(apiPath, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
