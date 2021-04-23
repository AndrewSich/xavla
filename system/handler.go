package system

import (
  "net/http"
  "html/template"
  "path"
  "xavla/route"
)

// Handler of static files route
func HandlerStatic(routeTarget, fileTarget string) {
  http.Handle(routeTarget,
    http.StripPrefix(routeTarget,
      http.FileServer(http.Dir(fileTarget))))
}

// Handler of index or home route
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
  var filePath = path.Join("views", route.Index)
  var tmpl, err = template.ParseFiles(filePath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  
  err = tmpl.Execute(w, route.Data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}