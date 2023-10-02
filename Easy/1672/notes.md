Looking across solutions I can see Goroutines are a fast and commonly used method *point of research here* "A WaitGroup is used to wait for a collection of goroutines to finish executing. The Add(1) method increments the WaitGroup's counter by one, indicating that one more goroutine is being launched.

For each account in the accounts slice, a goroutine is launched that sums up the values of the account and sends the result to the channel ch.

The defer wg.Done() line schedules the call to wg.Done() to be run after the anonymous function has returned. This decrements the WaitGroup's counter by one, indicating that one of the goroutines has finished executing.

The wg.Wait() line blocks until the WaitGroup's counter becomes zero, which means all the goroutines have finished executing.

Then the channel is closed, and the for loop iterates over the channel, comparing the value of accountSum to maxSum, if its greater than maxSum it updates the value of maxSum to accountSum.

Finally, the function returns the maxSum which is the maximum sum of all the account values.

```
func maximumWealth(accounts [][]int) int {
	maxSum := 0
	ch := make(chan int, len(accounts))
	wg := sync.WaitGroup{}

	for _, account := range accounts {
		account := account

		wg.Add(1)
		go func() {
			defer wg.Done()

			accountSum := 0
			for _, val := range account {
				accountSum += val
			}
			ch <- accountSum
		}()
	}

	wg.Wait()
	close(ch)
	for accountSum := range ch {
		if accountSum > maxSum {
			maxSum = accountSum
		}
	}

	return maxSum
}```

"