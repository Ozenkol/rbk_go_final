package shared

type HumanName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Address struct {
	Country   string
	City      string
	Street    string
	Building  string
	Apartment string
}

type PhoneNumber struct {
	CountryCode string
	Number      string
}

type Person struct {
	Name      HumanName
	BirthDate BirthDate
	Address   Address
	Phone     PhoneNumber
}

type StorageReference struct {
	ServiceName string
	URL         string
}

type Currency struct {
	Code string
	Name string
}

type Money struct {
	Amount   float64
	Currency Currency
}

func NewMoney(amount float64, currency Currency) Money {
	if amount < 0 {
		panic("Amount cannot be negative")
	}
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func NewHumanName(firstName, middleName, lastName string) HumanName {
	if firstName == "" || lastName == "" {
		panic("First name and last name cannot be empty")
	}
	if len(firstName) > 50 || len(middleName) > 50 || len(lastName) > 50 {
		panic("First name, middle name and last name cannot be longer than 50 characters")
	}
	return HumanName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}
}

func NewBirthDate(day, month, year int) BirthDate {
	if day < 1 || day > 31 {
		panic("Day must be between 1 and 31")
	}
	if month < 1 || month > 12 {
		panic("Month must be between 1 and 12")
	}
	if year < 0 {
		panic("Year must be a positive integer")
	}
	return BirthDate{
		Day:   day,
		Month: month,
		Year:  year,
	}
}

func NewAddress(country, city, street, building, apartment string) Address {
	if country == "" || city == "" || street == "" || building == "" {
		panic("Country, city, street and building cannot be empty")
	}
	if len(apartment) > 50 {
		panic("Apartment cannot be longer than 50 characters")
	}
	return Address{
		Country:   country,
		City:      city,
		Street:    street,
		Building:  building,
		Apartment: apartment,
	}
}

func NewPhoneNumber(countryCode, number string) PhoneNumber {
	if countryCode == "" || number == "" {
		panic("Country code and number cannot be empty")
	}
	return PhoneNumber{
		CountryCode: countryCode,
		Number:      number,
	}
}

func NewPerson(name HumanName, birthDate BirthDate, address Address, phone PhoneNumber) Person {
	return Person{
		Name:      name,
		BirthDate: birthDate,
		Address:   address,
		Phone:     phone,
	}
}
