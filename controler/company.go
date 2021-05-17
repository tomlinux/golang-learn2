package controler

import (
	"TwoProject/helper"
	"TwoProject/model"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

func RegisterRoutes() {

	http.HandleFunc("/", listCompanies)
	http.HandleFunc("/companies", listCompanies)
	http.HandleFunc("/companies/seed", seed)
	http.HandleFunc("/companies/add", addCompanies)
	http.HandleFunc("/companies/edit/", editCompanies)
	http.HandleFunc("/companies/delete/", deleteCompany)
	http.HandleFunc("/test", test)
}

func test(w http.ResponseWriter, r *http.Request) {
		//r.SetBasicAuth("admin", "123456")
		w.Write([]byte("今天天气正好"))


}

func seed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("seed"))
}

// 公司所有信息列表
func listCompanies(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("hello,world!")
	//resp, _ := http.Get("https://www.taobao.com")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(bytes))
	//w.Write(bytes)
	companies, err := model.GetAllCompany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		funcMap := template.FuncMap{"add": helper.Add}
		t := template.New("companies").Funcs(funcMap)
		t, err = t.ParseFiles("./templates/layout.html", "./templates/company/list.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		//log.Println(companies)
		//str2, _ := os.Getwd()
		//log.Println(str2)
		t.ExecuteTemplate(w, "layout", companies)
	}
}

// 添加公司信息数据
func addCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t := template.New("company-add")
		t, err := t.ParseFiles("./templates/layout.html", "./templates/company/add.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			t.ExecuteTemplate(w, "layout", nil)
		}
	case http.MethodPost:
		//log.Printf("hello world")
		//newCompany2 := new(model.Company)
		newCompany := model.Company{}
		newCompany.ID = r.PostFormValue("id")
		newCompany.Name = r.PostFormValue("name")
		newCompany.NickName = r.PostFormValue("nickName")
		log.Println(newCompany)

		err := newCompany.Insert()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
		}
	}
}

// 编辑公司信息数据
func editCompanies(w http.ResponseWriter, r *http.Request) {
	log.Println("hello 进入编辑栏目")
	//matched, err := regexp.MatchString("/companies/edit/([a-zA-Z0-9]*$)", str)
	idPattern := regexp.MustCompile("/companies/edit/([a-zA-Z0-9]*$)")
	matches := idPattern.FindStringSubmatch(r.URL.Path)
	log.Println(matches)
	if len(matches) > 0 {
		id := matches[1]
		switch r.Method {
		case http.MethodGet:
			//id, err := strconv.ParseInt(id, 10, 64)
			//if err != nil {
			//	w.WriteHeader(http.StatusInternalServerError)
			//	w.Write([]byte(err.Error()))
			//}
			company, err := model.GetCompanyById(id)
			log.Printf("编辑栏目：%s,%#v,%T \n", company, company, company)
			if err == nil {
				t := template.New("company-edit")
				t, err := t.ParseFiles("./templates/layout.html", "./templates/company/edit.html")
				if err == nil {
					t.ExecuteTemplate(w, "layout", company)
					return
				}
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return

		case http.MethodPost:
			company := &model.Company{}
			company.ID = r.PostFormValue("id")
			company.Name = r.PostFormValue("name")
			company.NickName = r.PostFormValue("nickName")
			err := company.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				//t, _ := template.New("success").Parse(`{{ define "content"}}{{.}}!{{end}}`)
				//t.ExecuteTemplate(w, "layout", "<script>alert('更新成功')</script>")
				//http.Redirect(w, r, "/companies", http.StatusSeeOther)
				//更新成功并且跳转
				w.Write([]byte(`<script>alert('更新成功');window.location.href="/companies";</script>`))
				//http.Redirect(w, r, "/companies", http.StatusSeeOther)
			}
			return
		}

	}
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	idPattern := regexp.MustCompile(`/companies/delete/([a-zA-Z0-9]*$)`)
	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		id := matches[1]

		if r.Method == http.MethodDelete {
			err := model.DeleteCompany(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

//
