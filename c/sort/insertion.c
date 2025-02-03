#include "../sort.h"

void insertion(Array arr) {
  for (int i = 1; i < arr.length; i++) {
    int cmp = at(&arr, i);
    int j = i - 1;
    while (j >= 0 && at(&arr, j) > cmp) {
      set(&arr, j + 1, at(&arr, j));
      j = j - 1;
    }
    set(&arr, j + 1, cmp);
  }
}

void insertionDesc(Array arr) {
  for (int i = 1; i < arr.length; i++) {
    int cmp = at(&arr, i);
    int j = i - 1;
    while (j >= 0 && at(&arr, j) < cmp) { // only change comparison sign
      set(&arr, j + 1, at(&arr, j));
      j = j - 1;
    }
    set(&arr, j + 1, cmp);
  }
}

void insertionLog(Array arr) {
  printArray(&arr, "initial array:             ", 0);
  
  int *hlIndexes = makeIndexesArray(1, 1);
  printf("start from the second item: ");
  highlightItems(&arr, hlIndexes, 1);
  free(hlIndexes);

  for (int i = 1; i < arr.length; i++) {
    int cmp = at(&arr, i);
    printf("compared item: %d\n", cmp);
    printf("subarray to compare: ");
    int j = i - 1;
    printDelimeter(&arr, j);
    while (j >= 0 && at(&arr, j) > cmp) {
      printf("shift (> %d): ", cmp);
      hlIndexes = makeIndexesArray(1, j, j + 1);
      highlightItems(&arr, hlIndexes, 2);
      free(hlIndexes);

      set(&arr, j + 1, at(&arr, j));
      j = j - 1;
    }
    printf("insert (%d at): ", cmp);
    hlIndexes = makeIndexesArray(1, j + 1);
    highlightItems(&arr, hlIndexes, 1);
    free(hlIndexes);
    set(&arr, j + 1, cmp);
    printArray(&arr, "current step array: ", 0);
  }

}