package route

import (
  "net/http"
  "html/template"
  "path"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
  // Merge and build folder and files to a path
  var filePath = path.Join("views", "welcome.html")
  var tmpl, err = template.ParseFiles(filePath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // Make a data
  var data = map[string]interface{}{
    "title": "Xavla Web Framework",
  }

  // Send data to Template Engine
  err = tmpl.Execute(w, data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}