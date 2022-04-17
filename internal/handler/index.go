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

//go:embed index.html.gotmpl
var indexRaw string
var indexTpl = template.Must(template.New("index").Parse(indexRaw))

type indexVars struct {
	Instances []string
}

func (h *handler) getIndex(w http.ResponseWriter, r *http.Request) {
	var vars indexVars
	for name := range h.targets {
		vars.Instances = append(vars.Instances, name)
	}
	sort.Strings(vars.Instances)

	if err := indexTpl.Execute(w, vars); err != nil {
		h.log.WithError(err).Error("failed to serve index template")
	}
}

func (h *handler) getFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	ico := bytes.NewBuffer(faviconRaw)
	if _, err := io.Copy(w, ico); err != nil {
		h.log.WithError(err).Error("failed to serve favicon")
	}
}
