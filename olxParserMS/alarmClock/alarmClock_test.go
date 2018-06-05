package alarmClock

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlarmClock_WorkCloseChannel(t *testing.T) {
	a := assert.New(t)

	ac := NewAlarmClock()
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []string{"Orange","Apple","Pinapple","Cucumber"}

	item := ListItem{
		Time: now + int64(2),
		Item: "Orange",
		To:   make(chan interface{}),
	}
	item1 := ListItem{
		Time: now + int64(3),
		Item: "Apple",
		To:   make(chan interface{}),
	}
	item2 := ListItem{
		Time: now + int64(4),
		Item: "Pinapple",
		To:   make(chan interface{}),
	}
	item3 := ListItem{
		Time: now + int64(5),
		Item: "Cucumber",
		To:   make(chan interface{}),
	}
	item4 := ListItem{
		Time: now + int64(6),
		Item: "Melon",
		To:   make(chan interface{}),
	}
	item5 := ListItem{
		Time: now + int64(7),
		Item: "Cherry",
		To:   make(chan interface{}),
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	close(item5.To)
	close(item4.To)

	res := <-item.To
	res1 := <-item1.To
	res2 := <-item2.To
	res3 := <-item3.To
	res4 := <-item4.To
	res5 := <-item5.To

	a.Nil(res4)
	a.Nil(res5)

	time.Sleep(2*time.Second)

	stopChan <- struct{}{}

	result := []string{res.(string),res1.(string),res2.(string),res3.(string)}

	a.Equal(ecpected, result)
}


func TestAlarmClock_Work(t *testing.T) {
	a := assert.New(t)

	ac := NewAlarmClock()
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []string{"Orange","Apple","Pinapple","Cucumber","Melon", "Cherry"}

	item := ListItem{
		Time: now + int64(2),
		Item: "Orange",
		To:   make(chan interface{}),
	}
	item1 := ListItem{
		Time: now + int64(3),
		Item: "Apple",
		To:   make(chan interface{}),
	}
	item2 := ListItem{
		Time: now + int64(4),
		Item: "Pinapple",
		To:   make(chan interface{}),
	}
	item3 := ListItem{
		Time: now + int64(5),
		Item: "Cucumber",
		To:   make(chan interface{}),
	}
	item4 := ListItem{
		Time: now + int64(6),
		Item: "Melon",
		To:   make(chan interface{}),
	}
	item5 := ListItem{
		Time: now + int64(7),
		Item: "Cherry",
		To:   make(chan interface{}),
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	res := <-item.To
	res1 := <-item1.To
	res2 := <-item2.To
	res3 := <-item3.To
	res4 := <-item4.To
	res5 := <-item5.To

	time.Sleep(2*time.Second)

	stopChan <- struct{}{}

	result := []string{res.(string),res1.(string),res2.(string),res3.(string),res4.(string),res5.(string)}

	a.Equal(ecpected, result)
}


func TestAlarmClock_WorkDublicateTime(t *testing.T) {
	a := assert.New(t)

	ac := NewAlarmClock()
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []string{"Orange","Apple","Pinapple","Cucumber","Melon", "Cherry"}

	c := make(chan interface{},1)
	c1 := make(chan interface{},1)
	c2 := make(chan interface{},1)
	c3 := make(chan interface{},1)
	c4 := make(chan interface{},1)
	c5 := make(chan interface{},1)

	item := ListItem{
		Time: now + int64(7),
		Item: "Orange",
		To:   c,
	}
	item1 := ListItem{
		Time: now + int64(7),
		Item: "Apple",
		To:   c1,
	}
	item2 := ListItem{
		Time: now + int64(7),
		Item: "Pinapple",
		To:  c2,
	}
	item3 := ListItem{
		Time: now + int64(4),
		Item: "Cucumber",
		To:   c3,
	}
	item4 := ListItem{
		Time: now + int64(7),
		Item: "Melon",
		To:   c4,
	}
	item5 := ListItem{
		Time: now + int64(7),
		Item: "Cherry",
		To:   c5,
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	res := <-item.To
	res1 := <-item1.To
	res2 := <-item2.To
	res3 := <-item3.To
	res4 := <-item4.To
	res5 := <-item5.To

	time.Sleep(2*time.Second)

	stopChan <- struct{}{}

	result := []string{res.(string),res1.(string),res2.(string),res3.(string),res4.(string),res5.(string)}

	a.Equal(ecpected, result)
}