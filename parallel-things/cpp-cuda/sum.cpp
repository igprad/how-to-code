#include <chrono>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>

#define MAX 1000000

int main() {
  std::ifstream csv("dataset.csv");
  double* first_value = (double*)malloc(sizeof(double) * MAX);
  double* second_value = (double*)malloc(sizeof(double) * MAX);
  double* third_value = (double*)malloc(sizeof(double) * MAX);
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
        if (rowIdx == 1) first_value[idx] = std::stod(number, &sz);
        if (rowIdx == 2) second_value[idx] = std::stod(number, &sz);
        if (rowIdx == 3) third_value[idx] = std::stod(number, &sz);
      } catch (const std::exception& e) {
        std::cerr << "error: " << e.what() << std::endl;
      }
      rowIdx++;
    }
    idx++;
  }

  /*
   * Will calculate the time spent to sum all values for each rows.
   * Then will show the diffs to show how long the summarizing process takes.
   * */
  auto start = std::chrono::high_resolution_clock::now();
  for (int i = 0; i < MAX; i++) {
    result[i] = first_value[i] + second_value[i] + third_value[i];
  }
  auto end = std::chrono::high_resolution_clock::now();
  double diff_time =
      std::chrono::duration<double, std::milli>(end - start).count();

  double total = 0;
  for (int i = 0; i < MAX; i++) {
    total += result[i];
  }
  std::cout << "To verify all sum was correct, will sum all result elements "
               "with total: "
            << total << std::endl;
  std::cout << "Processing time for each row summarization: " << diff_time
            << " ms." << std::endl;

  free(first_value);
  free(second_value);
  free(third_value);
  free(result);
  return 0;
}
