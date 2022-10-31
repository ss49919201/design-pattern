package main

import "fmt"

// Interface

// subject はobserverに通知するオブジェクトのインターフェース
type subject interface {
	// addObserver はsubjectにobserverを登録する
	addObserver(observer)
	// removeObserver はsubjectからobserverを削除する
	removeObserver(observer)
	// notifyObservers はobserverに通知する
	notifyObservers()
}

// observer は通知を受けるオブジェクトのIFです
type observer interface {
	// notify はsubjectから通知を受ける
	notify()
}

// Implement

type sliceSubject struct {
	observerList []observer
}

func (s *sliceSubject) addObserver(o observer) {
	s.observerList = append(s.observerList, o)
}

func (s *sliceSubject) removeObserver(o observer) {
	for i, v := range s.observerList {
		if v == o {
			s.observerList = append(s.observerList[:i], s.observerList[i+1:]...)
		}
	}
}

func (s *sliceSubject) notifyObservers() {
	for _, v := range s.observerList {
		v.notify()
	}
}

type emailObserver struct {
	title string
}

func (e *emailObserver) notify() {
	fmt.Println("send " + e.title)
}

// Usage

func main() {
	s := &sliceSubject{}

	o1 := &emailObserver{title: "email1"}
	o2 := &emailObserver{title: "email2"}
	s.addObserver(o1)
	s.addObserver(o2)
	s.notifyObservers()

	s.removeObserver(o1)
	o2.title = "changed email2"
	s.notifyObservers()

	// Output:
	// send email1
	// send email2
	// send changed email2
}
