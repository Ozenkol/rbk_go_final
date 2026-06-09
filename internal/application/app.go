package application

import (
	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/application/service"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Services Services
}

type Commands struct {
	CreateUser *command.CreateUserHandler
	UpdateUser *command.UpdateUserHandler
	DeleteUser *command.DeleteUserHandler

	CreateClient *command.CreateClientHandler
	UpdateClient *command.UpdateClientHandler
	DeleteClient *command.DeleteClientHandler

	CreateDeal *command.CreateDealHandler
	UpdateDeal *command.UpdateDealHandler
	DeleteDeal *command.DeleteDealHandler

	CreateTask *command.CreateTaskHandler
	UpdateTask *command.UpdateTaskHandler
	DeleteTask *command.DeleteTaskHandler

	CreateNote *command.CreateNoteHandler
	UpdateNote *command.UpdateNoteHandler
	DeleteNote *command.DeleteNoteHandler

	CreateProposal *command.CreateProposalHandler
	UpdateProposal *command.UpdateProposalHandler
	DeleteProposal *command.DeleteProposalHandler

	CreateDocument *command.CreateDocumentHandler
	UpdateDocument *command.UpdateDocumentHandler
	DeleteDocument *command.DeleteDocumentHandler

	CreateAnalytic *command.CreateAnalyticHandler
	UpdateAnalytic *command.UpdateAnalyticHandler
	DeleteAnalytic *command.DeleteAnalyticHandler

	CreateCommunication *command.CreateCommunicationHandler
	UpdateCommunication *command.UpdateCommunicationHandler
	DeleteCommunication *command.DeleteCommunicationHandler

	CreateCompany *command.CreateCompanyHandler
	UpdateCompany *command.UpdateCompanyHandler
	DeleteCompany *command.DeleteCompanyHandler

	CreateContract *command.CreateContractHandler
	UpdateContract *command.UpdateContractHandler
	DeleteContract *command.DeleteContractHandler

	CreateMeeting *command.CreateMeetingHandler
	UpdateMeeting *command.UpdateMeetingHandler
	DeleteMeeting *command.DeleteMeetingHandler

	CreateNotification *command.CreateNotificationHandler
	UpdateNotification *command.UpdateNotificationHandler
	DeleteNotification *command.DeleteNotificationHandler

	CreateFile *command.CreateFileHandler
	UpdateFile *command.UpdateFileHandler
	DeleteFile *command.DeleteFileHandler

	CreateInvoice *command.CreateInvoiceHandler
	UpdateInvoice *command.UpdateInvoiceHandler
	DeleteInvoice *command.DeleteInvoiceHandler

	CreateProduct *command.CreateProductHandler
	UpdateProduct *command.UpdateProductHandler
	DeleteProduct *command.DeleteProductHandler

	CreateSetting *command.CreateSettingHandler
	UpdateSetting *command.UpdateSettingHandler
	DeleteSetting *command.DeleteSettingHandler

	CreateTag *command.CreateTagHandler
	UpdateTag *command.UpdateTagHandler
	DeleteTag *command.DeleteTagHandler
}

type Queries struct {
	GetUserByID *query.FetchUserByIDHandler
	ListUsers   *query.FetchUserListHandler

	GetClientByID *query.FetchClientByIDHandler
	ListClients   *query.FetchClientListHandler

	GetDealByID *query.FetchDealByIDHandler
	ListDeals   *query.FetchDealListHandler

	GetTaskByID *query.FetchTaskByIDHandler
	ListTasks   *query.FetchTaskListHandler

	GetNoteByID *query.FetchNoteByIDHandler
	ListNotes   *query.FetchNoteListHandler

	GetProposalByID *query.FetchProposalByIDHandler
	ListProposals   *query.FetchProposalListHandler

	GetDocumentByID *query.FetchDocumentByIDHandler
	ListDocuments   *query.FetchDocumentListHandler

	GetAnalyticByID *query.FetchAnalyticByIDHandler
	ListAnalytics   *query.FetchAnalyticListHandler

	GetCommunicationByID *query.FetchCommunicationByIDHandler
	ListCommunications   *query.FetchCommunicationListHandler

	GetCompanyByID *query.FetchCompanyByIDHandler
	ListCompanies   *query.FetchCompanyListHandler

	GetContractByID *query.FetchContractByIDHandler
	ListContracts   *query.FetchContractListHandler

	GetMeetingByID *query.FetchMeetingByIDHandler
	ListMeetings   *query.FetchMeetingListHandler

	GetNotificationByID *query.FetchNotificationByIDHandler
	ListNotifications   *query.FetchNotificationListHandler

	GetFileByID *query.FetchFileByIDHandler
	ListFiles   *query.FetchFileListHandler

	GetInvoiceByID *query.FetchInvoiceByIDHandler
	ListInvoices   *query.FetchInvoiceListHandler

	GetProductByID *query.FetchProductByIDHandler
	ListProducts   *query.FetchProductListHandler

	GetSettingByID *query.FetchSettingByIDHandler
	ListSettings   *query.FetchSettingListHandler

	GetTagByID *query.FetchTagByIDHandler
	ListTags   *query.FetchTagListHandler
}

type Services struct {
	AuthService *service.AuthService
}
