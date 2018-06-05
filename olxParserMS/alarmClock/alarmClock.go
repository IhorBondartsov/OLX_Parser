package alarmClock

import (
	"time"
	"reflect"
	"unsafe"

	"github.com/powerman/narada-go/narada"
)

var log = narada.NewLog("alarmClock: ")

const (
	sizeBufferForAlarmClockChannel = 20
	defaultCapacityForList         = 16
)

type AlarmClockList interface {
	Sort()
	Add(ListItem)
	SortedAddByTime(ListItem)
	Delete(ListItem)
	DeleteByIndex(int)
	GetList() []ListItem
}

type AlarmClock struct {
	List     AlarmClockList
	addChan  chan ListItem
	stopChan chan struct{}
	ownTimer specialTimer
}

// specialTimer - its struct have timerIsOccupy.
// timerIsOccupy - if timer is working value is true, if timer is stop or expired value is false
type specialTimer struct {
	timer         *time.Timer
	timerIsOccupy bool
}

func NewAlarmClock() *AlarmClock {
	return &AlarmClock{
		List:    NewList(),
		addChan: make(chan ListItem, sizeBufferForAlarmClockChannel),
		stopChan:make(chan struct{}),
		ownTimer: specialTimer{
			timerIsOccupy: false,
		},
	}
}

func (ac *AlarmClock) Start() {
	log.DEBUG("AlarmClock started...")
	go ac.work()
	ac.ownTimer.timer = time.NewTimer(0)
}

// wake - check channel on closed, if not closed send Item
func (ac *AlarmClock) wake(li ListItem) {
	if IsChanClosed(li.To) {
		log.ERR("WakeError. Chanel for item was closed")
		return
	}
	li.To <- li.Item
}

// work - must work as goroutine.
// When into addChan come a value we check timer occupation if timer is free we switch on timer.
// When timer is ending we check all element in list
func (ac *AlarmClock) work() {
	log.DEBUG("work started...")
	for {
		select {
		case item := <-ac.addChan:
			ac.List.SortedAddByTime(item)

			if !ac.ownTimer.timerIsOccupy {
				sleepTime := time.Duration(item.Time - time.Now().Unix())
				ac.ownTimer.timer = time.NewTimer(time.Second * time.Duration(sleepTime))
				ac.ownTimer.timerIsOccupy = true
			}
		case <-ac.ownTimer.timer.C:
			ac.checkList()

		case <-ac.stopChan:
			return
		}
	}
}

// checkList - if element in list is expired we send this element into channel, and delete from array.
// If not expired we create new timer and interrupt execution. After checking list we set timer occupation as false
func (ac *AlarmClock) checkList() {
	for i := 0; i < len(ac.List.GetList()); i++ {

		now := time.Now().Unix()
		item := ac.List.GetList()[i]

		if item.Time > now {
			sleepTime := time.Duration(item.Time - now)
			ac.ownTimer.timer = time.NewTimer(time.Second * sleepTime)
			return
		}

		ac.wake(item)
		ac.List.DeleteByIndex(i)
		i--
	}

	if len(ac.List.GetList()) != 0 {
		ac.ownTimer.timer = time.NewTimer(time.Second * 2)
	} else {
		ac.ownTimer.timerIsOccupy = false
	}
}

func (ac *AlarmClock) GetAddChan() chan ListItem {
	return ac.addChan
}

func (ac *AlarmClock) GetStopChan() chan struct{} {
	return ac.stopChan
}

func IsChanClosed(ch interface{}) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {
		panic("only channels!")
	}

	// get interface value pointer, from cgo_export
	// typedef struct { void *t; void *v; } GoInterface;
	// then get channel real pointer
	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0))),
	))

	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **

	cptr += unsafe.Sizeof(uint(0)) * 2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(cptr)) > 0
}
