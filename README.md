# Study: Web Authentication With Golang

Tutorial link: https://www.udemy.com/course/oauth-authentication/

## Commits
1. Section 1 to 5: Introduction & Basic Authentication
2. Section 6: Exploring HMAC
3. Section 7: Exploring JWT
4. Section 8: Exploring Encryption

## Note
- Hashing 
  - MD5 - don’t use 
  - SHA 
  - Bcrypt 
  - Scrypt 
- Signing 
  - Symmetric Key 
    - HMAC 
    - same key to sign (encrypt) / verify (decrypt)
  - Asymmetric Key 
    - RSA ECDSA - better than RSA; faster; smaller keys 
    - private key to sign (encrypt) / public key to verify (decrypt)
  - JWT 
- Encryption 
  - Symmetric key 
    - AES 
  - Asymmetric Key 
    - RSA