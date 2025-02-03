#ifndef ARRAY
#define ARRAY

#include <stdlib.h>

typedef struct arr Array;

struct arr {
  int *values;
  int length;
  int capacity;
};

// init & manage storage
void initArray(Array *arr);
void arrRealloc(Array *arr);
Array copy(Array *src);

// manage data
void push(Array *arr, int value);
int at(Array *arr, int index);
void set(Array *arr, int index, int value);

// misc
void populate(Array *arr, int count);
void printArray(Array *arr, char *prefix, int showProps);
void highlightItems(Array *arr, int* indexes, int idxCount);
void printDelimeter(Array *arr, int after);

#endif