package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	State   int `json:"state"`
	Version int `json:"version"`
	Data    struct {
		Products []struct {
			Sort            int           `json:"__sort"`
			Ksort           int           `json:"ksort"`
			Time1           int           `json:"time1"`
			Time2           int           `json:"time2"`
			ID              int           `json:"id"`
			Root            int           `json:"root"`
			KindID          int           `json:"kindId"`
			SubjectID       int           `json:"subjectId"`
			SubjectParentID int           `json:"subjectParentId"`
			Name            string        `json:"name"`
			Brand           string        `json:"brand"`
			BrandID         int           `json:"brandId"`
			SiteBrandID     int           `json:"siteBrandId"`
			Sale            int           `json:"sale"`
			PriceU          int           `json:"priceU"`
			SalePriceU      int           `json:"salePriceU"`
			AveragePrice    int           `json:"averagePrice"`
			Benefit         int           `json:"benefit"`
			Pics            int           `json:"pics"`
			Rating          int           `json:"rating"`
			Feedbacks       int           `json:"feedbacks"`
			Colors          []interface{} `json:"colors"`
			Sizes           []struct {
				Name     string `json:"name"`
				OrigName string `json:"origName"`
				Rank     int    `json:"rank"`
				OptionID int    `json:"optionId"`
			} `json:"sizes"`
			DiffPrice    bool   `json:"diffPrice"`
			PanelPromoID int    `json:"panelPromoId,omitempty"`
			PromoTextCat string `json:"promoTextCat,omitempty"`
			IsNew        bool   `json:"isNew,omitempty"`
		} `json:"products"`
	} `json:"data"`
}

func main() {

	resp, err := http.Get("https://catalog.wb.ru/catalog/autoproduct7/catalog?appType=1&couponsGeo=2,12,7,3,6,13,21&curr=rub&dest=-1113276,-140291,-897992,12358481&emp=0&lang=ru&locale=ru&pricemarginCoeff=1.0&reg=0&regions=64,58,83,4,38,80,33,70,82,86,30,69,1,48,22,66,31,40&spp=0&subject=2011;2013;2216;2263;2394;2756;2781;2888;2965;3074;3139;3845;3848;4122;4199;4395;4657;4658;4699;4904;4905;4906;4907;4908;4930;5016;5409;5410;5411;5412;5413;5414;5464;5465;5533;5551;5566;5567;5568;5569;5570;5571;5728;5729;5749;5750;5751;5752;5753;5754;5755;5756;5757;5758;5759;5760;5761;5817;5818;5819;5820;5821;5822;5823;5824;5825;5831;5870;5871;5872;5873;5886;5887;6018;6019;6020;6173;6243;6247;6248;6267;6268;6278;6279;6280;6281;6282;6284;6285;6287;6289;6290;6291;6292;6293;6300;6301;6584;6587;7100")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result.Data.Products {
		fmt.Println(rec.Name, rec.Brand, rec.SalePriceU/100, rec.Rating)
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
