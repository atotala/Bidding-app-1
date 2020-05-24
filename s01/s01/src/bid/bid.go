package bid

import (
	"s01/src/db"
)

type itemBidDetail struct {
	itemName string
	bidPrice float64
	userName string
}

//UserBid structure store the info of item and their bid price
type UserBid struct {
	Item     string  `json: "Item Name"`
	BidPrice float64 `json: "Bid Price"`
}

//ItemBid structure store the info of user and bid price place by them
type ItemBid struct {
	UserName string  `json: "User Name"`
	BidPrice float64 `json: "Bid Price"`
}

var bidRecord = []itemBidDetail{}

//RecordBid function add the new bid from the user to bid Record
func RecordBid(item string, userName string, bidPrice float64) string {
	var addFlag = true
	if !db.ItemExistInDB(item) {
		return "Item is Not Available For Sale"
	}
	for i, elem := range bidRecord {
		if elem.itemName == item && elem.userName == userName {
			bidRecord[i].bidPrice = bidPrice
			addFlag = false
			break

		}
	}
	if addFlag {
		bidRecord = append(bidRecord, itemBidDetail{item, bidPrice, userName})
	}
	return "Bid Recorded"
}

//DeleteBid function delete the bid placed by user
func DeleteBid(item string, userName string) string {
	for i, elem := range bidRecord {
		if elem.itemName == item && (elem.userName == userName || userName == "") {
			bidRecord = append(bidRecord[:i], bidRecord[i+1:]...)
		}
	}

	return "Bid Deleted"
}

//GetAllBidsForUser function get the details of all the bids placed by a user
func GetAllBidsForUser(userName string) []UserBid {

	var data = []UserBid{}
	for _, elem := range bidRecord {
		if elem.userName == userName {
			data = append(data, UserBid{elem.itemName, elem.bidPrice})

		}
	}
	return data
}

//GetAllBidsForItem function get the details of all the bids placed for an item
func GetAllBidsForItem(item string) []ItemBid {
	var data = []ItemBid{}
	for _, elem := range bidRecord {
		if elem.itemName == item {
			data = append(data, ItemBid{elem.userName, elem.bidPrice})

		}
	}
	return data
}

// MaxBid function fetch the max bid placed for an item
func MaxBid(item string) float64 {
	var maxBid float64
	for _, elem := range bidRecord {
		if elem.itemName == item {
			if elem.bidPrice > maxBid {
				maxBid = elem.bidPrice
			}
		}
	}
	return maxBid
}
