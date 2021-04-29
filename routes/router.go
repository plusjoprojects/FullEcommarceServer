// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// gin.SetMode(gin.ReleaseMode)

	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	// Auth Routes
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.GET("/auth", controllers.Auth)
	auth.GET("/users/index", controllers.UsersListIndex)
	auth.GET("/users/delete/:id", controllers.DeleteUser)
	auth.POST("/update", controllers.UpdateUser)

	// --------- Users ----------/
	user := r.Group("/users")
	// ~~~ User Roles ~~~ //
	user.POST("/roles/store", controllers.StoreUserRoles)
	user.POST("/roles/update", controllers.UpdateUserRole)
	user.GET("/roles/index", controllers.IndexUserRoles)
	user.GET("/roles/delete/:id", controllers.DeleteUserRole)
	user.GET("/clients/index", controllers.ClientsIndex)

	user.POST("/employee/store", controllers.StoreEmployee)
	user.GET("/employee/index", controllers.IndexEmployee)
	user.GET("/employee/delete/:id", controllers.DeleteEmployee)
	user.POST("/employee/update", controllers.UpdateEmployee)

	// Delegates
	user.POST("/delegates/store", controllers.StoreDelegate)
	user.GET("/delegates/index", controllers.IndexDelegates)
	user.GET("/delegates/delete/:id", controllers.DeleteDelegate)
	user.POST("/delegates/update", controllers.UpdateDelegate)
	user.POST("/delegates/itemsToDelegate/store", controllers.StoreItemsToDelegate)

	// ---------- Clients --------- //
	clients := r.Group("/clients")

	clients.POST("/store", controllers.StoreClient)
	clients.POST("/update", controllers.UpdateClient)
	clients.GET("/index", controllers.IndexClients)
	clients.GET("/index/forPurchase", controllers.IndexClientsForPurchase)
	clients.GET("/index/forSales", controllers.IndexClientsForSales)
	clients.GET("/destroy/:id", controllers.DestroyClient)

	// --------- Basics ------- //
	basics := r.Group("/basics")

	// UploadImage => For All
	basics.POST("/upload_image/:imageType", controllers.UpdateImage)

	// categories
	basics.POST("/categories/store", controllers.StoreCategories)
	basics.POST("/categories/update", controllers.UpdateCategory)
	basics.GET("/categories/destroy/:id", controllers.DestroyCategory)
	basics.GET("/categories/index", controllers.IndexCategories)

	// SubCategories
	basics.POST("/subCategories/store", controllers.StoreSubCategories)
	basics.POST("/subCategories/update", controllers.UpdateSubCategory)
	basics.GET("/subCategories/index", controllers.IndexSubCategories)
	basics.GET("/subCategories/index/:categories_id", controllers.IndexSubCategoriesWithID)
	basics.GET("/subCategories/destroy/:id", controllers.DestroySubCategory)

	// Colors
	basics.POST("/colors/store", controllers.StoreColors)
	basics.POST("/colors/update", controllers.UpdateColor)
	basics.GET("/colors/index", controllers.IndexColors)
	basics.GET("/colors/destroy/:id", controllers.DestroyColor)

	// Items
	basics.POST("/items/store", controllers.StoreItems)
	basics.POST("/items/update", controllers.UpdateItem)
	basics.GET("/items/index", controllers.IndexItems)
	basics.GET("/items/destroy/:id", controllers.DestroyItem)
	basics.GET("/items/show/:id", controllers.ShowItem)
	basics.POST("/items/search", controllers.SearchItems)

	// Companies
	basics.POST("/companies/store", controllers.StoreCompany)
	basics.POST("/companies/update", controllers.UpdateCompany)
	basics.GET("/companies/index", controllers.IndexCompanies)
	basics.GET("/companies/destroy/:id", controllers.DestroyCompany)

	// Storages
	basics.GET("/storages/index", controllers.IndexStorages)
	basics.GET("/storage/items/:scope", controllers.IndexStorageItems)

	// ------- Advance ------- //
	advance := r.Group("/advance")

	// Ads Methods
	advance.GET("/ads/index", controllers.IndexAds)
	advance.POST("/ads/store", controllers.StoreAds)
	advance.POST("/ads/update", controllers.UpdateAds)
	advance.GET("/ads/remove/:id", controllers.RemoveAds)

	advance.GET("/discountCodes/index", controllers.IndexDiscountCodes)
	advance.POST("/discountCodes/store", controllers.StoreDiscountCode)
	advance.POST("/discountCodes/update", controllers.UpdateDiscountCode)
	advance.GET("/discountCodes/remove/:id", controllers.DestroyDiscountCode)

	// Ads Items Methods
	advance.GET("/ad_items/index/:id", controllers.IndexAdItems)
	advance.POST("/ad_items/store", controllers.StoreAdItem)
	advance.GET("/ad_items/remove/:id", controllers.RemoveAdItem)

	// Notification
	advance.POST("/SendNotification", controllers.GlobalSendNotification)

	// ------- Purchases -------- //
	purchases := r.Group("/purchases")

	// StorePurchases
	purchases.POST("/store", controllers.AddPurchases)
	purchases.GET("/index", controllers.IndexPurchases)

	// ---------- Sales ----------- //
	sales := r.Group("/sales")

	sales.POST("/store", controllers.AddSale)
	sales.GET("/LastID", controllers.GetSalesLastID)

	receipt := r.Group("/receipt")

	receipt.POST("/store/catchReceipt", controllers.StoreCatchReceipt)
	receipt.POST("/store/paymentReceipt", controllers.StorePaymentReceipt)
	receipt.POST("/store/exchageReceipt", controllers.StoreExchangeReceipt)

	// ------------ Expenses ------------- //
	expenses := r.Group("/expenses")
	expenses.POST("/storeMajorExpenses", controllers.StoreMajorExpenses)
	expenses.GET("/indexMajorExpenses", controllers.IndexMajorExpenses)
	expenses.GET("/destroyMajorExpenses/:id", controllers.DestroyMajorExpenses)
	expenses.POST("/updateMajorExpenses", controllers.UpdateMajorExpenses)

	expenses.POST("/report/:id", controllers.MajorExpensesReport)

	// ------ Storages ---------- //
	storages := r.Group("/storages")
	storages.POST("/items/moveFromStorageToStorage", controllers.MoveFromStorageToStorage)
	storages.POST("/items/ChangeQty", controllers.ChangeStorageItemQty)

	// ------- Orders ---------- //
	orders := r.Group("/orders")
	orders.GET("/index/:status", controllers.IndexNewOrders)
	orders.GET("/show/:id", controllers.ShowOrder)
	orders.POST("/action", controllers.OrderActions)

	// -------- Pos ----------- //
	pos := r.Group("/pos")
	pos.GET("/items/index", controllers.IndexPosItems)
	pos.GET("/items/withCategoriesID/:id", controllers.IndexPosItemsWithCategoriesID)
	pos.GET("/items/search/:text", controllers.SearchPosItems)
	pos.POST("/order/store", controllers.StoreAppOrder)
	pos.GET("/items/barcode/:barcode", controllers.GetBarcodeItem)
	pos.GET("/report/today", controllers.IndexPosTodayOrders)

	// --------- Reports Controllers ---------//
	reports := r.Group("/reports")
	reports.GET("/index/:reportType", controllers.IndexReport)
	reports.POST("/withDates", controllers.ReportWithDates)
	reports.POST("/clients/purchasesClients/:id", controllers.ReportPurchasesClients)
	reports.POST("/clients/salesClients/:id", controllers.ReportSalesClients)
	reports.POST("/receipts/Purchases/:id", controllers.ReportsPurchases)
	reports.POST("/receipts/Sales/:id", controllers.ReportsSales)

	// -------- App controllers --------//
	app := r.Group("/app")
	// Index App
	app.GET("/index", controllers.AppIndex)
	app.GET("/category/index/:id/:page", controllers.IndexItemsAndSubCategoriesWithCategoriesID)
	app.GET("/category/loadMore/:id/:page", controllers.CategoriesItemsLoadMore)
	app.GET("/item/:id", controllers.AppShowItem)
	app.GET("/subCategory/index/:id/:page", controllers.AppIndexItemsWithSubCategories)
	app.POST("/items/indexWithIDS", controllers.AppIndexItemsWithIDS)
	app.GET("/items/indexAllItems/:type/:page", controllers.IndexShowAllItem)
	app.GET("/test", controllers.TestExternalCall)

	// Addresses Methods
	app.POST("/user/address/store", controllers.StoreUserAddress)
	app.GET("/user/address/destroy/:id", controllers.DestroyUserAddress)
	app.GET("/user/address/index/:id", controllers.IndexUserAddresses)

	// Orders Methods
	app.POST("/user/orders/store", controllers.StoreAppOrder)
	app.GET("/user/order/show/:id", controllers.AppShowOrder)

	// Discount Code Checker ..
	app.GET("/discountCode/check/:code/:user_id", controllers.CheckDiscountCodeValid)

	// Auth App
	app.GET("/user/auth/auth", controllers.AppAuth)
	app.POST("/user/auth/login", controllers.AppLoginController)
	app.POST("/user/auth/update", controllers.AppUpdateUser)

	// Search App
	app.GET("/search/:text", controllers.AppSearch)

	// Notification Methods
	app.POST("/user/notification/store", controllers.StoreNotificationToken)
	app.POST("/user/notification/update", controllers.UpdateNotificationTokenDetails)

	// Testes Group
	tests := r.Group("/tests")
	tests.GET("/test_joins", controllers.TestJoins)
	tests.GET("/items", controllers.TestItemsWithStorageCount)

	r.Run(":8082")
}
