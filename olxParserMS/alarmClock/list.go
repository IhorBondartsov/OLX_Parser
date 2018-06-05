package alarmClock

import "sort"

// ListItem - is struct which save:
// Item - object which we keep
// Time - "wake up time" time when this obj must be send. Using as sorting value. Must be on UNIX Time (second)
// To - its channel, where we must return Item
type ListItem struct {
	Item interface{}
	Time int64
	To   chan interface{}
}

// List - keep items, have Sort() and the sorted addition
type List struct {
	List []ListItem
}

func NewList() AlarmClockList {
	return &List{
		List: make([]ListItem, 0, defaultCapacityForList),
	}
}

// Sort - sorted by time
func (l *List) Sort() {
	if len(l.List) == 0 {
		return
	}
	sort.Slice(l.List, func(i, j int) bool {
		return l.List[i].Time < l.List[j].Time
	})
}

// SortedAddByTime - added element, and keep slice sorted
func (l *List) SortedAddByTime(qi ListItem) {
	for k, v := range l.List {
		if qi.Time < v.Time {
			l.List = append(l.List[:k], append([]ListItem{qi}, l.List[k:]...)...)
			return
		}
	}
	l.List = append(l.List, qi)
}

// Add - added element to the slice end
func (l *List) Add(qi ListItem) {
	l.List = append(l.List, qi)
}

// DeleteByIndex - delete element from slice by index
func (l *List) DeleteByIndex(number int) {
	l.List = append(l.List[:number], l.List[number+1:]...)
}

// Delete - find element in slice end delete it
func (l *List) Delete(li ListItem) {
	for k, v := range l.List {
		if li == v {
			l.List = append(l.List[:k], l.List[k+1:]...)
			return
		}
	}
}

// GetList - return list with all item
func (l *List) GetList() []ListItem {
	return l.List
}
