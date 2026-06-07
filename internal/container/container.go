package container

import (
	"fmt"
	"sync"

	"github.com/Ozenkol/rbk-go-final/internal/application"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/application/service"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"github.com/Ozenkol/rbk-go-final/internal/infrastructure/persistence"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Container is the single source of truth for dependency wiring.
// All fields are lazily initialized and cached (singleton per process lifetime).
type Container struct {
	cfg Config

	// infrastructure
	once sync.Once
	db   *gorm.DB

	// repositories — lazy singletons
	muRepos     sync.Mutex
	clientRepo  client.ClientRepositoryInterface
	userRepo    user.UserRepositoryInterface
	offerRepo   offer.OfferRepositoryInterface
	taskRepo    task.TaskRepositoryInterface
	productRepo product.ProductRepositoryInterface

	// domain factories
	offerFactory offer.OfferFactoryInterface
	userFactory  user.UserFactoryInterface

	// application layer
	app *application.Application
}

type Config struct {
	DSN string // e.g. "test.db" or postgres DSN
}

func New(cfg Config) *Container {
	return &Container{cfg: cfg}
}

// --- Infrastructure ---

// DB returns a cached *gorm.DB. Panics only at startup, which is acceptable
// because the app literally cannot run without a DB — fail fast is correct here.
// If you disagree: wrap in an Init() error and call it explicitly in main.
func (c *Container) DB() *gorm.DB {
	c.once.Do(func() {
		db, err := gorm.Open(postgres.Open(c.cfg.DSN), &gorm.Config{})
		if err != nil {
			// panic here is intentional: no DB = no app.
			// In production swap sqlite for postgres and load DSN from env.
			panic(fmt.Sprintf("container: failed to open DB: %v", err))
		}
		// Run automigrate or migrations here so it's in one place.
		c.db = db
	})
	return c.db
}

func (c *Container) RedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // Set if your Redis instance requires authentication
		DB:       0,          // Use default DB
	})
}
// --- Repositories ---

func (c *Container) ClientRepository() client.ClientRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.clientRepo == nil {
		repo, err := persistence.NewClientRepository(c.DB())
		if err != nil {
			panic(fmt.Sprintf("container: ClientRepository init failed: %v", err))
		}
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

func (c *Container) ProductRepository() product.ProductRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.productRepo == nil {
		repo, err := persistence.NewProductRepository(c.DB())
		if err != nil {
			panic(fmt.Sprintf("container: ProductRepository init failed: %v", err))
		}
		c.productRepo = repo
	}
	return c.productRepo
}

func (c *Container) OfferRepository() offer.OfferRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.offerRepo == nil {
		repo, err := persistence.NewOfferRepository(c.DB())
		if err != nil {
			panic(fmt.Sprintf("container: OfferRepository init failed: %v", err))
		}
		c.offerRepo = repo
	}
	return c.offerRepo
}

func (c *Container) TaskRepository() task.TaskRepositoryInterface {
	c.muRepos.Lock()
	defer c.muRepos.Unlock()
	if c.taskRepo == nil {
		repo, err := persistence.NewTaskRepository(c.DB())
		if err != nil {
			panic(fmt.Sprintf("container: TaskRepository init failed: %v", err))
		}
		c.taskRepo = repo
	}
	return c.taskRepo
}

func (c *Container) TokenRepository() adapters.TokenRepositoryInterface {
	repo, err := persistence.NewTokenRepository(c.RedisDB())
	if err != nil {
		panic(fmt.Sprintf("container: TokenRepository init failed: %v", err))
	}
	return repo
}

func (c *Container) InvoiceRepository() invoice.InvoiceRepositoryInterface {
	repo, err := persistence.NewInvoiceRepository(c.DB())
	if err != nil {
		panic(fmt.Sprintf("container: InvoiceRepository init failed: %v", err))
	}
	return repo
}

func (c *Container) FileRepository() file.FileRepositoryInterface {
	repo, err := persistence.NewFileRepository(c.DB())
	if err != nil {
		panic(fmt.Sprintf("container: FileRepository init failed: %v", err))
	}
	return repo
}

func (c *Container) SettingRepository() setting.SettingRepositoryInterface {
	repo, err := persistence.NewSettingRepository(c.DB())
	if err != nil {
		panic(fmt.Sprintf("container: SettingRepository init failed: %v", err))
	}
	return repo
}

func (c *Container) TagRepository() tag.TagRepositoryInterface {
	repo, err := persistence.NewTagRepository(c.DB())
	if err != nil {
		panic(fmt.Sprintf("container: TagRepository init failed: %v", err))
	}
	return repo
}

// --- Domain ---

func (c *Container) OfferFactory() offer.OfferFactoryInterface {
	if c.offerFactory == nil {
		c.offerFactory = offer.NewOfferFactory(c.ProductRepository())
	}
	return c.offerFactory
}

func (c *Container) UserFactory() user.UserFactoryInterface {
	if c.userFactory == nil {
		c.userFactory = user.NewUserFactory(
			c.UserRepository(),
		)
	}
	return c.userFactory
}


// --- Application ---

// App assembles the CQRS Application struct once.
// Add handlers here as your command/query surface grows.
func (c *Container) App() *application.Application {
	if c.app == nil {
		c.app = &application.Application{
			Commands: application.Commands{
				// Clients command handler
				CreateUser: command.NewCreateUserHandler(
					c.UserRepository(),
					c.UserFactory(),
				),
				CreateClient: command.NewCreateClientHandler(
					c.ClientRepository(),
				),

				// Tasks command handlers
				CreateTask: command.NewCreateTaskHandler(
					c.TaskRepository(),
				),
				DeleteTask: command.NewDeleteTaskHandler(
					c.TaskRepository(),
				),
				UpdateTask: command.NewUpdateTaskHandler(
					c.TaskRepository(),
				),

			},
			Queries: application.Queries{
				// Clients query handler
				GetUserByID: query.NewFetchUserHandler(c.UserRepository()),

				// Tasks query handlers
				GetTaskByID: query.NewFetchTaskByIDHandler(c.TaskRepository()),
			},
			 Services: application.Services{
				// Auth service
				AuthService: service.NewAuthService(c.UserRepository(), c.TokenRepository()),
			 },			
		}
	}
	return c.app
}
