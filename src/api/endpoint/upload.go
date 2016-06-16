package endpoint

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/devinmcgloin/morph/src/api/AWS"
	"github.com/devinmcgloin/morph/src/api/SQL"
	"github.com/julienschmidt/httprouter"
)

// UploadHandler manages uploading the original file to aws.
// TODO: In the future it should also spin off worker threads to
// handle compression, and rendering other sizes for the image.
func UploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseMultipartForm(32 << 20)

	var err error

	var IID uint64

	IID = getNewIID()

	img := SQL.Img{
		IID:         IID,
		PublishTime: time.Now(),
		CaptureTime: time.Now(),
		UID:         1,
		LID:         SQL.ToNullInt64("0"),
	}

	source := SQL.ImgSource{
		IID:  IID,
		Size: "orig",
		SID:  getNewSID(),
	}

	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {

			// open uploaded
			var infile multipart.File
			infile, err = hdr.Open()

			if err != nil {
				log.Printf("Error while reading in image %s", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			var buf bytes.Buffer
			var written int64
			written, err = buf.ReadFrom(infile)
			if err != nil {
				log.Printf("Error while reading in image %s", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			filename := fmt.Sprintf("%d_orig.jpg", IID)

			source.URL, err = AWS.UploadImageAWS(buf.Bytes(), written, filename, "morph-content", "us-east-1")
			if err != nil {
				log.Printf("Error while uploading image %s", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

		}
	}

	log.Println(source.URL)

	err = SQL.AddSrc(source)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = SQL.AddImg(img)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	newURL := fmt.Sprintf("/i/%d/edit", IID)

	log.Println(newURL)

	http.Redirect(w, r, newURL, 302)

}
