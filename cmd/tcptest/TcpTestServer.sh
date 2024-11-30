#!/bin/bash

num_packages=12
payload_size_mb=10
server="localhost"
port=3333
num_parallel=4

send_payload() {
  package_num=$1
  server=$2
  port=$3

  payload_file="100mb_payload_$package_num.bin"

  dd if=/dev/urandom of=$payload_file bs=1M count=$payload_size_mb

  cat $payload_file | nc $server $port

  rm $payload_file
}

export -f send_payload

for ((i = 1; i <= $num_packages; i++)); do
  send_payload $i $server $port &

  while [ $(jobs | wc -l) -ge $num_parallel ]; do
    sleep 0.1
  done
done

wait

echo "All $num_packages packages of 100 MB sent!"
