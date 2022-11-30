package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var products []product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
func returnProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	}

	for _, product := range products {
		if int64(product.ID) == id {
			res, err := json.Marshal(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "Product ID not found", http.StatusNotFound)
	return
}

func createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Couldn't read request body", http.StatusInternalServerError)
		return

	}
	var item product
	if err := json.Unmarshal(reqBody, &item); err != nil {
		http.Error(w, "Incorrect request data type", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if item.ID == product.ID {
			http.Error(w, "Product already exists", http.StatusBadRequest)
			return
		}
	}
	products = append(products, item)

	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id, err := strconv.ParseInt(vars["id"], 10, 64)
	//if err != nil {
	//	http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	//}
	//
	//var updatedProduct product
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//if err = json.Unmarshal(reqBody, &updatedProduct); err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}

	//for i, product := range product {
	//	if int64(product.ID) == id {
	//		product.ID = updatedProduct.ID
	//		productName = updatedProduct.Name
	//		product.Description = updatedProduct.Description
	//		product.Price = updatedProduct.Price
	//		products[i] = product
	//		if err = json.NewEncoder(w).Encode(product); err != nil {
	//			w.WriteHeader(http.StatusInternalServerError)
	//		}
	//	}
	//}
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	}

	for index, product := range products {
		if int64(product.ID) == id {
			products = append(products[:index], products[index+1:]...)
		}
	}
}

func initProducts() {
	byteSlice, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(byteSlice, &products); err != nil {
		log.Fatal(err)
	}

}

func main() {

	initProducts()
	r := mux.NewRouter()

	r.HandleFunc("/products", getProductsHandler)
	r.HandleFunc("/product", createNewProductHandler).Methods("POST")
	r.HandleFunc("/product", updateProductHandler).Methods("PUT")
	r.HandleFunc("/product", deleteProductHandler).Methods("Delete")
	r.HandleFunc("/product", returnProductHandler)

	log.Println("listen on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
