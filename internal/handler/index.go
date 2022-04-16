package handler

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
	"net/http"
	"sort"
)

//go:embed favicon.ico
var faviconRaw []byte
var faviconReader = bytes.NewBuffer(faviconRaw)

//go:embed index.html.gotmpl
var indexRaw string
var indexTpl = template.Must(template.New("index").Parse(indexRaw))

type indexVars struct {
	Instances []string
}

func (h *Handler) getIndex(w http.ResponseWriter, r *http.Request) {
	var vars indexVars
	for name := range h.targets {
		vars.Instances = append(vars.Instances, name)
	}
	sort.Strings(vars.Instances)

	indexTpl.Execute(w, vars)
}

func (h *Handler) getFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, faviconReader)
}
