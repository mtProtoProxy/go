package net

type EventTimer struct {
	hIDx  int
	flags int
	// int (*wakeup)(event_timer_t *et);
	WakeUpTime     float64 // double wakeup_time;
	RealWakeUpTime float64 // double real_wakeup_time;
}
