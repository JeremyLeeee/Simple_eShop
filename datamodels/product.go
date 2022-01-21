package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" eshop:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" eshop:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" eshop:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" eshop:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" eshop:"ProductUrl"`
}
