package main
import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

var cardNumber struct {
        Number string `json:"number"`
    }

isValid, err := luhnAlgorithm(cardNumber.Number)
    if err != nil {
        http.Error(writer, "Error creating response", http.StatusInternalServerError)
        return
    }

func algorithm(creditNum string) bool
  sumOfTotal := 0
  isSecondDigit := false

    for i := len(creditNum) - 1; i >= 0; i-- {
        digit := int(cardNumber[i] - '0')

        if isSecondDigit {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sumOfTotal += digit
        isSecondDigit = !isSecondDigit
    }
    return sumOfTotal%10 == 0
}
