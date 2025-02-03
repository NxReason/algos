#include <stdio.h>

#include "array.h"
#include "sort.h"
#include "utils.h"

int main() {
  Array arr;
  initArray(&arr);
  populate(&arr, 15);

  Array second = copy(&arr);

  printArray(&arr, "first: ", 1);
  printArray(&second, "second: ", 1);

  insertion(arr);
  insertionDesc(second);

  printArray(&arr, "first asc: ", 1);
  printArray(&second, "second desc: ", 1);

  free(arr.values);
  free(second.values);
  return 0;
}

