package handlers

import(
	"log"
	"net/http"
	"../data"
	"regexp"
	"strconv"
	//"encoding/json"
)

type Product struct{
	l*log.Logger
}

func NewProduct(l*log.Logger) *Product {
	return &Product{l}
}

func (p*Product)ServeHTTP(rw http.ResponseWriter,rd *http.Request)  {

	if rd.Method == http.MethodGet{
		p.getProducts(rw,rd)
		return
	}

	if rd.Method == http.MethodPost{
		p.addProducts(rw,rd)
		return 
	}
	if rd.Method == http.MethodPut{
		p.l.Println("insidePUT")


		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(rd.URL.Path,-1)

		p.l.Println(rd.URL.Path)

		if len(g) != 1{
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2{
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}

		idString :=g[0][1]
		id, err := strconv.Atoi(idString)

		if err !=nil{
			http.Error(rw,"Invalid URL",http.StatusBadRequest)
			return
		}

		p.updateProducts(id,rw,rd)
		return
		//p.l.Println("got id", id)




	}



	rw.WriteHeader(http.StatusMethodNotAllowed)
}



func (p*Product)getProducts(rw http.ResponseWriter, r *http.Request)  {
	

	p.l.Println("Handle Get Products")

	lp:= data.GetProducts()
	err := lp.ToJSON(rw)
	
	if err != nil{
		http.Error(rw,"Unable to marshal json", http.StatusInternalServerError)
	}


}


func (p*Product)addProducts(rw http.ResponseWriter, rd *http.Request)  {
	
	p.l.Println("Handle Post Products")
	
	prod := &data.Product{}
	err := prod.FromJSON(rd.Body)
	

	if err != nil{
		http.Error(rw,"Unable to unmarshal json",http.StatusBadRequest)
	}

	p.l.Printf("Prod : %#v",prod)
	data.AddProduct(prod)
}


func (p Product) updateProducts(id int,rw http.ResponseWriter, rd*http.Request) {
	p.l.Println("Inside Update Products")
	prod := &data.Product{}
	err :=prod.FromJSON(rd.Body)
	if err !=nil{
		http.Error(rw,"Unable to unmarshal Json",http.StatusBadRequest)
	}

	e:=data.UpdateProduct(id,prod)
	if e == data.ErrProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}

	if err !=nil{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}


}