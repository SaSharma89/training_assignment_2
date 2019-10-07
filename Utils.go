package main

import "fmt"

func (e *Employee) AddBasicDetails(id int, first, last string, leave int) {
	e.eId = id
	e.firstName = first
	e.lastName = last
	e.leaveBalance = leave
	e.friendsList = make(map[int]string)
}

func (e * Employee) CheckName( first , last string)( bool){
	if first == e.firstName && last == e.lastName{
		return true
	} else {
		return false
	}
}

func (e *Employee) AddRating( rating float32){
	e.rating = rating
}

func (e *Employee) AddSalaryDetails( s SalaryDetails){
	e.salary = s
}

func (e *Employee) AddContactDetails( c ContactDetails){
	e.contact.emergencyNum = c.emergencyNum
	e.contact.secondaryNum = c.secondaryNum
	e.contact.primaryNum = c.primaryNum
}

func (e *Employee) AddDOB(d Date){
	e.DOB = d
}

func (e *Employee)AddCloseFriend(eid int, name string){
	e.friendsList[eid] = name
}

func (e *Employee)RemoveCloseFriend(eid int){
	delete(e.friendsList, eid)
}

func (e *Employee)DisplayCloseFriend(){
	fmt.Println(e.friendsList)
}

func (e Employee)DisplayDetails(){
	fmt.Println("Employee Id : ", e.eId)
	fmt.Println( "Name : ", e.firstName, " ", e.lastName)
	fmt.Println("Rating : ", e.rating)
	fmt.Println("Leave Balance :", e.leaveBalance)
	fmt.Println("DOB : ", e.DOB.day, "/", e.DOB.month, "/", e.DOB.year)
	fmt.Println("Salary details ::: Fix :", e.salary.fix, " Variable :", e.salary.variable)
	fmt.Println("Contact details ::: Primary : ", e.contact.primaryNum ,
		" Secondary : ", e.contact.secondaryNum, " Emergency : ", e.contact.emergencyNum)
	fmt.Println("Close friends list :", e.friendsList)
}

func (e Employee)GetEmergencyContact(){
	fmt.Println("Emergency contact details is ", e.contact.emergencyNum )
}

func (e *Employee)LeaveTaken(){
	e.leaveBalance--
}

func (e Employee)lossOfPay()(loss float32){
	perDayIncome := float32(e.salary.fix * 100000/30)
	if e.leaveBalance < MAX_LEAVE {
		leaveTaken := float32(MAX_LEAVE - e.leaveBalance)
		loss = leaveTaken * perDayIncome
	}
	return
}

func (e Employee)CalculateNetSalary()(salary float32){
	fixPay := float32(e.salary.fix * 100000) - e.lossOfPay()
	salary = fixPay + float32(e.salary.variable * 100000)
	fmt.Println("Total earned income is ", fixPay + float32(e.salary.variable * 100000), " and loss of pay ", e.lossOfPay())
	return
}

func (e Employee)CalculateSalaryHike(){
	hike := (e.rating/100) * e.CalculateNetSalary()
	fmt.Println("Employee hike is ", hike)
}