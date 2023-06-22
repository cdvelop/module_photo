package test

import (
	"testing"
)

func TestDevTools(t *testing.T) {
	// Create a new context and cancel function
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()

	// // Set the options to open the DevTools automatically
	// options := []chromedp.ExecAllocatorOption{
	// 	chromedp.Flag("auto-open-devtools-for-tabs", true),
	// }

	// // Create a new allocator with the options
	// allocator, err := chromedp.NewExecAllocator(ctx, options...)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer allocator.Done()

	// // Create a new context with the allocator
	// ctx, cancel = chromedp.NewContext(allocator)
	// defer cancel()

	// // Open a new browser window
	// err = chromedp.Run(ctx, chromedp.Navigate("chrome://newtab"), chromedp.Sleep(time.Second))
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
