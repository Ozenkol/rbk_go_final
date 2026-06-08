package api_requests

import (
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
)

// swagger:parameters createUser
type createUserParams struct {
	// in: body
	// required: true
	Body http_requests.CreateUserRequest
}

// swagger:parameters registerUser
type registerUserParams struct {
	// in: body
	// required: true
	Body http_requests.CreateUserRequest
}

// swagger:parameters loginUser
type loginUserParams struct {
	// in: body
	// required: true
	Body struct {
		// required: true
		// example: test@example.com
		Email string `json:"email"`

		// required: true
		// example: 123456
		Password string `json:"password"`
	}
}

// swagger:parameters createTask
type createTaskParams struct {
	// in: body
	// required: true
	Body http_requests.CreateTaskRequest
}

// swagger:parameters getTask deleteTask toggleTask
type taskIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters updateTask
type updateTaskParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
	// in: body
	// required: true
	Body http_requests.UpdateTaskRequest
}

// swagger:parameters createNote
type createNoteParams struct {
	// in: body
	// required: true
	Body http_requests.CreateNoteRequest
}

// swagger:parameters getNote updateNote deleteNote
type noteIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createDocument
type createDocumentParams struct {
	// in: body
	// required: true
	Body http_requests.CreateDocumentRequest
}

// swagger:parameters getDocument updateDocument deleteDocument
type documentIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createClient
type createClientParams struct {
	// in: body
	// required: true
	Body http_requests.CreateClientRequest
}

// swagger:parameters getClient updateClient deleteClient
type clientIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createAnalytic
type createAnalyticParams struct {
	// in: body
	// required: true
	Body http_requests.CreateAnalyticRequest
}

// swagger:parameters getAnalytic updateAnalytic deleteAnalytic
type analyticIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createCommunication
type createCommunicationParams struct {
	// in: body
	// required: true
	Body http_requests.CreateCommunicationRequest
}

// swagger:parameters getCommunication updateCommunication deleteCommunication
type communicationIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createCompany
type createCompanyParams struct {
	// in: body
	// required: true
	Body http_requests.CreateCompanyRequest
}

// swagger:parameters getCompany updateCompany deleteCompany
type companyIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createContract
type createContractParams struct {
	// in: body
	// required: true
	Body http_requests.CreateContractRequest
}

// swagger:parameters getContract updateContract deleteContract
type contractIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createFile
type createFileParams struct {
	// in: body
	// required: true
	Body http_requests.CreateFileRequest
}

// swagger:parameters getFile updateFile deleteFile
type fileIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createInvoice
type createInvoiceParams struct {
	// in: body
	// required: true
	Body http_requests.CreateInvoiceRequest
}

// swagger:parameters getInvoice updateInvoice deleteInvoice
type invoiceIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createMeeting
type createMeetingParams struct {
	// in: body
	// required: true
	Body http_requests.CreateMeetingRequest
}

// swagger:parameters getMeeting updateMeeting deleteMeeting
type meetingIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createNotification
type createNotificationParams struct {
	// in: body
	// required: true
	Body http_requests.CreateNotificationRequest
}

// swagger:parameters getNotification updateNotification deleteNotification
type notificationIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createOffer
type createOfferParams struct {
	// in: body
	// required: true
	Body http_requests.CreateOfferRequest
}

// swagger:parameters getOffer updateOffer deleteOffer
type offerIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createProduct
type createProductParams struct {
	// in: body
	// required: true
	Body http_requests.CreateProductRequest
}

// swagger:parameters getProduct updateProduct deleteProduct
type productIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createSetting
type createSettingParams struct {
	// in: body
	// required: true
	Body http_requests.CreateSettingRequest
}

// swagger:parameters getSetting updateSetting deleteSetting
type settingIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters createTag
type createTagParams struct {
	// in: body
	// required: true
	Body http_requests.CreateTagRequest
}

// swagger:parameters getTag updateTag deleteTag
type tagIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}
