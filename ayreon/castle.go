package ayreon

import "fmt"

// Castle is electric
type Castle struct {
	guests []Character
}

func (c *Castle) Enter(guest Character) string {
	c.guests = append(c.guests, guest)
	return fmt.Sprintf("%s enters the Electric Castle", guest.Name)
}

func (c *Castle) GetGuests() []Character {
	return c.guests
}

func (c *Castle) GetTraps() string {
	return "Many traps there!"
}
