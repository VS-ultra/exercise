package sort_test

import (
	"math/rand"
	"skillfactory/sort"
	"testing"
)

func generateSlice(n int) []int {
	ar := make([]int, n)
	for j := range ar {
		ar[j] = rand.Intn(1000)
	}
	return ar
}

func BenchmarkBubbleSortRecursive(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 10)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 1000)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 100000)
			b.StopTimer()
		}
	})

}

func BenchmarkInsertionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkBidirectionalSelectionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10)
			b.StartTimer()
			sort.BidirectionalSelectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000)
			b.StartTimer()
			sort.BidirectionalSelectionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000)
			b.StartTimer()
			sort.BidirectionalSelectionSort(ar)
			b.StopTimer()
		}
	})

}

func BenchmarkQuickSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})

}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})
}
func BenchmarkBubbleSortWorstCase(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(10, 10)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 10)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(100, 1000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 1000)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(10000, 100000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 100000)
			b.StopTimer()
		}
	})
}

func generateSliceWorstCase(max, size int) []int {
	ar := make([]int, size)
	for i := 0; i < size; i++ {
		ar[i] = max - i
	}

	return ar
}

func BenchmarkBubbleSortBestCase(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(10, 10)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 10)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(100, 1000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 1000)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(10000, 100000)
			b.StartTimer()
			sort.BubbleSortRecursive(ar, 100000)
			b.StopTimer()
		}
	})
}

func generateSliceBestCase(max, size int) []int {
	ar := make([]int, size)
	for i := 0; i < size; i++ {
		ar[i] = i
	}

	return ar
}
func BenchmarkInsertionSortBestCase(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(10, 10)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(100, 1000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceBestCase(10000, 100000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})
}
func BenchmarkInsertionSortWorstCase(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(10, 10)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(100, 1000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceWorstCase(10000, 100000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})
}
