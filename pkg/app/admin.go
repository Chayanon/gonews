package app

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Chayanon/gonews/pkg/model"
	"github.com/Chayanon/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		id := r.FormValue("id")
		if action == "delete" {
			err := model.DeleteNews(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminList(w, &view.AdminListData{
		List: list,
	})
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}

		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
			fp, err := os.Create("upload/" + fileName)
			if err == nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}
		err := model.CreateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// title := r.FormValue("title")
		// detail := r.FormValue("detail")
		// log.Println(title)
		// log.Println(detail)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		n.Title = r.FormValue("title")
		n.Detail = r.FormValue("detail")
		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
			fp, err := os.Create("upload/" + fileName)
			if err == nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}
		err := model.UpdateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin//list", http.StatusSeeOther)
		return
	}
	view.AdminEdit(w, n)
}
