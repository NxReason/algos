#include "utils.h"

// count must be equal to the number of variadic params provided
int* makeIndexesArray(int count, ...) {
  va_list ap; 
	va_start(ap, count);
  int *indexes = (int*)malloc(count * sizeof(int));

  for (int i = 0; i < count; i++) {
    *(indexes + i) = va_arg(ap, int);;
  }

	va_end(ap);
  return indexes;
}