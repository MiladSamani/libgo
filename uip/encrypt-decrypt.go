/* For license and copyright information please see LEGAL file in repository */

package uip

import "../crypto"

// Encrypt use in encrypted connection from Apps to Apps!
func Encrypt(packet []byte, cipher crypto.Cipher) (err error) {
	err = cipher.Encrypt(packet[32:])
	return
}

// Decrypt use in encrypted connection from Apps to Apps!
func Decrypt(packet []byte, cipher crypto.Cipher) (err error) {
	// Decrypt packet by encryptionKey & Checksum data in this protocol :
	// We check packet errors with encryption proccess together
	// and needed checksum data will be add to encrypted data. 32 bit checksum in end of Packet
	err = cipher.Decrypt(packet[32:])
	return
}

// EncryptRouting usually use in encrypted connection from OS to UIP Router!
func EncryptRouting(packet []byte, cipher crypto.Cipher) (err error) {
	err = cipher.Encrypt(packet[:31])
	return
}

// DecryptRouting usually use in encrypted connection from OS to UIP Router!
func DecryptRouting(packet []byte, cipher crypto.Cipher) (err error) {
	err = cipher.Decrypt(packet[:31])
	return
}
