#include <iostream>

void bubble_short_asc(int* arr, int max_arr) {
    for(int i=0; i < max_arr-1; i++)
    {
        for(int j=i+1; j < max_arr; j++) 
        {
            if (arr[j] <= arr[i]) 
            {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
    }
}

void bubble_short_desc(int* arr, int max_arr) {
    for(int i=0; i < max_arr-1; i++)
    {
        for(int j=i+1; j < max_arr; j++) 
        {
            if (arr[j] >= arr[i]) 
            {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
    }
}

void print_arr_values(int* arr, int max_arr) {
    for(int i=0; i<max_arr; i++)
    {
        std::cout << arr[i] << " ";
    }
    std::cout << std::endl;
}

int main() {
    int random_arr[10] = {4, 3, 213, 23, 883, 23, 8921, 8, 90090, 0};
    
    bubble_short_asc(random_arr, 10);
    print_arr_values(random_arr, 10);

    bubble_short_desc(random_arr, 10);
    print_arr_values(random_arr, 10);
}
