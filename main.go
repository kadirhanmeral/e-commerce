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
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASS", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "ecommerce")

	db, err := repository.NewMySQL(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", userController.Register)
	mux.HandleFunc("/login", userController.Login)
	mux.HandleFunc("/profile", userController.GetProfile)
	mux.HandleFunc("/profile/update", userController.UpdateProfile)

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

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
