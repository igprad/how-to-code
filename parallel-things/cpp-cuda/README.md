# Example Simple CUDA
## What
This is an example utilize nvidia to processing all values in the csv async.\
What is CUDA? Here: https://en.wikipedia.org/wiki/CUDA \
In this example program, will compare the sumarization result between `cpp` and `cuda`.

## Pre-req
Installed `nvcc` and `cuda utilities` https://developer.nvidia.com/cuda-downloads, also gcc.\
If you're using windows. You can utilize Visual Studio.\
Example dataset provided here (1 mil rows that containt 3 columns of float numbers)

## How-to
If you're using native compiler `nvcc` and `gcc` like me. You can compile and run it. Using default: \
`cpp` -> `gcc sum.cpp -o cpp.out` \
`cuda` -> `nvcc sum.cu -o cuda.out` \
then run the shells to getting the result on your local machine.

## Result
The result varies, but almost consistent on my local machine for the difference.\
The `cuda` program got almost +- 15 times faster.\
When this README updated, the `cpp` got result +- `6ms` and the `cuda` got `0.5ms`\
The sumarization result also correct.

## My Spec
Running on NVIDIA-1070Ti x i7-4ish :(
