package main

const MAX_LEAVE = 12

type Date struct {
	day int
	month int
	year int
}

type SalaryDetails struct {
	fix int64
	variable int64
}

type ContactDetails struct {
	primaryNum string
	secondaryNum string
	emergencyNum string
}

type CloseFriendsList map[int]string

type Employee struct {
	eId int
	firstName string
	lastName string
	rating float32
	leaveBalance int

	DOB Date
	salary SalaryDetails
	contact ContactDetails
	friendsList CloseFriendsList
}