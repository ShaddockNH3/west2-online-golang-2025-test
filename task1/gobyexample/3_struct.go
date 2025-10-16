package gobyexample

import "fmt"

var nextID = 1

type Animal struct {
	id        int
	Name      string
	IsAdopted bool
}

func (a *Animal) ID() int {
	return a.id
}

type ShelterAnimal interface {
	ID() int
	GetDetails() string
	PerformCheckup()
	Adopt()
}

type ShelterCat struct{ Animal }

func (c *ShelterCat) GetDetails() string {
	details := fmt.Sprintf("ID: %d, 可爱的猫咪: %s", c.ID(), c.Name)
	if c.IsAdopted {
		details += " (已被领养)"
	}
	return details
}
func (c *ShelterCat) PerformCheckup() { fmt.Println("为", c.Name, "检查了牙齿！") }
func (c *ShelterCat) Adopt()          { c.IsAdopted = true }

type ShelterDog struct{ Animal }

func (d *ShelterDog) GetDetails() string {
	details := fmt.Sprintf("ID: %d, 忠诚的狗狗: %s", d.ID(), d.Name)
	if d.IsAdopted {
		details += " (已被领养)"
	}
	return details
}
func (d *ShelterDog) PerformCheckup() { fmt.Println("给", d.Name, "测了心率！") }
func (d *ShelterDog) Adopt()          { d.IsAdopted = true }

func NewShelterCat(name string) *ShelterCat {
	c := ShelterCat{Animal: Animal{id: nextID, Name: name, IsAdopted: false}}
	nextID++
	return &c
}
func NewShelterDog(name string) *ShelterDog {
	d := ShelterDog{Animal: Animal{id: nextID, Name: name, IsAdopted: false}}
	nextID++
	return &d
}

type Shelter struct {
	Name    string
	Animals map[int]ShelterAnimal
}

func (s *Shelter) AdmitAnimal(animal ShelterAnimal) {
	s.Animals[animal.ID()] = animal
	fmt.Println(animal.GetDetails(), "来到了收容所！")
}

func (s *Shelter) ListAnimals() {
	fmt.Printf("\n--- %s 的动物列表 ---\n", s.Name)
	for _, animal := range s.Animals {
		fmt.Println(animal.GetDetails())
	}
}

func test3() {
	shelter := Shelter{
		Name:    "喵汪之家",
		Animals: make(map[int]ShelterAnimal),
	}

	shelter.AdmitAnimal(NewShelterCat("咪咪"))
	shelter.AdmitAnimal(NewShelterDog("旺财"))

	shelter.ListAnimals()

	animalToAdopt := shelter.Animals[1]
	if animalToAdopt != nil {
		fmt.Println("\n...", animalToAdopt.GetDetails(), "准备被领养...")
		animalToAdopt.Adopt()
	}

	fmt.Println("\n--- 咪咪被领养后 ---")
	shelter.ListAnimals()
}
