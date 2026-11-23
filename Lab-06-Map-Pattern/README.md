# Examples of Map Pattern

Map is the simplest pattern.  It is embarrisingly parallel with no need for any data sharing between threads.

Two examples are included here with code showing how to get timing results demonstrating speedup.

## SAXPY

This is a standard algorithm used a lot in simulation.
It  stands for Single-precision y := alpha * x + y where alpha is a scalar and x and y are n-vectors. SAXPY is a combination of scalar multiplication and vector addition, and it’s very simple: it takes as input two vectors of 32-bit floats X and Y with N elements each, and a scalar value A. It multiplies each element X[i] by A and adds the result to Y[i]. There are Double and Complex data type versions as well - DAxpy, Caxpy, Zaxpy.

## Mandelbroit Set

The Mandelbrot set is a two-dimensional set with a relatively simple definition that exhibits great complexity, especially as it is magnified. It is popular for its aesthetic appeal and fractal structures. The set is defined in the complex plane as the complex numbers c c for which the function  f_{c}(z)=z^{2}+c  does not diverge to infinity when iterated starting at z = 0 z=0, i.e., for which the sequence f_{c}(0)}, f_{c}(f_{c}(0))}, etc., remains bounded in absolute value.

This set was first defined and drawn by Robert W. Brooks and Peter Matelski in 1978, as part of a study of Kleinian groups.[3] Afterwards, in 1980, Benoit Mandelbrot obtained high-quality visualizations of the set while working at IBM's Thomas J. Watson Research Center in Yorktown Heights, New York. [label](https://en.wikipedia.org/wiki/Mandelbrot_set)

## Timing Results

The programs were tested using one OpenMP thread and four OpenMP threads.

### SAXPY

| Threads | Execution Time |
|---|---:|
| 1 | 0.002925 seconds |
| 4 | 0.002984 seconds |

The four-thread version was slightly slower in this run. This is likely because the SAXPY operation completes very quickly, meaning that the overhead involved in managing multiple threads can outweigh the benefit of parallel execution.

### Mandelbrot

| Threads | Execution Time |
|---|---:|
| 1 | 0.043097 seconds |
| 4 | 0.052917 seconds |

The four-thread Mandelbrot version was also slower in this run. Parallel execution does not always guarantee improved performance, particularly when the workload is relatively small or when thread-management overhead becomes significant.

The results may also vary between runs depending on processor load and scheduling.

## Conclusion

Both examples demonstrate the OpenMP Map Pattern, where independent iterations of a computation can be distributed across multiple threads.

Although the tested four-thread versions did not produce a speed-up on this machine, the programs successfully demonstrate how OpenMP can parallelise independent work using `parallel for`.
