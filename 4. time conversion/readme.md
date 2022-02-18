# Time conversion

Given a time in -hour AM/PM format, convert it to military (24-hour) time.
Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

### example
> s= '12:01:00PM'

> output -> 12:01:00

### example 2
> s= '12:01:00AM'

> output -> 00:01:00

## function description
Complete the timeConversion function in the editor below. It should return a new string representing the input time in 24 hour format.