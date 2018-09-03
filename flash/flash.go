package flash

import (
	"encoding/base64"
	"net/http"
	"time"
)

// SetFlash function
// set cookie by name
func SetFlash(w http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{
		Name:  name,
		Value: Encode(value),
	}
	http.SetCookie(w, c)
}

// GetFlash function
// Get cookie value from name
func GetFlash(w http.ResponseWriter, req *http.Request, name string) ([]byte, error) {
	c, err := req.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := Decode(c.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{
		Name:    name,
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}
	http.SetCookie(w, dc)
	return value, nil
}

// Encode function encodes to string
func Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// Decode function decodes string
func Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
