#!/bin/sh

PATH=/home/naif/Documents/squeezeCNN/training-images

for d in $PATH; do
	if [ -d "$d" ]; then
		echo "$d"
	fi
done
