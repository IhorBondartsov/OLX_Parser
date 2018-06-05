package alarmClock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Sort(t *testing.T) {
	a := assert.New(t)

	expectedItemTime := []int64{1, 2, 3, 4, 5}

	qi1 := ListItem{
		Time: 1,
		Item: 5,
	}
	qi2 := ListItem{
		Time: 2,
		Item: 4,
	}
	qi3 := ListItem{
		Time: 3,
		Item: 3,
	}
	qi4 := ListItem{
		Time: 4,
		Item: 2,
	}
	qi5 := ListItem{
		Time: 5,
		Item: 1,
	}

	queue := List{
		List: []ListItem{
			qi2, qi4, qi1, qi5, qi3,
		},
	}

	queue.Sort()

	for k, v := range queue.List {
		a.Equal(expectedItemTime[k], v.Time)
	}
}

func TestQueue_Add(t *testing.T) {
	a := assert.New(t)

	expectedItemTime := []int64{1, 2, 3, 5, 4}

	qi1 := ListItem{
		Time: 1,
		Item: 5,
	}
	qi2 := ListItem{
		Time: 2,
		Item: 4,
	}
	qi3 := ListItem{
		Time: 3,
		Item: 3,
	}
	qi4 := ListItem{
		Time: 4,
		Item: 2,
	}
	qi5 := ListItem{
		Time: 5,
		Item: 1,
	}

	queue := List{}
	queue.Add(qi1)
	queue.Add(qi2)
	queue.Add(qi3)
	queue.Add(qi5)
	queue.Add(qi4)

	for k, v := range queue.List {
		a.Equal(expectedItemTime[k], v.Time)
	}
}

func TestQueue_SortedAddByTime(t *testing.T) {
	a := assert.New(t)

	qi1 := ListItem{
		Time: 1,
		Item: 5,
	}
	qi2 := ListItem{
		Time: 2,
		Item: 4,
	}
	qi3 := ListItem{
		Time: 3,
		Item: 3,
	}
	qi4 := ListItem{
		Time: 4,
		Item: 2,
	}
	qi5 := ListItem{
		Time: 5,
		Item: 1,
	}
	queue := List{}
	queue.Add(qi1)
	queue.Add(qi3)
	queue.Add(qi5)

	queue.SortedAddByTime(qi4)
	queue.SortedAddByTime(qi2)

	a.Equal(qi2, queue.List[1])
	a.Equal(qi4, queue.List[3])
}

func TestQueue_SortedAddByTime2(t *testing.T) {
	a := assert.New(t)

	queue := List{}
	time := []int64{3, 7, 8, 2, 4, 6, 5, 1, 9, 0} // when we add to alarmClock, last elem must be first

	for _, v := range time {
		queue.SortedAddByTime(ListItem{Time: v})
	}

	for i := 0; i < 9; i++ {
		a.Equal(int64(i), queue.List[i].Time)
	}
}

func TestQueue_SortedAddMatchTime(t *testing.T) {
	a := assert.New(t)

	queue := List{}
	time := []int64{3, 7, 8, 1, 4, 4, 5, 1, 9, 1} // when we add to alarmClock, last elem must be first

	expected := []int64{1, 1, 1, 3, 4, 4, 5, 7, 8, 9}

	for _, v := range time {
		queue.SortedAddByTime(ListItem{Time: v})
	}

	for i := 0; i < 9; i++ {
		a.Equal(expected[i], queue.List[i].Time)
	}
}

func TestQueue_Delete(t *testing.T) {
	a := assert.New(t)

	expectedItemTime := []int64{1, 2, 3, 4}

	qi1 := ListItem{
		Time: 1,
		Item: 5,
	}
	qi2 := ListItem{
		Time: 2,
		Item: 4,
	}
	qi3 := ListItem{
		Time: 3,
		Item: 3,
	}
	qi4 := ListItem{
		Time: 4,
		Item: 2,
	}
	qi5 := ListItem{
		Time: 5,
		Item: 1,
	}

	queue := List{}
	queue.Add(qi1)
	queue.Add(qi2)
	queue.Add(qi3)
	queue.Add(qi5)
	queue.Add(qi4)

	queue.Delete(qi5)

	for k, v := range queue.List {
		a.Equal(expectedItemTime[k], v.Time)
	}
}
