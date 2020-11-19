package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ebonynon/xEatz/customers"
	"github.com/ebonynon/xEatz/fooditems"
	"github.com/ebonynon/xEatz/orders"
	"github.com/ebonynon/xEatz/restaurants"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	getPort := func(KEY string) string {

		ENV, Okay := os.LookupEnv(KEY)

		if !Okay {

			ENV = "3080"

		}
		return ENV

	}

	r.Post("/addrestaurant", restaurants.AddARestaurant)
	r.Get("/getrestaurants", restaurants.GetRestaurants)
	r.Get("/getrestaurant/{id}", restaurants.GetARestaurant)
	r.Delete("/droprestaurant/{id}", restaurants.DeleteARestaurant)

	r.Post("/addfooditem", fooditems.AddAFoodItem)
	r.Get("/getfooditems", fooditems.GetFoodItems)
	r.Get("/getfooditem/{id}", fooditems.GetAFoodItem)
	r.Delete("/dropfooditem/{id}", fooditems.DeleteAFoodItem)

	r.Post("/addcustomer", customers.AddACustomer)
	r.Get("/getcustomers", customers.GetCustomers)
	r.Get("/getcustomer/{id}", customers.GetACustomer)
	r.Delete("/dropcustomer/{id}", customers.DeleteACustomer)

	r.Post("/addorder", orders.AddAOrder)
	r.Get("/getorders", orders.GetOrders)
	r.Get("/getorder/{id}", orders.GetAOrder)
	r.Delete("/droporder/{id}", orders.DeleteAOrder)

	http.ListenAndServe(":"+getPort("PORT"), r)

}
