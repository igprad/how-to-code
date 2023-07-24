#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <chrono>

#include <cuda_runtime.h>
#include <cuda_runtime_api.h>
#include <device_launch_parameters.h>

#define MAX 1000000


__device__ double sum(double first_val, double second_val, double third_val) {
    return first_val + second_val + third_val;
}

__global__ void sum_csv_data_with_cuda(double *result, double *first_row, double *second_row, double *third_row, int max) {
    int offset_x = threadIdx.x;
    int offset_y = blockIdx.x;
    int offset = offset_y + offset_x * max;
    if (offset < MAX) result[offset] = sum(first_row[offset], second_row[offset], third_row[offset]);
}

int main() {
    std::ifstream csv ("dataset.csv");
    double* first_value = (double*) malloc(sizeof(double) * MAX);
    double* second_value = (double*) malloc(sizeof(double) * MAX);
    double* third_value = (double*) malloc(sizeof(double) * MAX);
    double* result = (double*)malloc(sizeof(double) * MAX);

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
        int rowIdx = 0;
        while (std::getline(lineToInputBuffered, number, ',')) {
            try {
                if (rowIdx == 1)
                   first_value[idx] = std::stod(number, &sz);
                if (rowIdx == 2)
                   second_value[idx] = std::stod(number, &sz);
                if (rowIdx == 3)
                   third_value[idx] = std::stod(number, &sz);
            } catch (const std::exception& e) {
                std::cerr << "error: "<< e.what() << std::endl;
            }
            rowIdx++;
        }
        idx++;
    }

    /*
     * Will calculate the time spent to sum all values for each rows.
     * This section will be quite different with `sum.cpp`,
     * since will use cuda api to do the parallel tasks.
     * */
    double* first_dev = (double*) malloc(sizeof(double) * MAX);
    cudaMalloc(&first_dev, sizeof(double) * MAX);
    cudaMemcpy(first_dev, first_value, sizeof(double) * MAX, cudaMemcpyHostToDevice);

    double* second_dev = (double*) malloc(sizeof(double) * MAX);
    cudaMalloc(&second_dev, sizeof(double) * MAX);
    cudaMemcpy(second_dev, second_value, sizeof(double) * MAX, cudaMemcpyHostToDevice);

    double* third_dev = (double*) malloc(sizeof(double) * MAX);
    cudaMalloc(&third_dev, sizeof(double) * MAX);
    cudaMemcpy(third_dev, third_value, sizeof(double) * MAX, cudaMemcpyHostToDevice);
    
    double* result_dev = (double*) malloc(sizeof(double) * MAX);
    cudaMalloc(&result_dev, sizeof(double) * MAX);

    cudaEvent_t start, stop;
    float time = 0;
    cudaEventCreate(&start);
    cudaEventCreate(&stop);

    cudaEventRecord(start, 0);
    sum_csv_data_with_cuda<<<1000, 1000>>>(result_dev, first_dev, second_dev, third_dev, 1000); 
    cudaEventRecord(stop, 0);
    cudaEventSynchronize(stop);
    cudaEventElapsedTime(&time, start, stop);

    cudaMemcpy(result, result_dev, sizeof(double) * MAX, cudaMemcpyDeviceToHost);

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
    for(int i=0; i < MAX; i++) {
        total+=result[i];
    }
    std::cout<<"Total sum of all rows = " << total << std::endl;
    std::cout<<"With elapsed time while using cuda threads (ms) = " << time << std::endl;

    free(first_value);
    free(second_value);
    free(third_value);
    free(result);
    return 0;
}
