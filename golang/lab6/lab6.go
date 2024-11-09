package lab6

import (
	"fmt"
)

type Character struct {
	Name   string
	Class  string
	Level  int
	Health float64
}

func (c *Character) LevelUp() {
	c.Level++
	fmt.Printf("%s достиг уровня %d!\n", c.Name, c.Level)
}

func (c *Character) Heal(amount float64) {
	c.Health += amount
	fmt.Printf("%s восстановил %.2f здоровья! Теперь %.2f здоровья.\n", c.Name, amount, c.Health)
}

func (c *Character) DisplayInfo() {
	fmt.Printf("Имя: %s\nКласс: %s\nУровень: %d\nЗдоровье: %.2f\n", c.Name, c.Class, c.Level, c.Health)
}

func RunLab6() {
	hero := Character{
		Name:   "Арсен",
		Class:  "Воин",
		Level:  10,
		Health: 100.0,
	}

	hero.DisplayInfo()

	hero.LevelUp()

	hero.Heal(15.0)

	hero.DisplayInfo()
}
