```python
def func(stop_event):
	while True:
		if stop_event.is_set():
			break
		# do something

thread = StoppableThread(target=func)
thread.start()
# ...
thread.stop()
# thread.stopped() == True

thread = ReStartableThread(target=func, event_name='stop_event') # can specify the parameter's name
thread.start()
# ...
thread.stop()
# ...
thread.start()
# ...
thread.stop()
# ...
```

