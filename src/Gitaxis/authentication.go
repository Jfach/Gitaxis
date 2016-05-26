package main

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "os"
  "common"
)

func Authenticate() {
  fmt.Println("[*] Enter Key:")

  var keyInput string
  fmt.Scanln(&keyInput)
  firstLine := common.ReadFirstLine(".key.hkp")
  compareError := bcrypt.CompareHashAndPassword([]byte(firstLine), []byte(keyInput))

  if compareError != nil {
    denied()
  } else {
    granted()
  }
}

func CreateKey() {
  fmt.Println("[*] Enter new key:")
  var keyInput string
  fmt.Scanln(&keyInput)
  generateKey(keyInput)
}

func DestroyKey() {
  common.DeleteFile("Dungeon/.key.hkp")
}

func generateKey(password string) {
  hashKey, hashError := bcrypt.GenerateFromPassword([]byte(password), 11)
  common.ErrorCheck(hashError)

  keyFile, keyError := os.Create("Dungeon/.key.hkp")
  common.ErrorCheck(keyError)
  defer keyFile.Close()

  keyFile.Write([]byte(hashKey))
  keyFile.Sync()
}

//TODO
// Query other nodes for decryption key, allow user
// to enter instructions
func granted() {
    fmt.Println("ACCESS GRANTED")
}

// End Program, check for amount of past attempts.
// If the past attempts are over 10 wipe everything
// If not, increment past attempts
//TODO
func denied() {
    fmt.Println("\nACCESS DENIED")
}
