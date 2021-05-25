/* For license and copyright information please see LEGAL file in repository */

package http

// ContentType store
type ContentType struct {
	Type     mimeType
	SubType  mimeSubType
	Charset  string
	Boundary string
}

// GetContentType read all value about content in header
func (h *header) GetContentType() (c ContentType) {
	var contentType = h.Get(HeaderKeyContentType)
	return getContentType(contentType)
}

// getContentType read all value about content in header
func getContentType(contentType string) (c ContentType) {
	var mediaTypeFirst, mediaTypeSecond string
	var index int
	for i := 0; i < len(contentType); i++ {
		switch contentType[i] {
		case '/':
			mediaTypeFirst = contentType[:i]
			index = i + 1
		case ';':
			mediaTypeSecond = contentType[index:i]
			contentType = contentType[i+2:] // +2 due to have space after semicolon
		}
	}

	switch contentType[0] {
	case 'c': // charset=
		c.Charset = contentType[8:]
	case 'b': // boundary=
		c.Boundary = contentType[9:]
	}

	switch mediaTypeFirst {
	case "text":
		c.Type = ContentTypeMimeTypeText
	case "application":
		c.Type = ContentTypeMimeTypeApplication
	}

	switch mediaTypeSecond {
	case "html":
		c.SubType = ContentTypeMimeSubTypeHTML
	case "json":
		c.SubType = ContentTypeMimeSubTypeJSON
	}

	return
}

type mimeType uint16

// Standard HTTP content type
const (
	ContentTypeMimeTypeUnset       mimeType = 0
	ContentTypeMimeTypeText                 = 1
	ContentTypeMimeTypeApplication          = 2
)

type mimeSubType uint16

// Standard HTTP content type
const (
	ContentTypeMimeSubTypeUnset mimeSubType = 0
	ContentTypeMimeSubTypeHTML              = 1
	ContentTypeMimeSubTypeJSON              = 2
)
