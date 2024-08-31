package version

import (
	"encoding/json"
	"github.com/dkrizic/testserver/meta"
	"net/http"
)

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	version := Version{Name: meta.ServiceName, Version: meta.Version}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(version)
}
