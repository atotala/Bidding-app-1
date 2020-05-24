###Setup###
1.Setup environment variable GOPATH on your machine to a folder which will go workspace going forward.
2. Create folder $GOPATH/src.
3. Copy enitre s01 folder under above path

###Usage###
To run application locally, please excute the following command
1. dep ensure
2. go run app.go

###APIs description###
	router.HandleFunc("/delteitem", DeleteItem).Methods("POST")
1. Place a bid - POST (http://localhost:8080/recordbid)
    DATA:{
        "item":"",
        "user_name":"",
        "bid_price":""
    }

2. Delete a bid - POST (http://localhost:8080/deletebid)
    DATA:{
        "item":"",
        "user_name":""
    }

3. Max bid on an item -  GET (http://localhost:8080/maxbid/{item})
4. All bids placed on an item - GET (http://localhost:8080/allbid/{item})
5. All bids placed by one user - GET (http://localhost:8080/userbid/{user})
6. Add new item for bidding - POST (http://localhost:8080/additem) (For Admin user)
    DATA:{
        "item":""
    }
7. Delete item for bidding and remove all the bids placed against the ime- POST (http://localhost:8080/deleteitem) (For Admin user)
    DATA:{
        "item":""
    }



