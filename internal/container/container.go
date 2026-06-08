package container

import (
	"fmt"
	"sync"

	"github.com/Ozenkol/rbk-go-final/internal/application"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/application/service"
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"github.com/Ozenkol/rbk-go-final/internal/infrastructure/persistence"
	"github.com/Ozenkol/rbk-go-final/internal/infrastructure/storage"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Container struct {
	cfg Config

	once sync.Once
	db   *gorm.DB

	muRepos     sync.Mutex
	clientRepo  client.ClientRepositoryInterface
	userRepo    user.UserRepositoryInterface
	offerRepo   offer.OfferRepositoryInterface
	taskRepo    task.TaskRepositoryInterface
	productRepo product.ProductRepositoryInterface
	noteRepo    note.NoteRepositoryInterface
	docRepo     document.DocumentRepositoryInterface
	analyticRepo analytic.AnalyticRepositoryInterface
	commRepo     communication.CommunicationRepositoryInterface
	companyRepo  company.CompanyRepositoryInterface
	contractRepo contract.ContractRepositoryInterface
	meetingRepo  meeting.MeetingRepositoryInterface
	notifRepo    notification.NotificationRepositoryInterface
	fileRepo     file.FileRepositoryInterface
	invoiceRepo  invoice.InvoiceRepositoryInterface
	settingRepo  setting.SettingRepositoryInterface
	tagRepo      tag.TagRepositoryInterface

	storageProvider adapters.ObjectStorageProvider
	app *application.Application
}

type Config struct {
	DSN string
}

func New(cfg Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) DB() *gorm.DB {
	c.once.Do(func() {
		db, err := gorm.Open(postgres.Open(c.cfg.DSN), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("container: failed to open DB: %v", err))
		}
		c.db = db
	})
	return c.db
}

func (c *Container) RedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       0,
	})
}

func (c *Container) ClientRepository() client.ClientRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.clientRepo == nil {
		repo, _ := persistence.NewClientRepository(c.DB())
		c.clientRepo = repo
	}
	return c.clientRepo
}

func (c *Container) UserRepository() user.UserRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.userRepo == nil {
		c.userRepo = persistence.NewUserRepository(c.DB())
	}
	return c.userRepo
}

func (c *Container) OfferRepository() offer.OfferRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.offerRepo == nil {
		repo, _ := persistence.NewOfferRepository(c.DB())
		c.offerRepo = repo
	}
	return c.offerRepo
}

func (c *Container) TaskRepository() task.TaskRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.taskRepo == nil {
		repo, _ := persistence.NewTaskRepository(c.DB())
		c.taskRepo = repo
	}
	return c.taskRepo
}

func (c *Container) ProductRepository() product.ProductRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.productRepo == nil {
		repo, _ := persistence.NewProductRepository(c.DB())
		c.productRepo = repo
	}
	return c.productRepo
}

func (c *Container) NoteRepository() note.NoteRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.noteRepo == nil {
		repo, _ := persistence.NewNoteRepository(c.DB())
		c.noteRepo = repo
	}
	return c.noteRepo
}

func (c *Container) DocumentRepository() document.DocumentRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.docRepo == nil {
		repo, _ := persistence.NewDocumentRepository(c.DB())
		c.docRepo = repo
	}
	return c.docRepo
}

func (c *Container) AnalyticRepository() analytic.AnalyticRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.analyticRepo == nil {
		repo, _ := persistence.NewAnalyticRepository(c.DB())
		c.analyticRepo = repo
	}
	return c.analyticRepo
}

func (c *Container) CommunicationRepository() communication.CommunicationRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.commRepo == nil {
		repo, _ := persistence.NewCommunicationRepository(c.DB())
		c.commRepo = repo
	}
	return c.commRepo
}

func (c *Container) CompanyRepository() company.CompanyRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.companyRepo == nil {
		repo, _ := persistence.NewCompanyRepository(c.DB())
		c.companyRepo = repo
	}
	return c.companyRepo
}

func (c *Container) ContractRepository() contract.ContractRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.contractRepo == nil {
		repo, _ := persistence.NewContractRepository(c.DB())
		c.contractRepo = repo
	}
	return c.contractRepo
}

func (c *Container) MeetingRepository() meeting.MeetingRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.meetingRepo == nil {
		repo, _ := persistence.NewMeetingRepository(c.DB())
		c.meetingRepo = repo
	}
	return c.meetingRepo
}

func (c *Container) NotificationRepository() notification.NotificationRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.notifRepo == nil {
		repo, _ := persistence.NewNotificationRepository(c.DB())
		c.notifRepo = repo
	}
	return c.notifRepo
}

