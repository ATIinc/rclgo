package test

// #include <malloc.h>
import "C"

// cHeapInUse returns the bytes glibc malloc has handed out and not yet freed,
// summed over all arenas. Chunks served by mmap are not counted.
func cHeapInUse() int64 {
	return int64(C.mallinfo2().uordblks)
}
