# TrueAes
A simple go file containing function to Encrypt and Decrypt in AES
# Requirements
  - go
  - [pkcs7pad]("https://github.com/zenazn/pkcs7pad")
# Why this lib?
This lib assist the process of split the target to encrypt in block of 16 bytes, to match the cyper size of AES128, so you can encrypt everything in just one function.
With some modification you can even achieve AES192 or AES256, it's a matter of changing the forloop condition and the keysize
