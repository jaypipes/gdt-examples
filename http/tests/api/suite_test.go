// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.

package api_test

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jaypipes/gdt"
	gdtfix "github.com/jaypipes/gdt-fixtures"
	gdthttp "github.com/jaypipes/gdt-http"

	"github.com/jaypipes/gdt-examples/http/api"
)

const (
	dataFilePath = "testdata/fixtures.json"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

type dataset struct {
	Authors    interface{}
	Publishers interface{}
	Books      []*api.Book
}

func data() *dataset {
	f, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	data := &dataset{}
	if err = json.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}
	return data
}

func dataFixture() gdt.Fixture {
	f, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	f.Seek(0, io.SeekStart)
	fix, err := gdtfix.NewJSONFixture(f)
	if err != nil {
		panic(err)
	}
	return fix
}

func TestHTTP(t *testing.T) {
	// Register an HTTP server fixture that spins up the API service on a
	// random port on localhost
	logger := log.New(os.Stdout, "books_api_http: ", log.LstdFlags)
	c := api.NewControllerWithBooks(logger, data().Books)
	apiFixture := gdthttp.NewServerFixture(c.Router(), false /* useTLS */)
	gdt.Fixtures.Register("books_api", apiFixture)
	gdt.Fixtures.Register("books_data", dataFixture())

	// Construct a new gdt.TestSuite from the directory of this file
	ts, err := gdt.From(currentDir())
	if err != nil {
		t.Fatal(err)
	}
	ts.Run(gdt.NewContext(), t)
}

func TestHTTPS(t *testing.T) {
	// Register an HTTPS server fixture that spins up the API service on a
	// random port on localhost and a well-known cert for localhost/127.0.0.1
	logger := log.New(os.Stdout, "books_api_https: ", log.LstdFlags)
	c := api.NewControllerWithBooks(logger, data().Books)
	apiFixture := gdthttp.NewServerFixture(c.Router(), true /* useTLS */)
	gdt.Fixtures.Register("books_api", apiFixture)
	gdt.Fixtures.Register("books_data", dataFixture())

	// Construct a new gdt.TestSuite from the directory of this file
	ts, err := gdt.From(currentDir())
	if err != nil {
		t.Fatal(err)
	}
	ts.Run(gdt.NewContext(), t)
}
