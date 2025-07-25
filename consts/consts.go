package consts

const (
	RoleIdAdmin    = iota + 1
	RoleIdManager  = 2
	RoleIdCustomer = 3

	RoleAdmin    = "ADMIN"
	RoleManager  = "MANAGER"
	RoleCustomer = "CUSTOMER"

	DefaultPageSize = 10
	DefaultPage     = 1

	PermissionUserCreate = "user.create" // Permission to create a new user
	PermissionUserUpdate = "user.update" // Permission to update an existing user's information
	PermissionUserFetch  = "user.fetch"  // Permission to fetch a specific user's data
	PermissionUserList   = "user.list"   // Permission to list all users
	PermissionUserDelete = "user.delete" // Permission to delete a user

	PermissionProductCreate = "product.create" // Permission to create a new product
	PermissionProductUpdate = "product.update" // Permission to update an existing product
	PermissionProductFetch  = "product.fetch"  // Permission to fetch a specific product
	PermissionProductList   = "product.list"   // Permission to list products
	PermissionProductDelete = "product.delete" // Permission to delete an product
)

var RoleMap = map[int]string{
	RoleIdAdmin:    RoleAdmin,
	RoleIdManager:  RoleManager,
	RoleIdCustomer: RoleCustomer,
}
