#!/bin/sh

LIST="spectral-norm-time \
      spectral-norm-parallel \
      spectral-orig"

for i in $LIST; do
  echo > ./$i.time
  for n in 2000 3000 5500 11000; do
    for ncpu in $(seq 1 16); do
      for iters in $(seq 1 30); do
        echo "Spectral norm '$i' ($n) ($ncpu) ($iters of 30)"
        ./$i/$i -ncpu $ncpu -n $n >> ./$i.time
      done
    done
  done
done
