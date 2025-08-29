# Advanced Topics: Garbage Collection (GC)

Go provides automatic memory management through a garbage collector. This frees developers from manually allocating and deallocating memory, which is a common source of bugs in languages like C.

## How the Go GC Works

Go's garbage collector has evolved significantly over the years. The current implementation (since Go 1.8) is a **concurrent, tri-color, mark-and-sweep garbage collector**.

**Key Characteristics:**
- **Concurrent:** The GC runs concurrently with the main program, minimizing the "stop-the-world" (STW) pauses where the application's execution is halted.
- **Tri-color Algorithm:** This is an abstraction used to track which objects are reachable. Objects are conceptually painted one of three colors:
    - **White:** An object that is a candidate for garbage collection. Initially, all objects are white.
    - **Gray:** An object that is reachable from the roots (stacks, globals) but whose pointers have not yet been scanned. The GC maintains a list of gray objects.
    - **Black:** An object that is reachable and whose pointers have all been scanned. Black objects will not be collected.
- **Mark and Sweep:** The process has two main phases:
    1.  **Mark Phase:** The GC starts from the "roots" (global variables and variables on the stacks of active goroutines) and traverses the object graph. It finds all reachable objects and marks them as "live" (coloring them black). This phase is done concurrently.
    2.  **Sweep Phase:** The GC scans all objects in the heap. Any object that is still "white" is unreachable and its memory can be reclaimed.

## Write Barrier

A crucial component of a concurrent GC is the **write barrier**. When the GC is running, the main program is still executing and can change the object graph (e.g., `A.ptr = B`).

The write barrier is a small amount of code inserted by the compiler on every pointer write. It notifies the garbage collector about the change. Specifically, if the program writes a pointer to a white object into a black object, the write barrier will re-color the white object to gray, ensuring it doesn't get collected incorrectly. This is what allows the marking phase to happen concurrently with the main program.

## Minimizing GC Pauses

While Go's GC is very fast, with STW pauses often in the sub-millisecond range, it's not free. A lead engineer should be aware of how to write code that is GC-friendly.

**Strategies:**
1.  **Reduce Allocations:** The less memory you allocate, the less work the GC has to do.
    - Use object pools (`sync.Pool`) for short-lived, frequently allocated objects.
    - Avoid unnecessary allocations in tight loops.
    - Prefer updating a struct/slice in place over creating a new one, where possible.
2.  **Use Pointers Wisely:** A large number of pointers can increase the amount of work the GC has to do during the mark phase. For very large, simple data structures (e.g., a large slice of simple structs), it can sometimes be more efficient to store the values directly rather than pointers to them.
3.  **Understand Escape Analysis:** As covered in the "Pointers" section, the compiler uses escape analysis to allocate variables on the stack where possible. Stack memory is not managed by the GC, so this is a key optimization. Writing code that allows the compiler to use the stack more effectively can reduce GC pressure.

## Tuning the GC

The `GOGC` environment variable can be used to tune the garbage collector. It's a percentage that controls when a new GC cycle is triggered.
- `GOGC=100` (the default): A new GC cycle is triggered when the heap size has doubled since the previous cycle.
- `GOGC=200`: Trigger when the heap size has tripled.
- `GOGC=off`: Disable the GC entirely.

Changing `GOGC` is a trade-off: a higher value will run the GC less often, reducing its CPU cost, but will increase the program's memory usage. It should only be changed after careful benchmarking and profiling. 