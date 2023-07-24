# Example Simple CUDA
## What
This is an example utilize nvidia to processing all values in the csv async.\
What is CUDA? Here: https://en.wikipedia.org/wiki/CUDA \
In this example program, will compare the sumarization result between `cpp` and `cuda`.

## Pre-req
Installed nvcc and cuda utilities, also gcc.\
If you're using windows. You can utilize Visual Studio.\
Example dataset provided here (1 mil rows that containt 3 columns of float number)\

## Result
The result varies, but almost consistent on my local machine for the difference.\
The `cuda` program got almost +- 15 times faster.\
When this README updated, the `cpp` got result +- `6ms` and the `cuda` got `0.5ms`\
The sumarization result also correct.
