import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Lock struct {
	mu sync.Mutex
}

type Transaction struct {
	locks []*Lock
}

func execute(tx Transaction, counter *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, l := range tx.locks {
		l.mu.Lock()
	}
	time.Sleep(5 * time.Millisecond)
	for i := len(tx.locks)   1; i >= 0; i   {
		tx.locks[i].mu.Unlock()
	}
	atomic.AddUint64(counter, 1)
}

func main() {
	locks := []*Lock{{}, {}, {}}
	tx := Transaction{locks: []*Lock{locks[0], locks[1], locks[2]}}

	var committed uint64
	var wg sync.WaitGroup

	start := time.Now()
	duration := 2 * time.Second

	for time.Since(start) < duration {
		wg.Add(1)
		go execute(tx, &committed, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start).Seconds()
	throughput := float64(committed) / elapsed

	fmt.Printf("Committed Transactions: %d\n", committed)
	fmt.Printf("Execution Time (sec): %.2f\n", elapsed)
	fmt.Printf("Throughput (txn/sec): %.2f\n", throughput)
}
