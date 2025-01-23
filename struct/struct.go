package main

import "fmt"

type Student struct {
	id    string
	name  string
	age   int
	phone string
}

func main() {
	// NOTE - student
	student := Student{
		id:    "1",
		name:  "Big",
		age:   18,
		phone: "2059296742",
	}
	fmt.Println("student:", student)

	//NOTE - student list array
	studentListArray := [3]Student{}
	studentListArray[0] = Student{
		id:    "1",
		name:  "Ton",
		age:   22,
		phone: "2055736369",
	}
	studentListArray[1] = Student{
		id:    "2",
		name:  "Phon",
		age:   20,
		phone: "2055992261",
	}
	studentListArray[2] = Student{
		id:    "3",
		name:  "Kilaeng",
		age:   23,
		phone: "55799663",
	}
	fmt.Println("students:", studentListArray)

	//NOTE - student list slice
	studentListSlice := []Student{}
	student1 := Student{
		id:    "1",
		name:  "Ker",
		age:   20,
		phone: "2055667877",
	}
	student2 := Student{
		id:    "2",
		name:  "Jack",
		age:   23,
		phone: "2078485989",
	}
	student3 := Student{
		id:    "3",
		name:  "Far",
		age:   25,
		phone: "2058965989",
	}

	studentListSlice = append(studentListSlice, student1)
	studentListSlice = append(studentListSlice, student2)
	studentListSlice = append(studentListSlice, student3)

	fmt.Println("studentListSlice:", studentListSlice)
}
