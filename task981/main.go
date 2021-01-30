package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func (m Man) Print() {
	fmt.Println("Name:", m.Name, m.LastName)
	fmt.Println("Age:", m.Age)
	fmt.Println("Gender:", m.Gender)
	fmt.Println("Number of crimes:", m.Crimes)
}

func main() {

	people := make(map[string]Man)
	people["John"] = Man{Name: "John", LastName: "Dillinger", Age: 31, Gender: "Male", Crimes: 27}
	people["Bonnie"] = Man{Name: "Bonnie", LastName: "Parker", Age: 23, Gender: "Female", Crimes: 59}
	people["Clyde"] = Man{Name: "Clyde", LastName: "Barrow", Age: 25, Gender: "Male", Crimes: 61}
	people["Butch"] = Man{Name: "Butch", LastName: "Cassidy", Age: 40, Gender: "Male", Crimes: 21}
	people["Мишка"] = Man{Name: "Мишка", LastName: "Япончик", Age: 27, Gender: "Male", Crimes: 25}

	suspects := []string{"John", "Мишка", "Bonnie", "Butch"}
	var mainSuspect Man
	for _, name := range suspects {
		p, ok := people[name]
		if !ok {
			continue
		}
		if p.Crimes > mainSuspect.Crimes {
			mainSuspect = p
		}
	}
	if (mainSuspect == Man{}) {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	} else {
		fmt.Println("Main suspect")
		mainSuspect.Print()
	}
}
