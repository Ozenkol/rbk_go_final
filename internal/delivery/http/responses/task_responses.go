package http_responses

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
	"github.com/Ozenkol/rbk-go-final/internal/domain/proposal"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
)

// swagger:response getTaskResponse
type getTaskResponse struct {
	// in: body
	Body task.Task
}

// swagger:response createTaskResponse
type createTaskResponse struct {
	// in: body
	Body struct {
		ID string `json:"id"`
	}
}

// swagger:response deleteTaskResponse
type deleteTaskResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response getNoteResponse
type getNoteResponse struct {
	// in: body
	Body note.Note
}

// swagger:response getDocumentResponse
type getDocumentResponse struct {
	// in: body
	Body document.Document
}

// swagger:response getClientResponse
type getClientResponse struct {
	// in: body
	Body client.Client
}

// swagger:response listClientsResponse
type listClientsResponse struct {
	// in: body
	Body []client.Client
}

// swagger:response getAnalyticResponse
type getAnalyticResponse struct {
	// in: body
	Body analytic.Analytic
}

// swagger:response listAnalyticsResponse
type listAnalyticsResponse struct {
	// in: body
	Body []analytic.Analytic
}

// swagger:response getCommunicationResponse
type getCommunicationResponse struct {
	// in: body
	Body communication.Communication
}

// swagger:response listCommunicationsResponse
type listCommunicationsResponse struct {
	// in: body
	Body []communication.Communication
}

// swagger:response getCompanyResponse
type getCompanyResponse struct {
	// in: body
	Body company.Company
}

// swagger:response listCompaniesResponse
type listCompaniesResponse struct {
	// in: body
	Body []company.Company
}

// swagger:response getContractResponse
type getContractResponse struct {
	// in: body
	Body contract.Contract
}

// swagger:response listContractsResponse
type listContractsResponse struct {
	// in: body
	Body []contract.Contract
}

// swagger:response getFileResponse
type getFileResponse struct {
	// in: body
	Body file.File
}

// swagger:response listFilesResponse
type listFilesResponse struct {
	// in: body
	Body []file.File
}

// swagger:response getInvoiceResponse
type getInvoiceResponse struct {
	// in: body
	Body invoice.Invoice
}

// swagger:response listInvoicesResponse
type listInvoicesResponse struct {
	// in: body
	Body []invoice.Invoice
}

// swagger:response getMeetingResponse
type getMeetingResponse struct {
	// in: body
	Body meeting.Meeting
}

// swagger:response listMeetingsResponse
type listMeetingsResponse struct {
	// in: body
	Body []meeting.Meeting
}

// swagger:response getNotificationResponse
type getNotificationResponse struct {
	// in: body
	Body notification.Notification
}

// swagger:response listNotificationsResponse
type listNotificationsResponse struct {
	// in: body
	Body []notification.Notification
}

// swagger:response getProposalResponse
type getProposalResponse struct {
	// in: body
	Body proposal.Proposal
}

// swagger:response getProductResponse
type getProductResponse struct {
	// in: body
	Body product.Product
}

// swagger:response listProductsResponse
type listProductsResponse struct {
	// in: body
	Body []product.Product
}

// swagger:response getSettingResponse
type getSettingResponse struct {
	// in: body
	Body setting.Setting
}

// swagger:response listSettingsResponse
type listSettingsResponse struct {
	// in: body
	Body []setting.Setting
}

// swagger:response getTagResponse
type getTagResponse struct {
	// in: body
	Body tag.Tag
}

// swagger:response listTagsResponse
type listTagsResponse struct {
	// in: body
	Body []tag.Tag
}
