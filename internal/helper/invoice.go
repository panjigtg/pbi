package helper

import (
	"fmt"
	"time"
)

func GenerateInvoice() string {
	return fmt.Sprintf("INV-%d", time.Now().Unix())
}