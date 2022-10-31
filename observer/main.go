package main

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
