#!/bin/bash

declare -a items

count=0
while read line; do
  items[$count]=$line
  (( count++ ))
  total=$(( total + line ))
done < input

function interval {
  greater=0
  end=$(( count - $1 ))
  for (( i=0; i < end; i++ )) do
    [[ ${items[$i + $1]} -gt ${items[$i]} ]] && (( greater++ ))
  done
  echo "$greater"
}

for i in 1 3; do interval $i; done
