package misc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// To get the server running - "pip install flask && python3 app.py"
// *I'm to lazy to re-write this in Go atm (I wrote the flask app a while back)

/*
assume there is a REST API available at http://www.employee.com/api for accessing employee information The employee information endpoint is /employee/<id> Each employee record you retrieve will be a JSON object with the following keys:

    name refers to a String that contains the employee’s first and last name
    title refers to a String that contains the employee’s job title
    reports refers to an Array of Strings containing the IDs of the employee’s direct reports

Write a function that will take an employee ID and print out the entire hierarchy of employees under that employee. For example, suppose that Flynn Mackie’s employee id is 'A123456789' and his only direct reports are Wesley Thomas and Nina Chiswick. If you provide 'A123456789' as input to your function, you will see the sample output below.

Flynn Mackie - Senior VP of Engineering
    Wesley Thomas - VP of Design
        Randall Cosmo - Director of Design
            Brenda Plager - Senior Designer
    Nina Chiswick - VP of Engineering
        Tommy Quinn - Director of Engineering
            Jake Farmer - Frontend Manager
                Liam Freeman - Junior Code Monkey
            Sheila Dunbar - Backend Manager
                Peter Young - Senior Code Cowboy
*/

type Employee struct {
	Name    string
	Title   string
	Reports []string
}

func getEmployeeInfo(id string) (employee *Employee, err error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:37899/linkedin/api/employee/%s", id), nil)
	if err != nil {
		return
	}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &employee)
	if err != nil {
		return
	}
	return
}

func PrintEmployeeHeirarchy(id string, numSpace int) {
	e, _ := getEmployeeInfo(id)
	spaceChars := ""
	for i := 0; i < numSpace; i++ {
		spaceChars += " "
	}
	fmt.Printf("%s%s - %s\n", spaceChars, e.Name, e.Title)
	for _, employeeId := range e.Reports {
		PrintEmployeeHeirarchy(employeeId, numSpace+2)
	}
}
