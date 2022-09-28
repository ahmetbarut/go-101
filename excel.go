package main

// Kullandığım paketler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
	"github.com/xuri/excelize/v2"
)

// Veri çekerken gelen yapıya göre bir struct oluşturup ona göre işlenmesi için
// `jsonresp` ve `product` structlarını kullandım.

// Başlangıç seviyeme göre, öğrendiğim;
// struct'lar 3 girintiye ayrılıyor. 1. girinti de gelen veriye nasıl erişmek
// için isim belrtmek. Bunu bir değişken gibi düşünebiliriz.
// 2. girinti de gelecek olan verinin türü
// 3. girinti de ise belirttiğimiz özellik ismine karşı gelecek olan değer.

type jsonresp struct {
	Products []product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

type product struct {
	Id                 int64    `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float32  `json:"price"`
	DiscountPercentage float32  `json:"discountPernetage"`
	Rating             float32  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func main() {
	// https://dummyjson.com/products

	// Yeni bir istek örneği döndürür ve kullanılmak üzere RelaseRequest'i kullanabiliriz.
	req := fasthttp.AcquireRequest()

	// Gelecek olan verinin struct'ını hazırlıyoruz
	var prod jsonresp

	// uygulama bitmeden önce isteği kapatmak için defer keyword'ü kullanılır
	defer fasthttp.ReleaseRequest(req)

	// şimdi istek atacağımız url'i belirterek gidiyoruz
	req.SetRequestURI("https://dummyjson.com/products")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(resp.Body()), &prod)

	if err != nil {
		fmt.Println(err)
	}
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Title")
	f.SetCellValue("Sheet1", "C1", "Description")
	f.SetCellValue("Sheet1", "D1", "Price")
	f.SetCellValue("Sheet1", "E1", "DiscountPercentage")
	f.SetCellValue("Sheet1", "F1", "Rating")
	f.SetCellValue("Sheet1", "G1", "Stock")
	f.SetCellValue("Sheet1", "H1", "Brand")
	f.SetCellValue("Sheet1", "I1", "Category")
	f.SetCellValue("Sheet1", "J1", "Thumbnail")

	for i := 0; i < len(prod.Products); i++ {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), prod.Products[i].Id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), prod.Products[i].Title)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), prod.Products[i].Description)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), prod.Products[i].Price)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), prod.Products[i].DiscountPercentage)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), prod.Products[i].Rating)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+2), prod.Products[i].Stock)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+2), prod.Products[i].Brand)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+2), prod.Products[i].Category)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+2), prod.Products[i].Thumbnail)
	}

	f.SetActiveSheet(len(prod.Products))

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
