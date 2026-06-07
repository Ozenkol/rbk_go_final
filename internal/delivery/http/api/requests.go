package api_requests

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
)

// swagger:parameters createUser
type createUserParams struct {
    // in: body
    // required: true
    Body struct {
        FirstName string `json:"first_name"`
		MiddleName string `json:"middle_name,omitempty"`
        LastName  string `json:"last_name"`
        Email     string `json:"email"`
        Password  string `json:"password"`
    }
}

// swagger:parameters registerUser
type registerUserParams struct {
	// in: body
	// required: true
	Body struct {
		// required: true
		// default: John
		FirstName  string `json:"first_name"`
		// required: true
		// default: Smith
		MiddleName string `json:"middle_name,omitempty"`

		// required: true
		// default: Adams
		LastName   string `json:"last_name"`

		// required: true
		// default: test@example.com
		Email      string `json:"email"`

		// required: true
		// default: 123456
		Password   string `json:"password"`
	}
}

// swagger:parameters loginUser
type loginUserParams struct {
	// in: body
	// required: true
	Body struct {
		// required: true
		// default: test@example.com
		Email string `json:"email"`

		// required: true
		// default: 123456
		Password string `json:"password"`
	}
}

// swagger:parameters createTask
type createTaskParams struct {
	// in: body
	// required: true
	Body task.Task
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
	Body task.Task
}

// swagger:parameters createNote
type createNoteParams struct {
	// in: body
	// required: true
	Body note.Note
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
	Body document.Document
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
	Body client.Client
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
	Body analytic.Analytic
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
	Body communication.Communication
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
	Body company.Company
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
	Body contract.Contract
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
	Body file.File
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
	Body invoice.Invoice
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
	Body meeting.Meeting
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
	Body notification.Notification
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
	Body offer.Offer
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
	Body product.Product
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
	Body setting.Setting
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
	Body tag.Tag
}

// swagger:parameters getTag updateTag deleteTag
type tagIDParams struct {
	// in: path
	// required: true
	ID string `json:"id"`
}
