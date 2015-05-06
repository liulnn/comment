package comment

import (
	"testing"
)

func TestCommentsCtrlPost(t *testing.T) {
	var topicId string = "1"
	req := http.NewRequest("POST", fmt.Printf("/topics/%d/comments"), nil)
	c := CommentsCtrl{Request: req, PathParams: []string{topicId}, AccountId: 1}
	resp := c.Post()
	switch resp.StatusCode {
	case 500:
		t.Error("server error")
	case 201:
		url, err := resp.Location()
		if err != nil || url == nil {
			t.Error("response's location error")
		}
		etag := resp.Header.Get("ETag")
		if etag == nil {
			t.Error("response header has no etag")
		}
		h := md5.New()
		io.WriteString(h, string(resp.Body))
		buffer := bytes.NewBuffer(nil)
		fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
		if etag != buffer.String() {
			t.Error("response's etag is err")
		}
	}
}
