package ayreon

import "fmt"

// Castle is electric
type Castle struct {
	guests []string
}

func (c *Castle) Enter(guest string) string {
	c.guests = append(c.guests, guest)
	return fmt.Sprintf("%s enters the Electric Castle", guest)
}

func (c *Castle) GetGuests() []string {
	return c.guests
}

func (c *Castle) GetTraps() string {
	return "Many traps there!"
}
