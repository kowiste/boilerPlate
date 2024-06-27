package asset

import "context"



type AssetsGetFilter struct {
	//filters for get asset list
}

func (as Assets) Get(c context.Context, filter *AssetsGetFilter) (err error) {
	return
}
