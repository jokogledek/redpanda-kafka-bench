#!/usr/bin/env sh
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-16-partitions-100b.yaml >> max-rate-1-topic-16-partitions-100b.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-20-partitions-20p-20c-100b.yaml >> max-rate-1-topic-20-partitions-20p-20c-100b.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-1-partition-1p-1c-1kb.yaml >> max-rate-1-topic-1-partition-1p-1c-1kb.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-10-partitions-10p-10c-1kb.yaml >> max-rate-1-topic-10-partitions-10p-10c-1kb.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-16-partitions-1kb.yaml >> max-rate-1-topic-16-partitions-1kb.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-20-partitions-20p-20c-1kb.yaml >> max-rate-1-topic-20-partitions-20p-20c-1kb.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-1-partition-1p-1c-100b.yaml >> max-rate-1-topic-1-partition-1p-1c-100b.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785
sudo bin/benchmark --drivers driver-kafka/kafka-throughput.yaml workloads/max-rate-1-topic-10-partitions-10p-10c-100b.yaml >> max-rate-1-topic-10-partitions-10p-10c-100b.log
rpk topic delete -r '.*' --brokers 0.0.0.0:33785