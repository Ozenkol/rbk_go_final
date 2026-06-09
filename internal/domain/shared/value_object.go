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

func (*Money) NewMoney(amount float64, currency Currency) Money {
	if amount < 0 {
		panic("Amount cannot be negative")
	}
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (*HumanName) NewHumanName(firstName, middleName, lastName string) HumanName {
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

func (*BirthDate) NewBirthDate(day, month, year int) BirthDate {
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

func (*Address) NewAddress(country, city, street, building, apartment string) Address {
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

func (*PhoneNumber) NewPhoneNumber(countryCode, number string) PhoneNumber {
	if countryCode == "" || number == "" {
		panic("Country code and number cannot be empty")
	}
	return PhoneNumber{
		CountryCode: countryCode,
		Number:      number,
	}
}

func (*Person) NewPerson(name HumanName, birthDate BirthDate, address Address, phone PhoneNumber) Person {
	return Person{
		Name:      name,
		BirthDate: birthDate,
		Address:   address,
		Phone:     phone,
	}
}

type ClientType string

const (
	ClientTypePerson  ClientType = "PERSON"
	ClientTypeCompany ClientType = "COMPANY"
)

type DealStage string

const (
	DealStageNew           DealStage = "NEW"
	DealStageContacted     DealStage = "CONTACTED"
	DealStageNeedsAnalysis DealStage = "NEEDS_ANALYSIS"
	DealStageProposalSent  DealStage = "PROPOSAL_SENT"
	DealStageNegotiation   DealStage = "NEGOTIATION"
	DealStageContract      DealStage = "CONTRACT"
	DealStageWon           DealStage = "WON"
	DealStageLost          DealStage = "LOST"
)

type TaskStatus string

const (
	TaskStatusNew        TaskStatus = "NEW"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
	TaskStatusCancelled  TaskStatus = "CANCELLED"
	TaskStatusOverdue    TaskStatus = "OVERDUE"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "LOW"
	TaskPriorityMedium TaskPriority = "MEDIUM"
	TaskPriorityHigh   TaskPriority = "HIGH"
)

type CommunicationType string

const (
	CommTypeCall     CommunicationType = "CALL"
	CommTypeEmail    CommunicationType = "EMAIL"
	CommTypeWhatsApp CommunicationType = "WHATSAPP"
	CommTypeMeeting  CommunicationType = "MEETING"
	CommTypeSMS      CommunicationType = "SMS"
	CommTypeOther    CommunicationType = "OTHER"
)

type ProposalStatus string

const (
	ProposalStatusDraft    ProposalStatus = "DRAFT"
	ProposalStatusSent     ProposalStatus = "SENT"
	ProposalStatusViewed   ProposalStatus = "VIEWED"
	ProposalStatusAccepted ProposalStatus = "ACCEPTED"
	ProposalStatusRejected ProposalStatus = "REJECTED"
	ProposalStatusExpired  ProposalStatus = "EXPIRED"
)

type InvoiceStatus string

const (
	InvoiceStatusDraft         InvoiceStatus = "DRAFT"
	InvoiceStatusSent          InvoiceStatus = "SENT"
	InvoiceStatusPaid          InvoiceStatus = "PAID"
	InvoiceStatusPartiallyPaid InvoiceStatus = "PARTIALLY_PAID"
	InvoiceStatusOverdue       InvoiceStatus = "OVERDUE"
	InvoiceStatusCancelled     InvoiceStatus = "CANCELLED"
)

type ContractStatus string

const (
	ContractStatusDraft      ContractStatus = "DRAFT"
	ContractStatusOnApproval ContractStatus = "ON_APPROVAL"
	ContractStatusSigned     ContractStatus = "SIGNED"
	ContractStatusActive     ContractStatus = "ACTIVE"
	ContractStatusExpired    ContractStatus = "EXPIRED"
	ContractStatusTerminated ContractStatus = "TERMINATED"
)

type UserRole string

const (
	RoleAdmin        UserRole = "ADMIN"
	RoleSalesLead    UserRole = "SALES_LEAD"
	RoleSalesManager UserRole = "SALES_MANAGER"
	RoleAccountant   UserRole = "ACCOUNTANT"
	RoleSupport      UserRole = "SUPPORT"
	RoleViewer       UserRole = "VIEWER"
)
