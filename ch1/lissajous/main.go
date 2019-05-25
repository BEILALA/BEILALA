@@ -13,17 +13,21 @@ import (
	"image/color"
	"image/gif"
	"io"
	//!-main
	"log"
	//!+main
	"math"
	"math/rand"
	//!-main
	"net/http"
	//!+main
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
 @@ -33,6 +37,11 @@ const (

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
