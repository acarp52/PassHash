package main

import (
    "crypto/sha512"
    "encoding/base64"
)

// Computes the SHA512 hash of the user's password, then returns it
// as a base64 encoded tring
func hashPassword(passwd string) string{
    // Returns new hash.Hash that computes SHA-512 checksum
    hash := sha512.New()

    // Writes hash of password, sliced into an array of bytes
    hash.Write([]byte(passwd))

    // Appends hash of sliced bytes together
    hashSum := hash.Sum(nil)

    // Converts byte array to a base64 encoded string
    strHash := base64.StdEncoding.EncodeToString(hashSum)

    return strHash
}
