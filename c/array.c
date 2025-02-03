#include "array.h"

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>

#define DEFAULT_ARR_CAPACITY 10
#define LOWER_VALUE 0
#define UPPER_VALUE 10

void initArray(Array *arr) {
  int *values = (int *)malloc(DEFAULT_ARR_CAPACITY * sizeof(int));
  arr->values = values;
  arr->length = 0;
  arr->capacity = DEFAULT_ARR_CAPACITY;
}

void arrRealloc(Array *arr) {
  int *newStorage = (int*)malloc(arr->capacity * 2 * sizeof(int));
  for (int i = 0; i < arr->length; i++) {
    newStorage[i] = arr->values[i];
  }
  // reset storage & free the old one
  int* oldValues = arr->values;
  arr->values = newStorage;
  free(oldValues);

  // update array props
  arr->capacity = arr->capacity * 2;
}

void push(Array *arr, int value) {
  if (arr->length >= arr->capacity) {
    arrRealloc(arr);
  }
  arr->values[arr->length] = value;
  arr->length = arr->length + 1;
}

int at(Array *arr, int index) {
  return arr->values[index];
}

void set(Array *arr, int index, int value) {
  arr->values[index] = value;
}

void populate(Array *arr, int count) {
  srand(time(NULL));
  int range = UPPER_VALUE - LOWER_VALUE + 1;
  while (count--) {
    int newValue = (rand() % range) + LOWER_VALUE;
    push(arr, newValue);
  }
}

void printArray(Array *arr, char *prefix, int showProps) {
  if (prefix == NULL) {
    prefix = "";
  }

  printf("%s [", prefix);
  for (int i = 0; i < arr->length; i++) {
    printf(" %d", at(arr, i));
  }
  printf(" ]\n");
  if (showProps) {
    printf("length: %d, capacity: %d\n", arr->length, arr->capacity);
  }
}
