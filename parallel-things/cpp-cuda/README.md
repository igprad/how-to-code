# Example Simple CUDA
## What
This is an example utilize nvidia to processing all values in the csv async.\
What is CUDA? Here: https://en.wikipedia.org/wiki/CUDA \
In this example program, will compare the sumarization result between `cpp` and `cuda`.

## Requirements
- Installed C++ compiler.
```
gcc version 14.2.1 20250207 (GCC)

```
- Installed `nvcc` and `cuda utilities` https://developer.nvidia.com/cuda-downloads, also gcc.
If you're using windows. You can utilize Visual Studio.
```
nvcc: NVIDIA (R) Cuda compiler driver
Copyright (c) 2005-2025 NVIDIA Corporation
Built on Wed_Jan_15_19:20:09_PST_2025
Cuda compilation tools, release 12.8, V12.8.61
Build cuda_12.8.r12.8/compiler.35404655_0
```

## Dataset
Example dataset provided here is `dataset.csv` that contains 1 mil rows data with 3 columns of float numbers) \
The example program will sum all the data for each rows first, then later on will sum all the results.

## How to
If you're using native compiler `nvcc` and `g++`, you can compile and run the program with default commands:
### C++ 
```
g++ sum.cpp -o cpp.out && ./cpp.out
```
### Cuda
```
nvcc -Wno-deprecated-gpu-targets sum.cu -o cuda.out && ./cuda.out
```

## Result
The result varies, but almost consistent on my local machine for the difference.\
The `cuda` program should be faster.

## My Result
My result in my local machine (i5-12600K && RTX 3060 (12GB)) run on Arch.
- cpp -> result: 1499355 (int value) with time cost ~ 2.55431 ms.
- cuda -> result: 1499355 (int value) with time cost ~ 0.310688 ms. 

### To be improved
[ ] Use dynamic parameter for kernel thread / block size for the calculation.
[ ] Any other suggestions.
