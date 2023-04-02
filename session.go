
package stockfighter

import (
	"sync"
	"fmt"
	"encoding/json"
)

type Session struct {
	Cash          int