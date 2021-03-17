#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "foo.h"

int main() {
  char buf[10];
  GoString go_str;

  for (int i = 0; i < 10; i++) {
    buf[i] = 'a';
  }
  buf[9] = 0;

  go_str.p = buf;
  go_str.n = 10;
  char *r = bar(go_str, buf);

  printf("return: %s\n", r);
  free(r);
  sleep(3);

  for (int i = 0; i < 10; i++) {
    buf[i] = 'b';
  }
  buf[9] = 0;

  sleep(100000000);
  return 0;
}