func (c *Container) FileRepository() file.FileRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.fileRepo == nil {
		repo, _ := persistence.NewFileRepository(c.DB())
		c.fileRepo = repo
	}
	return c.fileRepo
}

func (c *Container) InvoiceRepository() invoice.InvoiceRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.invoiceRepo == nil {
		repo, _ := persistence.NewInvoiceRepository(c.DB())
		c.invoiceRepo = repo
	}
	return c.invoiceRepo
}

func (c *Container) SettingRepository() setting.SettingRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.settingRepo == nil {
		repo, _ := persistence.NewSettingRepository(c.DB())
		c.settingRepo = repo
	}
	return c.settingRepo
}

func (c *Container) TagRepository() tag.TagRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.tagRepo == nil {
		repo, _ := persistence.NewTagRepository(c.DB())
		c.tagRepo = repo
	}
	return c.tagRepo
}

func (c *Container) TokenRepository() adapters.TokenRepositoryInterface {
	repo, _ := persistence.NewTokenRepository(c.RedisDB())
	return repo
}

func (c *Container) StorageProvider() adapters.ObjectStorageProvider {
	if c.storageProvider == nil {
		c.storageProvider = storage.NewMinioStorage()
	}
	return c.storageProvider
}

func (c *Container) UserFactory() user.UserFactoryInterface {
	return user.NewUserFactory(c.UserRepository())
}

