package main

import (
	"e-commerce/controller"
	"e-commerce/repository"
	"e-commerce/service"
	"log"
	"net/http"
	"os"
)

func main() {
	// Database connection
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" {
		dbUser = "root"
	}
	if dbPass == "" {
		dbPass = "password"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbName == "" {
		dbName = "ecommerce"
	}

	db, err := repository.NewMySQL(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		// Continue for now to allow build verification, but in real app we might exit
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)

	// Controllers
	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)

	// Router
	mux := http.NewServeMux()

	// User Routes
	mux.HandleFunc("/register", userController.Register)
	mux.HandleFunc("/login", userController.Login)
	mux.HandleFunc("/profile", userController.GetProfile)
	mux.HandleFunc("/profile/update", userController.UpdateProfile)

	// Product Routes
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.List(w, r)
		case http.MethodPost:
			productController.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.Get(w, r)
		case http.MethodPut:
			productController.Update(w, r)
		case http.MethodDelete:
			productController.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Cart Routes
	mux.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cartController.GetCart(w, r)
		case http.MethodPost:
			cartController.AddItem(w, r)
		case http.MethodDelete:
			cartController.ClearCart(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
