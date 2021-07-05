package data

import(
	"fmt"
	"time"
	"io"
	"encoding/json"
)

type Product struct{

	Name string `json:"PrdName"`
	ID int `json:"id"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string`json:"sku"`
	CreatedOn string`json:"-"`
	UpdatedOn string`json:"-"`
	DeletedOn string`json:"-"`

}

type Products []*Product


func (p*Product) FromJSON(r io.Reader) error  {
	
	e := json.NewDecoder(r)
	return e.Decode(p)
}



func (p*Products) ToJSON(w io.Writer) error  {
	err := json.NewEncoder(w)
	return err.Encode(p)
	
}


func GetProducts() Products  {
	return productList
	
}




var productList = []*Product{
	&Product{
		ID: 1,
		Name: "coffe MUG",
		Description: "milky coffe",
		SKU:  "abc234",
		Price: 2.5,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:2,
		Name: "Indian Tea",
		Description:"hard tea milky",
		SKU:"yuy132",
		Price:3.5,
		CreatedOn:time.Now().UTC().String(),
		UpdatedOn:time.Now().UTC().String(),

	},
}



func AddProduct(p*Product)  {
	p.ID = getNextID()
	productList = append(productList,p)	
}


func getNextID() int  {
	
	lp := productList[len(productList)-1]
	return (lp.ID+1)

}


func UpdateProduct(id int, p*Product) error  {

_,pos,err := findProduct(id)

if err !=nil{
	return err
}	

p.ID= id
productList[pos]= p
return nil
}


var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product,int, error)  {
	
	for i,p:=range productList{
		if p.ID == id{
			return p,i,nil
		}
	}

return nil,-1, ErrProductNotFound
}