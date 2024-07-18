package handlers

import (
	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/mocks"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	dbrepo "github.com/Orololuwa/collect_am-api/src/repository/db-repo"
	"github.com/Orololuwa/collect_am-api/src/types"
)

type ErrorData struct {
	Message string
	Error   error
	Status  int
}

type HandlerFunc interface {
	// Auth
	SignUp(payload dtos.UserSignUp) (userId uint, errData *ErrorData)
	LoginUser(payload dtos.UserLoginBody) (data types.LoginSuccessResponse, errData *ErrorData)

	// Business
	CreateBusiness(payload dtos.AddBusiness, options ...*Extras) (id uint, errData *ErrorData)
	GetBusiness(id uint, options ...*Extras) (data *models.Business, errData *ErrorData)
	UpdateBusiness(id uint, payload map[string]interface{}, options ...*Extras) (errData *ErrorData)

	// Products
	AddProduct(payload dtos.AddProduct, options ...*Extras) (id uint, errData *ErrorData)
	UpdateProduct(payload dtos.UpdateProduct, options ...*Extras) (errData *ErrorData)
	GetAllProducts(query map[string]interface{}, options ...*Extras) (products []models.Product, pagination repository.Pagination, errData *ErrorData)
	GetProduct(id uint, options ...*Extras) (product models.Product, errData *ErrorData)

	// Customers
	AddCustomer(payload dtos.CreateCustomer, options ...*Extras) (id uint, errData *ErrorData)
	EditCustomer(payload types.EditCustomerPayload, options ...*Extras) (errData *ErrorData)
	GetCustomer(payload types.GetACustomerPayload, options ...*Extras) (customer models.Customer, errData *ErrorData)
	GetAllCustomers(query map[string]interface{}, options ...*Extras) (customers []models.Customer, pagination repository.Pagination, errData *ErrorData)

	// Invoice
	CreateInvoice(payload types.CreateInvoicePayload, options ...*Extras) (id uint, errData *ErrorData)
	GetInvoice(payload types.GetAnInvoicePayload, options ...*Extras) (customer models.Invoice, errData *ErrorData)
	GetAllInvoices(query map[string]interface{}, options ...*Extras) (customers []models.Invoice, pagination repository.Pagination, errData *ErrorData)
}

type Repository struct {
	App           *config.AppConfig
	conn          repository.DBInterface
	User          repository.UserDBRepo
	Business      repository.BusinessDBRepo
	Kyc           repository.KycDBRepo
	Product       repository.ProductDBRepo
	Customer      repository.CustomerDBRepo
	Address       repository.AddressDBRepo
	Invoice       repository.InvoiceDBRepo
	ListedProduct repository.ListedProductDBRepo
}

// NewHandlers function initializes the Repo
func NewHandlers(a *config.AppConfig, db *driver.DB) HandlerFunc {
	return &Repository{
		App:           a,
		conn:          db.Gorm,
		User:          dbrepo.NewUserDBRepo(db),
		Business:      dbrepo.NewBusinessDBRepo(db),
		Kyc:           dbrepo.NewKycDBRepo(db),
		Product:       dbrepo.NewProductDBRepo(db),
		Customer:      dbrepo.NewCustomerDBRepo(db),
		Address:       dbrepo.NewAddressDBRepo(db),
		Invoice:       dbrepo.NewInvoiceDBRepo(db),
		ListedProduct: dbrepo.NewListedProductDBRepo(db),
	}
}

// NewHandlers function initializes the Repo
func NewTestHandlers(a *config.AppConfig) HandlerFunc {
	mockDB := mocks.NewMockDB()

	return &Repository{
		App:           a,
		conn:          mockDB,
		User:          dbrepo.NewUserTestingDBRepo(),
		Business:      dbrepo.NewBusinessTestingDBRepo(),
		Kyc:           dbrepo.NewKycTestingDBRepo(),
		Product:       dbrepo.NewProductTestingDBRepo(),
		Customer:      dbrepo.NewCustomerTestingDBRepo(),
		Address:       dbrepo.NewAddressTestingDBRepo(),
		Invoice:       dbrepo.NewInvoiceTestingDBRepo(),
		ListedProduct: dbrepo.NewListedProductTestingDBRepo(),
	}
}

type Extras struct {
	User     *models.User
	Business *models.Business
}
