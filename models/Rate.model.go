package models

type Rate struct {
	Rate float64 `json:"rate"`
}

func (c *Rate) Clone() *Rate {
	return &Rate{
		Rate: c.Rate,
	}
}
