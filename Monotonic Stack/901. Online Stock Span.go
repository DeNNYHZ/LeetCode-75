package Monotonic_Stack

type StockSpanner struct {
	stack  []int
	prices []int
	index  int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack:  make([]int, 0),
		prices: make([]int, 0),
		index:  0,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.prices = append(this.prices, price)
	for len(this.stack) > 0 && this.prices[this.stack[len(this.stack)-1]] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	result := 0
	if len(this.stack) > 0 {
		result = this.index - this.stack[len(this.stack)-1]
	} else {
		result = this.index + 1
	}
	this.stack = append(this.stack, this.index)
	this.index++
	return result
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
