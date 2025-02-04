// Copyright 2022 Arduino SA
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pkgs

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/arduino/arduino-create-agent/gen/indexes"
	"github.com/sirupsen/logrus"
	"go.bug.st/downloader/v2"
)

// Indexes is a client that implements github.com/arduino/arduino-create-agent/gen/indexes.Service interface
type Indexes struct {
	Log    *logrus.Logger
	Folder string
}

// Add downloads the index file found at the url contained in the payload, and saves it in the Indexes Folder.
// If called with an already existing index, it overwrites the file.
// It can fail if the payload is not defined, if it contains an invalid url.
func (c *Indexes) Add(ctx context.Context, payload *indexes.IndexPayload) (*indexes.Operation, error) {
	// Parse url
	indexURL, err := url.Parse(payload.URL)
	if err != nil {
		return nil, indexes.MakeInvalidURL(err)
	}

	// Download tmp file
	filename := b64.StdEncoding.EncodeToString([]byte(url.PathEscape(payload.URL)))
	path := filepath.Join(c.Folder, filename+".tmp")
	d, err := downloader.Download(path, indexURL.String())
	if err != nil {
		return nil, err
	}
	err = d.Run()
	if err != nil {
		return nil, err
	}

	// Move tmp file
	err = os.Rename(path, filepath.Join(c.Folder, filename))
	if err != nil {
		return nil, err
	}

	return &indexes.Operation{Status: "ok"}, nil
}

// Get reads the index file from the Indexes Folder, unmarshaling it
func (c *Indexes) Get(ctx context.Context, uri string) (index Index, err error) {
	filename := b64.StdEncoding.EncodeToString([]byte(url.PathEscape(uri)))
	path := filepath.Join(c.Folder, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return index, err
	}

	err = json.Unmarshal(data, &index)
	if err != nil {
		return index, err
	}

	return index, nil
}

// List reads from the Indexes Folder and returns the indexes that have been downloaded
func (c *Indexes) List(context.Context) ([]string, error) {
	// Create folder if it doesn't exist
	_ = os.MkdirAll(c.Folder, 0755)
	// Read files
	files, err := os.ReadDir(c.Folder)

	if err != nil {
		return nil, err
	}

	res := []string{}
	for _, file := range files {
		// Select only files that begin with http
		decodedFileName, _ := b64.URLEncoding.DecodeString(file.Name())
		fileName := string(decodedFileName)
		if !strings.HasPrefix(fileName, "http") {
			continue
		}
		path, err := url.PathUnescape(fileName)
		if err != nil {
			c.Log.Warn(err)
		}
		res = append(res, path)
	}

	return res, nil
}

// Remove deletes the index file from the Indexes Folder
func (c *Indexes) Remove(ctx context.Context, payload *indexes.IndexPayload) (*indexes.Operation, error) {
	filename := b64.StdEncoding.EncodeToString([]byte(url.PathEscape(payload.URL)))
	err := os.RemoveAll(filepath.Join(c.Folder, filename))
	if err != nil {
		return nil, err
	}
	return &indexes.Operation{Status: "ok"}, nil
}
