package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Richard": {Name: "Richard", LastName: "Pickman", Age: 45, Gender: "male", Crimes: 10},
		"Sukuna":  {Name: "Sukuna", LastName: "Ryōmen", Age: 1000, Gender: "male", Crimes: 1000000},
		"Satoru":  {Name: "Satoru", LastName: "Gojo", Age: 28, Gender: "male", Crimes: 0},
		"Sam":     {Name: "Sam", LastName: "Serious", Age: 40, Gender: "male", Crimes: 6},
		"Erza":    {Name: "Erza", LastName: "Scarlet", Age: 19, Gender: "female", Crimes: 0},
		"Anya":    {Name: "Anya", LastName: "Forger", Age: 6, Gender: "female", Crimes: 0},
		"Toga":    {Name: "Toga", LastName: "Himiko", Age: 16, Gender: "female", Crimes: 256},
		"Eren":    {Name: "Eren", LastName: "Yeager", Age: 22, Gender: "male", Crimes: 1},
		"Rick":    {Name: "Rick", LastName: "Sanchez", Age: 70, Gender: "male", Crimes: 0},
		"Mikasa":  {Name: "Mikasa", LastName: "Ackerman", Age: 22, Gender: "female", Crimes: 0},
	}
	//"MostCriminal" and "MostCriminalFound" to keep track of the person with the most crimes and whether such a person was found in the "people" map.
	var MostCriminal Man
	var MostCriminalFound bool
	//Сhecks if each name exists as a key in the "people" map.
	//If it does, it compares the number of crimes of that person with the current "MostCriminal" and updates the variables accordingly.
	suspects := []string{"Richard", "Sukuna", "Sam", "Eren", "Toga"}
	for _, name := range suspects {
		person, ok := people[name]
		if !ok {
			continue
		}
		if person.Crimes > MostCriminal.Crimes {
			MostCriminal = person
			MostCriminalFound = true
		}
	}
	if MostCriminalFound {
		fmt.Printf("The most criminal personality: %s %s", MostCriminal.Name, MostCriminal.LastName)
	} else {
		fmt.Println("There is no information on the requested suspects in the database")
	}
}
