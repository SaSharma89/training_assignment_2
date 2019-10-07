package main

import (
	"fmt"
)

var empList []Employee

func removeIndex(s []Employee, index int) []Employee {
	return append(s[:index], s[index+1:]...)
}

func getEmployeeIndex() (found bool, index int) {
	var first, last string
	fmt.Println("Enter first name :")
	fmt.Scan(&first)
	fmt.Println("Enter last name")
	fmt.Scan(&last)

	found = false
	for idx, _ := range empList {
		if empList[idx].CheckName(first, last) {
			index = idx
			found = true
			break
		}
	}

	if !found {
		fmt.Println("<<<<< EMPLOYEE RECORD NOT FOUND >>>>>")
	}

	return
}

func main() {
	fmt.Println("This is employee management")
	//var emp Employee
	id := 0
	i := 0
	flag := true
	for flag {
		fmt.Println("\n\n\nChoose any option")
		fmt.Println("1 to add new employee")
		fmt.Println("2 to update Salary details")
		fmt.Println("3 to update contact details")
		fmt.Println("4 to update rating")
		fmt.Println("5 to get emergency contact")
		fmt.Println("6 to detect leave")
		fmt.Println("7 to calculate salary")
		fmt.Println("8 to calculate hike")
		fmt.Println("9 to remove employee")
		fmt.Println("10 to display employee details")
		fmt.Println("11 to add friend")
		fmt.Println("12 to remove friend")
		fmt.Println("13 to display friend list")
		fmt.Println("0 to exit")
		fmt.Scan(&i)

		switch i {
		case 0:
			flag = false
		case 1:
			//! basic details
			id++
			var first, last string
			var newEmp Employee
			fmt.Println("Enter first name :")
			fmt.Scan(&first)
			fmt.Println("Enter last name")
			fmt.Scan(&last)
			newEmp.AddBasicDetails(id, first, last, MAX_LEAVE)

			var day, month, year int
			fmt.Println("Enter date of birth")
			fmt.Println("Enter day(dd) : ")
			fmt.Scan(&day)
			fmt.Println("Enter month(mm) : ")
			fmt.Scan(&month)
			fmt.Println("Enter year(yyyy) : ")
			fmt.Scan(&year)

			//! DOB
			date := Date{day, month, year}
			newEmp.AddDOB(date)

			//! add into list
			empList = append(empList, newEmp)
		case 2:
			found, index := getEmployeeIndex()
			if found {
				var fix, variable int64
				fmt.Println("Enter fix :")
				fmt.Scan(&fix)
				fmt.Println("Enter variable")
				fmt.Scan(&variable)
				//! salary details
				salary := SalaryDetails{fix, variable}
				empList[index].AddSalaryDetails(salary)
			}
		case 3:
			found, index := getEmployeeIndex()
			if found {
				var p, s, e string
				fmt.Println("Enter primary number :")
				fmt.Scan(&p)
				fmt.Println("Enter secondary number :")
				fmt.Scan(&s)
				fmt.Println("Enter emergency number :")
				fmt.Scan(&e)

				//! contact details
				contact := ContactDetails{p, s, e}
				empList[index].AddContactDetails(contact)
			}

		case 4:
			found, index := getEmployeeIndex()
			if found {
				var r float32
				fmt.Println("Enter rating :")
				fmt.Scan(&r)

				//! rating
				empList[index].AddRating(r)
			}
		case 5:
			found, index := getEmployeeIndex()
			if found {
				empList[index].GetEmergencyContact()
			}

		case 6:
			found, index := getEmployeeIndex()
			if found {
				empList[index].LeaveTaken()
			}
		case 7:
			found, index := getEmployeeIndex()
			if found {
				empList[index].CalculateNetSalary()
			}
		case 8:
			found, index := getEmployeeIndex()
			if found {
				empList[index].CalculateSalaryHike()

			}
		case 9:
			found, index := getEmployeeIndex()
			if found {
				empList = removeIndex(empList, index)
			}
		case 10:
			found, index := getEmployeeIndex()
			if found {
				fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<employee details >>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				empList[index].DisplayDetails()
			}
		case 11:
			found, index := getEmployeeIndex()
			if found {
				var eid int
				var name string
				fmt.Println("Enter eid :")
				fmt.Scan(&eid)
				fmt.Println("Enter name")
				fmt.Scan(&name)
				empList[index].AddCloseFriend( eid, name)
			}
		case 12 :
			found, index := getEmployeeIndex()
			if found {
				var eid int
				fmt.Println("Enter eid :")
				fmt.Scan(&eid)
				empList[index].RemoveCloseFriend(eid)
			}
		case 13 :
			found, index := getEmployeeIndex()
			if found{
				empList[index].DisplayCloseFriend()
			}
		}
	}
}
