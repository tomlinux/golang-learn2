package controler

import (
	"TwoProject/helper"
	"TwoProject/model"
	"encoding/json"
	"net/http"
)

func RegisterApiRoutes() {
	http.HandleFunc("/api/companies", helper.CheckAuth(getCompanies))
}


func getCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		companies, err := model.GetAllCompany()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			enc := json.NewEncoder(w)
			err = enc.Encode(companies)
		//[
		//	{
		//	"ID": "33232",
		//	"Name": "招财猫金融科技有限公司",
		//	"NickName": "招财猫"
		//	}
		//	]
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	case http.MethodPost:
		//r.Body
		//{
		//"id": "004",
		//"name": "Facebook",
		//"nickName": "FB"
		//}
		//w.Header().Set("Content-Type", "application/json; charset=utf-8")
		dec := json.NewDecoder(r.Body)
		c := model.Company{}
		err := dec.Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			err = c.Insert()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	}

}
