/* For license and copyright information please see LEGAL file in repository */

package http

import "strings"

// TransferEncoding return transfer encoding and notify if multiple exist
// To read multiple just call this method in a loop to get multiple became false
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding
func (h *header) TransferEncoding() (te string, multiple bool) {
	var transferEncoding = h.Get(HeaderKeyTransferEncoding)
	var commaIndex int = strings.IndexByte(transferEncoding, Comma)
	if commaIndex == -1 {
		commaIndex = len(transferEncoding)
	} else {
		h.Replace(HeaderKeyTransferEncoding, transferEncoding[commaIndex+1:])
		multiple = true
	}
	te = transferEncoding[:commaIndex]
	return
}

// SetTransferEncoding set transfer encoding.
func (h *header) SetTransferEncoding(te string) {
	h.Set(HeaderKeyTransferEncoding, te)
}
