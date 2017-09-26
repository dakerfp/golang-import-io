// ...
event.Events = syscall.EPOLLIN
event.Fd = int32(fd) // fd is set as the listenter socket
e = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event) // HL
// handle error ...
for {
	nevents, e := syscall.EpollWait(epfd, events[:], -1) // HL
	// handle error ...
	for ev := 0; ev < nevents; ev++ {
		if int(events[ev].Fd) == fd {
			connFd, _, err := syscall.Accept(fd)
			if err != nil {
				// handle error ...
			}
			syscall.SetNonblock(fd, true) // HL
			event.Events = syscall.EPOLLIN | EPOLLET
			event.Fd = int32(connFd)
			if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil { // HL
				os.Exit(1)
			}
		} else {
			go echo(int(events[ev].Fd)) // HL
		}
	}
}