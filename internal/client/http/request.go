package httpclient

import (
	"sync"
	"time"

	"github.com/steinfletcher/apitest"

	apiRunner "api-tests-template/internal/helpers/api-runner"
)

const minRequestInterval = 120 * time.Millisecond

var (
	requestMu     sync.Mutex
	nextRequestAt time.Time
)

// NewRequest создаёт API-запрос с общей паузой между запросами в тестовом процессе.
func NewRequest(token string) *apitest.APITest {
	requestMu.Lock()
	defer requestMu.Unlock()
	if delay := time.Until(nextRequestAt); delay > 0 {
		time.Sleep(delay)
	}
	nextRequestAt = time.Now().Add(minRequestInterval)
	return apiRunner.GetRunner().Auth(token).Create()
}
