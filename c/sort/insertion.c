#include "../array.h"

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