#!/bin/bash

#executing network
cd /home/naif/Documents/squeezeCNN
rm aligned-images/cache.t7
./squeezecnn --train --align --align_folder_in=./training-images/ --align_folder_out=aligned-images
