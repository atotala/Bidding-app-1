package routes

import (
	"encoding/json"
	"net/http"
	"s01/src/bid"
	"s01/src/db"

	"github.com/gorilla/mux"
)

//MaxBidData structure store the info of max bid placed for an item
type MaxBidData struct {
	MaxBId float64 `json:"Max_Bid"`
}

//RecordBid function take user input and store the bid in db
func RecordBid(w http.ResponseWriter, r *http.Request) {

	var reqData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("json format is expected"))
		return
	}
	item := reqData["item"].(string)
	userName := reqData["user_name"].(string)
	bidPrice := reqData["bid_price"].(float64)
	otpt := bid.RecordBid(item, userName, bidPrice)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(otpt)
}

//DeleteBid function delte the bid placed by user for an item
func DeleteBid(w http.ResponseWriter, r *http.Request) {

	var reqData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("json format is expected"))
		return
	}
	item := reqData["item"].(string)
	userName := reqData["user_name"].(string)
	otpt := bid.DeleteBid(item, userName)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(otpt)
}

//MaxBid function respond with the max bid for an item
func MaxBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemName := params["item"]
	data := MaxBidData{MaxBId: bid.MaxBid(itemName)}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)

}

//ItemAllBid function respond with the all bids for an item
func ItemAllBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemName := params["item"]
	data := bid.GetAllBidsForItem(itemName)
	w.Header().Set("Content-type", "application/json")
	if len(data) != 0 {
		json.NewEncoder(w).Encode(data)
	}

	json.NewEncoder(w).Encode("No Data Found")
}

//UserAllBid function respond with the all the bids placed by an user
func UserAllBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["user"]
	data := bid.GetAllBidsForUser(userName)
	w.Header().Set("Content-type", "application/json")
	if len(data) != 0 {
		json.NewEncoder(w).Encode(data)

	} else {

		json.NewEncoder(w).Encode("No Data Found")
	}
}

//AddItem function add an item in existing list of items for sale
func AddItem(w http.ResponseWriter, r *http.Request) {

	var reqData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("json format is expected"))
		return
	}
	item := reqData["item"].(string)
	added := db.AddItem(item)
	w.Header().Set("Content-type", "application/json")
	if added {
		json.NewEncoder(w).Encode("Item Added for Sale")
	} else {
		json.NewEncoder(w).Encode("Item Already available for Sale")

	}
}

//DeleteItem function delete the item from list of item for sale and also delete all the bid placed against that item
func DeleteItem(w http.ResponseWriter, r *http.Request) {

	var reqData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("json format is expected"))
		return
	}
	item := reqData["item"].(string)
	deleted := db.DeleteItem(item)
	w.Header().Set("Content-type", "application/json")
	if deleted {
		bid.DeleteBid(item, "")
		json.NewEncoder(w).Encode("Item Delete for Sale")
	} else {
		json.NewEncoder(w).Encode("Item is not present Sale")

	}
}

//NewRouter for accepting the requests
func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/recordbid", RecordBid).Methods("POST")
	router.HandleFunc("/deletebid", DeleteBid).Methods("POST")
	router.HandleFunc("/maxbid/{item}", MaxBid).Methods("GET")
	router.HandleFunc("/allbid/{item}", ItemAllBid).Methods("GET")
	router.HandleFunc("/userbid/{user}", UserAllBid).Methods("GET")
	router.HandleFunc("/additem", AddItem).Methods("POST")
	router.HandleFunc("/deleteitem", DeleteItem).Methods("POST")

	return router
}
