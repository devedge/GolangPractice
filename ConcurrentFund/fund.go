package funding

type Fund struct {
  // balance is private since it's in lowercase
  balance int
}

func NewFund(initialBalance int) *Fund {
  // Return a pointer to a new struct
  return &Fund{
    balance: initialBalance
  }
}

// Method starts with a receiver, eg., a Fund pointer
func (f *Fund) Balance() int {
  return f.balance
}

func (f *Fund) Withdraw(amount int)  {
  f.balance -= amount
}
