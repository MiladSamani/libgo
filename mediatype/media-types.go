/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"../protocol"
)

// The currently registered types
const (
	MimeMainTypeDomain      = "domain"
	MimeMainTypeApplication = "application"
	MimeMainTypeAudio       = "audio"
	MimeMainTypeFont        = "font"
	MimeMainTypeExample     = "example"
	MimeMainTypeImage       = "image"
	MimeMainTypeMessage     = "message"
	MimeMainTypeModel       = "model"
	MimeMainTypeMultipart   = "multipart"
	MimeMainTypeText        = "text"
	MimeMainTypeVideo       = "video"
)

// MediaTypes store all data structure details
type MediaTypes struct{}

// RegisterMediaType register given datastructure
func (dss *MediaTypes) RegisterMediaType(mt protocol.MediaType) {
	if mt.MediaType() == "" {
		panic("Mediatype doesn't has an MediaType. Can't register empty mediatype.")
	}

	poolByMediaType[mt.MediaType()] = mt
	poolByID[mt.ID()] = mt
	var fe = mt.FileExtension()
	if fe != "" {
		poolByFileExtension[fe] = mt
	}
}

func (dss *MediaTypes) GetMediaTypeByID(id uint64) protocol.MediaType { return ByID(id) }
func (dss *MediaTypes) GetMediaTypeByFileExtension(ex string) protocol.MediaType {
	return ByFileExtension(ex)
}
func (dss *MediaTypes) GetMediaType(mediaType string) protocol.MediaType {
	return ByMediaType(mediaType)
}

var (
	poolByMediaType     = map[string]protocol.MediaType{}
	poolByID            = map[uint64]protocol.MediaType{}
	poolByFileExtension = map[string]protocol.MediaType{}
)

func ByMediaType(mediaType string) protocol.MediaType { return poolByMediaType[mediaType] }
func ByID(id uint64) protocol.MediaType               { return poolByID[id] }
func ByFileExtension(ex string) protocol.MediaType    { return poolByFileExtension[ex] }
