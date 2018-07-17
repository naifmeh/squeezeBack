#!/bin/bash
JARPATH=/home/naif/Documents/squeezeKafka/squeezekafka.jar
HOST=localhost:9092
CLASSIFIER=$SQUEEZEPATH/generated-embeddings/classifier.pkl
#executing the classifier producer
 java -jar $JARPATH -type classifier-producer -brokers $HOST -file $CLASSIFIER
