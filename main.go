package main

import (
    "log"

    "github.com/stellar/go/keypair"
)

func main() {
    log.Println("Starting key generation...")

    kp, err := keypair.Random()
    if err != nil {
        log.Panic(err)
        return
    }

    log.Println(kp.Address())
    log.Println(kp.Seed())
}
