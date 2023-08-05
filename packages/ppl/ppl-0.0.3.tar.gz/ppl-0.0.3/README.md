<h1 align="center">pb</h1>

<p align="center">
<img src="https://i.imgur.com/UGgcdf1.gif">
<br>
The progress bar library that started with the idea of looking pretty
</p>

## Installation

```
pip3 install mpb
```
> By default the bar length is the full width of the terminal window

## How to use

### Simple usage

```python
import time
from mpb import pb

for i in pb(range(100)):
    time.sleep(0.1)
```

### Show task name along with the progress bar

```python
import time
import random
from mpb import pb

total = 120
tasks = [
    'Make paintball', 'Find dragons', 'Code in python', 'Take out the trash',
    'Fill up water bottles for trip'
]
for task in tasks:
    i = 0
    for i in pb(range(total), task=task):
        sleep_time = [.05, .04, .03, .02, .01][random.randint(0, 4)]
        time.sleep(sleep_time)  # emulating long-playing task
```

### Custom bar length

```python
import time
from mpb import pb

for i in pb(range(100), bar_len=20):
    time.sleep(0.1)
```
