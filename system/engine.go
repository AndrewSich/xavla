package system

import (
  "net/http"
  "html/template"
  "path"
  "xavla/route"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
  var filePath = path.Join("views", route.Index)
  var tmpl, err = template.ParseFiles(filePath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = tmpl.Execute(w, route.Data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

