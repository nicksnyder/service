package debug

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/user"

	"github.com/nicksnyder/service/pkg/version"
)

// Serve serves a debug http endpoint.
func Serve(w http.ResponseWriter, r *http.Request) error {
	return WriteData(w)
}

// WriteData writes debug data to w.
func WriteData(w io.Writer) error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	u, err := user.Current()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(map[string]interface{}{
		"Version": version.Version(),
		"Host":    host,
		"User":    u,
	})
}
