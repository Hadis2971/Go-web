package dataAccess

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Hadis2971/go_web/models"
)

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type FoundUserReponse struct {
	ID models.UserId `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Password string
}



type IUserDataAccess interface {
	CreateUser(user models.User) error
	DeleteUser(id int) error
	UpdateUser(updateUserRequest UpdateUserRequest) error
	GetUserByUsernameOrEmail(user models.User) (*FoundUserReponse, error)
	GetUserById(id int) (*FoundUserReponse, error)
	GetAllUsersAndTheirOrders() ([]models.UserWithOrders, error)
}

type UserDataAccess struct {
	dbConnection *sql.DB
}

var ( 
	InternalServerError = errors.New("Internal Server Error!!!")
	ErrorMissingID = errors.New("Missing ID Field!!!")
	ErrorMissingUsernameOrEmail = errors.New("Missing Username Or Email!!!")
	ErrorUserNotFound = errors.New("User Not Found!!!")
	ErrorNoUserOrderProductsFound = errors.New("No Users And Product Orders Found!!!")
)

func NewUserDataAccess(dbConnection *sql.DB) *UserDataAccess {
	return &UserDataAccess{dbConnection: dbConnection}
}

func (da UserDataAccess) CreateUser(user models.User) error {
	query := "INSERT INTO User (username, email, password) VALUES (?, ?, ?)"

	_, err := da.dbConnection.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return InternalServerError
	}


	return nil
}

func (da UserDataAccess) DeleteUser(id int) error {
	query := "DELETE FROM User WHERE id = ?"

	if (id == 0) {
		return ErrorMissingID
	}

	_, err := da.dbConnection.Exec(query, id)

	if err != nil {
		return InternalServerError
	}

	return nil
}

func (da UserDataAccess) UpdateUser(updateUserRequest UpdateUserRequest) error {
	query := "UPDATE User SET username = ?, email = ? WHERE id = ?"

	if (updateUserRequest.ID == 0) {
		return ErrorMissingID
	}

	_, err := da.dbConnection.Query(query, updateUserRequest.Username, updateUserRequest.Email, updateUserRequest.ID)

	if err != nil {
		return InternalServerError
	}

	return nil
}

func (da UserDataAccess) GetUserByUsernameOrEmail(user models.User) (*FoundUserReponse, error) {
	query := "SELECT * FROM User WHERE username = ? OR email = ?"
	
	var foundUserReponse FoundUserReponse

	if (user.Username == "" && user.Email == "") {
		return nil, ErrorMissingUsernameOrEmail
	}

	row := da.dbConnection.QueryRow(query, user.Username, user.Email)

	err := row.Scan(&foundUserReponse.ID, &foundUserReponse.Username, &foundUserReponse.Email, &foundUserReponse.Password, &foundUserReponse.CreatedOn, &foundUserReponse.UpdatedOn)

	if err != nil {
		return nil, err
	}

	return &foundUserReponse, nil
}

func (da UserDataAccess) GetUserById(id int) (*FoundUserReponse, error) {
	query := "SELECT * FROM User WHERE id = ?"
	var foundUserReponse FoundUserReponse

	if (id == 0) {
		return nil, ErrorMissingID
	}

	row := da.dbConnection.QueryRow(query, id)

	err := row.Scan(&foundUserReponse.ID, &foundUserReponse.Username, &foundUserReponse.Email, &foundUserReponse.Password, &foundUserReponse.CreatedOn, &foundUserReponse.UpdatedOn)


	
	if err != nil {
		return nil, ErrorUserNotFound
	}

	return &foundUserReponse, nil
}

func (da UserDataAccess) GetAllUsersAndTheirOrders() ([]models.UserWithOrders, error) {
	query := `SELECT User.id, User.username, User.email, 
	Product_Order.id as product_order_id, Product_Order.user_id, Product_Order.order_id, Product_Order.product_id, Product_Order.quantity, 
	Product.id as product_id, Product.name as product_name, Product.description AS product_description, Product.price, Product.stock, Product.created_on, Product.updated_on
	FROM User LEFT JOIN Product_Order ON User.id = Product_Order.user_id 
	LEFT JOIN Product ON Product_Order.product_id = Product.id;`

	var basicUser models.BasicUser
	var fullOrder models.FullOrder
	var fullOrders []models.FullOrder
	var product models.Product
	var productOrderWithProduct models.ProductOrderWithProduct
	var userWithOrders []models.UserWithOrders

	userAndOrdersMap := make(map[int]models.UserWithOrders)

	rows, err := da.dbConnection.Query(query)

	if err == sql.ErrNoRows {
		return nil, ErrorNoUserOrderProductsFound
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&fullOrder.Id, &fullOrder.Username, &fullOrder.Email, &fullOrder.ProductOrderId, &fullOrder.UserId, &fullOrder.OrderId, &fullOrder.ProductId, &fullOrder.Quantity, &fullOrder.ProductId, &fullOrder.ProductName, &fullOrder.ProductDescription, &fullOrder.ProductPrice, &fullOrder.ProductStock, &fullOrder.ProductCreatedOn, &fullOrder.ProductUpdatedOn)

		fullOrders = append(fullOrders, fullOrder)
	}

	for _, val := range fullOrders {
		_, ok := userAndOrdersMap[val.Id]

		if !ok {
			basicUser = models.BasicUser{ID: val.Id, Username: val.Username, Email: val.Email}
			userAndOrdersMap[val.Id] = models.UserWithOrders{ID: basicUser.ID, Username: basicUser.Username, Email: basicUser.Email}
		}

		if val.OrderId != nil {

			product = models.Product{ID: *val.ProductId, Name: *val.ProductName, Description: *val.ProductDescription, Price: *val.ProductPrice, Stock: *val.ProductStock, CreatedOn: *val.ProductCreatedOn, UpdatedOn: *val.ProductCreatedOn}
			productOrderWithProduct = models.ProductOrderWithProduct{ProductOrderId: *val.ProductOrderId, UserId: *val.UserId, OrderId: *val.OrderId, ProductId: *val.ProductId, Quantity: *val.Quantity, Product: product}
	
			user := userAndOrdersMap[val.Id]
			user.Orders = append(user.Orders, productOrderWithProduct)

			userAndOrdersMap[val.Id] = user
		}
	}

	for _, val := range userAndOrdersMap {
		userWithOrders = append(userWithOrders, val)
	}
 
	return userWithOrders, nil
}