func (c *Container) App() *application.Application {
	if c.app == nil {
		c.app = &application.Application{
			Commands: application.Commands{
				CreateUser: command.NewCreateUserHandler(c.UserRepository(), c.UserFactory()),
				UpdateUser: command.NewUpdateUserHandler(c.UserRepository()),
				DeleteUser: command.NewDeleteUserHandler(c.UserRepository()),

				CreateClient: command.NewCreateClientHandler(c.ClientRepository()),
				UpdateClient: command.NewUpdateClientHandler(c.ClientRepository()),
				DeleteClient: command.NewDeleteClientHandler(c.ClientRepository()),

				CreateTask: command.NewCreateTaskHandler(c.TaskRepository()),
				UpdateTask: command.NewUpdateTaskHandler(c.TaskRepository()),
				DeleteTask: command.NewDeleteTaskHandler(c.TaskRepository()),

				CreateNote: command.NewCreateNoteHandler(c.NoteRepository()),
				UpdateNote: command.NewUpdateNoteHandler(c.NoteRepository()),
				DeleteNote: command.NewDeleteNoteHandler(c.NoteRepository()),

				CreateOffer: command.NewCreateOfferHandler(c.OfferRepository()),
				UpdateOffer: command.NewUpdateOfferHandler(c.OfferRepository()),
				DeleteOffer: command.NewDeleteOfferHandler(c.OfferRepository()),

				CreateDocument: command.NewCreateDocumentHandler(c.DocumentRepository(), c.StorageProvider()),
				UpdateDocument: command.NewUpdateDocumentHandler(c.DocumentRepository()),
				DeleteDocument: command.NewDeleteDocumentHandler(c.DocumentRepository()),

				CreateAnalytic: command.NewCreateAnalyticHandler(c.AnalyticRepository()),
				UpdateAnalytic: command.NewUpdateAnalyticHandler(c.AnalyticRepository()),
				DeleteAnalytic: command.NewDeleteAnalyticHandler(c.AnalyticRepository()),

				CreateCommunication: command.NewCreateCommunicationHandler(c.CommunicationRepository()),
				UpdateCommunication: command.NewUpdateCommunicationHandler(c.CommunicationRepository()),
				DeleteCommunication: command.NewDeleteCommunicationHandler(c.CommunicationRepository()),

				CreateCompany: command.NewCreateCompanyHandler(c.CompanyRepository()),
				UpdateCompany: command.NewUpdateCompanyHandler(c.CompanyRepository()),
				DeleteCompany: command.NewDeleteCompanyHandler(c.CompanyRepository()),

				CreateContract: command.NewCreateContractHandler(c.ContractRepository()),
				UpdateContract: command.NewUpdateContractHandler(c.ContractRepository()),
				DeleteContract: command.NewDeleteContractHandler(c.ContractRepository()),

				CreateMeeting: command.NewCreateMeetingHandler(c.MeetingRepository()),
				UpdateMeeting: command.NewUpdateMeetingHandler(c.MeetingRepository()),
				DeleteMeeting: command.NewDeleteMeetingHandler(c.MeetingRepository()),

				CreateNotification: command.NewCreateNotificationHandler(c.NotificationRepository()),
				UpdateNotification: command.NewUpdateNotificationHandler(c.NotificationRepository()),
				DeleteNotification: command.NewDeleteNotificationHandler(c.NotificationRepository()),

				CreateFile: command.NewCreateFileHandler(c.FileRepository()),
				UpdateFile: command.NewUpdateFileHandler(c.FileRepository()),
				DeleteFile: command.NewDeleteFileHandler(c.FileRepository()),

				CreateInvoice: command.NewCreateInvoiceHandler(c.InvoiceRepository()),
				UpdateInvoice: command.NewUpdateInvoiceHandler(c.InvoiceRepository()),
				DeleteInvoice: command.NewDeleteInvoiceHandler(c.InvoiceRepository()),

				CreateProduct: command.NewCreateProductHandler(c.ProductRepository()),
				UpdateProduct: command.NewUpdateProductHandler(c.ProductRepository()),
				DeleteProduct: command.NewDeleteProductHandler(c.ProductRepository()),

				CreateSetting: command.NewCreateSettingHandler(c.SettingRepository()),
				UpdateSetting: command.NewUpdateSettingHandler(c.SettingRepository()),
				DeleteSetting: command.NewDeleteSettingHandler(c.SettingRepository()),

				CreateTag: command.NewCreateTagHandler(c.TagRepository()),
				UpdateTag: command.NewUpdateTagHandler(c.TagRepository()),
				DeleteTag: command.NewDeleteTagHandler(c.TagRepository()),
			},
			Queries: application.Queries{
				GetUserByID: query.NewFetchUserByIDHandler(c.UserRepository()),
				ListUsers:   query.NewFetchUserListHandler(c.UserRepository()),

				GetClientByID: query.NewFetchClientByIDHandler(c.ClientRepository()),
				ListClients:   query.NewFetchClientListHandler(c.ClientRepository()),

				GetTaskByID: query.NewFetchTaskByIDHandler(c.TaskRepository()),
				ListTasks:   query.NewFetchTaskListHandler(c.TaskRepository()),

				GetNoteByID: query.NewFetchNoteByIDHandler(c.NoteRepository()),
				ListNotes:   query.NewFetchNoteListHandler(c.NoteRepository()),

				GetOfferByID: query.NewFetchOfferByIDHandler(c.OfferRepository()),
				ListOffers:   query.NewFetchOfferListHandler(c.OfferRepository()),

				GetDocumentByID: query.NewFetchDocumentByIDHandler(c.DocumentRepository()),
				ListDocuments:   query.NewFetchDocumentListHandler(c.DocumentRepository()),

				GetAnalyticByID: query.NewFetchAnalyticByIDHandler(c.AnalyticRepository()),
				ListAnalytics:   query.NewFetchAnalyticListHandler(c.AnalyticRepository()),

				GetCommunicationByID: query.NewFetchCommunicationByIDHandler(c.CommunicationRepository()),
				ListCommunications:   query.NewFetchCommunicationListHandler(c.CommunicationRepository()),

				GetCompanyByID: query.NewFetchCompanyByIDHandler(c.CompanyRepository()),
				ListCompanies:   query.NewFetchCompanyListHandler(c.CompanyRepository()),

				GetContractByID: query.NewFetchContractByIDHandler(c.ContractRepository()),
				ListContracts:   query.NewFetchContractListHandler(c.ContractRepository()),

				GetMeetingByID: query.NewFetchMeetingByIDHandler(c.MeetingRepository()),
				ListMeetings:   query.NewFetchMeetingListHandler(c.MeetingRepository()),

				GetNotificationByID: query.NewFetchNotificationByIDHandler(c.NotificationRepository()),
				ListNotifications:   query.NewFetchNotificationListHandler(c.NotificationRepository()),

				GetFileByID: query.NewFetchFileByIDHandler(c.FileRepository()),
				ListFiles:   query.NewFetchFileListHandler(c.FileRepository()),

				GetInvoiceByID: query.NewFetchInvoiceByIDHandler(c.InvoiceRepository()),
				ListInvoices:   query.NewFetchInvoiceListHandler(c.InvoiceRepository()),

				GetProductByID: query.NewFetchProductByIDHandler(c.ProductRepository()),
				ListProducts:   query.NewFetchProductListHandler(c.ProductRepository()),

				GetSettingByID: query.NewFetchSettingByIDHandler(c.SettingRepository()),
				ListSettings:   query.NewFetchSettingListHandler(c.SettingRepository()),

				GetTagByID: query.NewFetchTagByIDHandler(c.TagRepository()),
				ListTags:   query.NewFetchTagListHandler(c.TagRepository()),
			},
			Services: application.Services{
				AuthService: service.NewAuthService(c.UserRepository(), c.TokenRepository()),
			},
		}
	}
	return c.app
}
