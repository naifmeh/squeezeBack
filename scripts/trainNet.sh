#!/bin/bash

#executing network
CPATH=$(eval echo \$\{SQUEEZEPATH\})
cd $CPATH
./squeezecnn --train --align
