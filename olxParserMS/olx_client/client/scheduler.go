package client

type Scheduler struct {
	List map[int]int64 // [order id] unix time
}
