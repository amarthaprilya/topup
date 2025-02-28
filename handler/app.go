package handler

import (
	"camera-rent/auth"
	"camera-rent/database"
	"camera-rent/middleware"
	"camera-rent/repository"
	"camera-rent/service"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartApp() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatal("Eror Db Connection")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin , Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods:    []string{"POST, OPTIONS, GET, PUT, DELETE"},
	}))

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}
	//load config
	ServerKeys := os.Getenv("SERVER_KEY")
	if ServerKeys == "" {
		log.Fatal("Server key not found in environment variables")
	}
	gateway, err := service.NewMidtransGateway(&service.Config{
		ServerKey: ServerKeys,
	})

	if err != nil {
		log.Fatal("Failed to initialize Midtrans gateway: ", err)
	}

	// user
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewService(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)

	if err != nil {
		panic(err)
	}

	user := router.Group("/api/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.DELETE("/:slug", userHandler.DeletedUser)
	user.PUT("/:slug", userHandler.UpdateUser)

	// product
	categoryRepository := repository.NewRepositoryCategory(db)
	categoryService := service.NewServiceCategory(categoryRepository)
	categoryHandler := NewCategoryHandler(categoryService)

	apiCategory := router.Group("/api/category")
	apiCategory.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), categoryHandler.CreateCategory)
	apiCategory.GET("/:id", categoryHandler.GetCategory)
	apiCategory.GET("/", categoryHandler.GetAllCategory)
	apiCategory.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), categoryHandler.DeleteCategory)
	apiCategory.PUT("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), categoryHandler.UpdateCategory)

	// product
	productRepository := repository.NewRepositoryProduct(db)
	productService := service.NewServiceProduct(productRepository, categoryRepository)
	productHandler := NewProductHandler(productService)

	apiProduct := router.Group("/api/products")
	apiProduct.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), productHandler.CreateProduct)
	apiProduct.GET("/:id", productHandler.GetProduct)
	apiProduct.GET("/", productHandler.GetAllProduct)
	apiProduct.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), productHandler.DeleteProduct)
	apiProduct.PUT("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), productHandler.UpdateProduct)

	// product
	bookingRepository := repository.NewRepositoryBooking(db)
	bookingService := service.NewServiceBooking(bookingRepository, productRepository, userRepository)
	bookingHandler := NewBookingHandler(bookingService)

	apiBooking := router.Group("/api/booking")
	apiBooking.POST("/rent-product", middleware.AuthMiddleware(authService, userService), bookingHandler.CreateBooking)
	apiBooking.GET("/:id", bookingHandler.GetBookingById)
	// apiBooking.GET("/booking-report", bookingHandler.GetBookingByUserID)
	apiBooking.GET("/booking-report", bookingHandler.GetAllBookings)
	apiBooking.DELETE("/:id", middleware.AuthRole(authService, userService), bookingHandler.DeleteBooking)
	// apiBooking.PUT("/:id", middleware.AuthRole(authService, userService), bookingHandler.UpdateProduct)

	// topUp
	topUpRepository := repository.NewRepositoryTopUp(db)
	topUpService := service.NewServiceTopUp(topUpRepository, userRepository)
	topUpHandler := NewTopUpHandler(topUpService)

	apiTopUp := router.Group("/api/topup")
	apiTopUp.POST("/", middleware.AuthMiddleware(authService, userService), topUpHandler.CreatetopUp)
	apiTopUp.GET("/:id", topUpHandler.GetTopUp)

	// topUp
	paymentSaldoRepository := repository.NewRepositoryPaymentSaldo(db)
	paymentSaldoService := service.NewServicePaymentSaldo(paymentSaldoRepository, userRepository, topUpRepository, gateway)
	paymentSaldoHandler := NewPaymentSaldoHandler(paymentSaldoService, authService)

	paymentSaldo := router.Group("/api/paymentsaldo")
	paymentSaldo.POST("/", paymentSaldoHandler.GetPaymentSaldoNotification)
	paymentSaldo.POST("/:id", middleware.AuthMiddleware(authService, userService), paymentSaldoHandler.DoPaymentSaldo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Jalankan server pada port yang sudah ditentukan
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
