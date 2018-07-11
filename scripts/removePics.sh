#!/bin/sh

PATH=$SQUEEZEPATH/training-images

for d in $PATH; do
	if [ -d "$d" ]; then
		echo "$d"
	fi
done
