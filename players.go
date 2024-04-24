package palworld

type Player struct {
	Name      string  `json:"name"`
	LocationX float64 `json:"location_x"`
	LocationY float64 `json:"location_y"`
	Level     int     `json:"level"`
	PlayerId  string  `json:"playerId"`
	UserId    string  `json:"userId"`
	IP        string  `json:"ip"`
	Ping      float64 `json:"ping"`
}

type Players struct {
	Players []*Player `json:"players"`
}

func (c *Client) Players() (*Players, error) {
	apiPath := "/v1/api/players"

	var v Players
	if err := c.get(apiPath, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
