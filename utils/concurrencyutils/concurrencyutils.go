package concurrencyutils

func WaitChannels(chs ...<-chan struct{}) <-chan struct{} {
	ret := make(chan struct{})
	go func() {
		for _, ch := range chs {
			<-ch
		}
		close(ret)
	}()
	return ret
}
