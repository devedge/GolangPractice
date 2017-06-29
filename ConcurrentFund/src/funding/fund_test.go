package funding

import (
  "sync"
  "testing"
)

const WORKERS = 10

func BenchmarkFund(b *testing.B)  {
  // Skip N=1
  if b.N < WORKERS {
    return
  }

  // Add as many dollars as we have iterations this run
  fund := NewFund(b.N)

  // Assume that b.N divides cleanly
  dollarsPerFounder := b.N/WORKERS

  // WaitGroup structs don't need to be initialized,
  // their "zero value" is ready to use
  // So, just declare one and then use it
  var wg sync.WaitGroup

  // Burn through them one at a time until they are all gone
  for i := 0; i < WORKERS; i++ {
    // Let the WaitGroup know we're adding a goroutine
    wg.Add(1)

    // Spawn off a founder worker, as a closure
    go func()  {
      // Mark this worker done when the function finishes
      defer wg.Done()

      for i := 0; i < dollarsPerFounder; i++ {
        fund.Withdraw(1)
      }
    }() // call the closure
  }

  // Wait for all the workers to finish
  wg.Wait()

  if fund.Balance() != 0 {
    b.Error("Balance wasn't zero:", fund.Balance())
  }
}

func BenchmarkWithdrawals(b *testing.B) {
    // ...

    server := NewFundServer(b.N)

    // ...

    // Spawn off the workers
    for i := 0; i < WORKERS; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < dollarsPerFounder; i++ {
                server.Commands <- WithdrawCommand{ Amount: 1 }
            }
        }()
    }

    // ...

    balanceResponseChan := make(chan int)
    server.Commands <- BalanceCommand{ Response: balanceResponseChan }
    balance := <- balanceResponseChan

    if balance != 0 {
        b.Error("Balance wasn't zero:", balance)
    }
}
