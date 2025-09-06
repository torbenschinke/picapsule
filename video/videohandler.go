package video

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.wdy.de/nago/pkg/blob"
)

func NewHandler(uc UseCases, store blob.Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		optVid, err := uc.FindByID(ID(id))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if optVid.IsNone() {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		vid := optVid.Unwrap()

		optReader, err := store.NewReader(request.Context(), vid.BlobKey)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if optReader.IsNone() {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		reader := optReader.Unwrap()
		defer reader.Close()

		file := reader.(io.ReadSeeker)
		serveVideoFromReadSeeker(writer, request, file, vid.Name, vid.CreatedAt, "video/mp4", vid.Size)
	}
}

func serveVideoFromReadSeeker(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, filename string, modTime time.Time, contentType string, size int64) {
	// Content-Type
	if contentType == "" {
		contentType = "video/mp4" // fallback
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Accept-Ranges", "bytes")

	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" {
		// Ganzer Stream
		w.Header().Set("Content-Length", fmt.Sprint(size))
		http.ServeContent(w, r, filename, modTime, rs) // nutzt ReadSeeker
		return
	}

	if !strings.HasPrefix(rangeHeader, "bytes=") {
		http.Error(w, "invalid range", http.StatusRequestedRangeNotSatisfiable)
		return
	}

	rangeSpec := strings.TrimPrefix(rangeHeader, "bytes=")
	parts := strings.Split(rangeSpec, "-")
	if len(parts) != 2 {
		http.Error(w, "invalid range", http.StatusRequestedRangeNotSatisfiable)
		return
	}

	var start, end int64
	var err error

	// start
	if parts[0] != "" {
		start, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			http.Error(w, "invalid range", http.StatusRequestedRangeNotSatisfiable)
			return
		}
	}
	// end
	if parts[1] != "" {
		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			http.Error(w, "invalid range", http.StatusRequestedRangeNotSatisfiable)
			return
		}
	} else {
		end = size - 1
	}

	if start > end || start < 0 || end >= size {
		http.Error(w, "invalid range", http.StatusRequestedRangeNotSatisfiable)
		return
	}

	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, size))
	w.Header().Set("Content-Length", fmt.Sprint(end-start+1))
	w.WriteHeader(http.StatusPartialContent)

	_, err = rs.Seek(start, io.SeekStart)
	if err != nil {
		http.Error(w, "seek error", http.StatusInternalServerError)
		return
	}
	if _, err := io.CopyN(w, rs, end-start+1); err != nil {
		slog.Error("failed to read seek", slog.String("error", err.Error()))
	}
}
