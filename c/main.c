#include <stdio.h>

#include "array.h"
#include "sort.h"

int main() {
  Array arr;
  initArray(&arr);
  populate(&arr, 15);

  printArray(&arr, "initial array:", 1);
  insertion(arr);
  printArray(&arr, "insertion sort:", 1);

  return 0;
}

