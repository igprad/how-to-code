#include <exception>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>

#include <cuda_runtime.h>
#include <cuda_runtime_api.h>
#include <device_launch_parameters.h>

#define DEFAULT_MAX_ROWS 1000000

__device__ double sum(double first_val, double second_val, double third_val) {
  return first_val + second_val + third_val;
}

__global__ void sum_csv_data_with_cuda(double *result, double *first_col,
                                       double *second_col, double *third_col) {
  int offset_x = threadIdx.x;
  int offset_y = blockIdx.x;
  int offset = offset_x + offset_y * 1000;
  if (offset < DEFAULT_MAX_ROWS)
    result[offset] =
        sum(first_col[offset], second_col[offset], third_col[offset]);
}

int main(int argc, char **argv) {

  std::ifstream csv("dataset.csv");
  double *first_value = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  double *second_value = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  double *third_value = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  double *result = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);

  // Getting the value from dataset
  std::string value;
  int idx = 0;
  bool csv_title_skipped = false;
  while (std::getline(csv, value)) {
    if (!csv_title_skipped) {
      csv_title_skipped = true;
      continue;
    }
    std::string number;
    std::stringstream lineToInputBuffered(value);
    std::string::size_type sz;
    int colIdx = 0;
    while (std::getline(lineToInputBuffered, number, ',')) {
      try {
        if (colIdx == 1)
          first_value[idx] = std::stod(number, &sz);
        if (colIdx == 2)
          second_value[idx] = std::stod(number, &sz);
        if (colIdx == 3)
          third_value[idx] = std::stod(number, &sz);
      } catch (const std::exception &e) {
        std::cerr << "error: " << e.what() << std::endl;
      }
      colIdx++;
    }
    idx++;
  }

  /*
   * Will calculate the time spent to sum all values for each rows.
   * This section will be quite different with `sum.cpp`,
   * since will use cuda api to do the parallel tasks.
   * */
  double *first_dev = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMalloc(&first_dev, sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMemcpy(first_dev, first_value, sizeof(double) * DEFAULT_MAX_ROWS,
             cudaMemcpyHostToDevice);

  double *second_dev = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMalloc(&second_dev, sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMemcpy(second_dev, second_value, sizeof(double) * DEFAULT_MAX_ROWS,
             cudaMemcpyHostToDevice);

  double *third_dev = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMalloc(&third_dev, sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMemcpy(third_dev, third_value, sizeof(double) * DEFAULT_MAX_ROWS,
             cudaMemcpyHostToDevice);

  double *result_dev = (double *)malloc(sizeof(double) * DEFAULT_MAX_ROWS);
  cudaMalloc(&result_dev, sizeof(double) * DEFAULT_MAX_ROWS);

  cudaEvent_t start, stop;
  float time = 0;
  cudaEventCreate(&start);
  cudaEventCreate(&stop);

  cudaEventRecord(start, 0);
  sum_csv_data_with_cuda<<<1000, 1000>>>(result_dev, first_dev, second_dev,
                                         third_dev);
  cudaEventRecord(stop, 0);
  cudaEventSynchronize(stop);
  cudaEventElapsedTime(&time, start, stop);

  cudaMemcpy(result, result_dev, sizeof(double) * DEFAULT_MAX_ROWS,
             cudaMemcpyDeviceToHost);

  cudaFree(first_dev);
  cudaFree(second_dev);
  cudaFree(third_dev);
  cudaFree(result_dev);
  cudaFree(&start);
  cudaFree(&stop);
  /*
   * end section
   */

  double total = 0;
  for (int i = 0; i < DEFAULT_MAX_ROWS; i++) {
    total += result[i];
  }
  std::cout << "Total sum of all rows = " << (int)total << std::endl;
  std::cout << "With elapsed time while using cuda threads (ms) = " << time
            << std::endl;

  free(first_value);
  free(second_value);
  free(third_value);
  free(result);
}
