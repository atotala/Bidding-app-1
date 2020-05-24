package db

type items struct {
	itemName    string
	itemID      int
	maxBidPrice float64
}

var dataStore = make(map[string]items)

//InitializeDB function load the basic items in memory for sale
func InitializeDB() bool {
	dataStore["Television"] = items{"Television", 123, 0}
	dataStore["Sofa"] = items{"Sofa", 456, 0}
	dataStore["Suitcase"] = items{"Suitcase", 789, 0}
	return true
}

//AddItem function add the items in data store
func AddItem(item string) bool {
	_, exist := dataStore[item]
	if !exist {
		dataStore[item] = items{itemName: item, itemID: 123, maxBidPrice: 0}
		return true
	}
	return false

}

//ItemExistInDB check if the item is available for sale
func ItemExistInDB(item string) bool {
	_, exist := dataStore[item]
	return exist
}

//DeleteItem function delete the item available for sale
func DeleteItem(item string) bool {
	_, exist := dataStore[item]
	if exist {
		delete(dataStore, item)
		return true
	}
	return false

}
