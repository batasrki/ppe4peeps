#/bin/bash
brew services start zookeeper && brew services start kafka

# Create orders-received topic
kafka-topics --zookeeper 127.0.0.1:2181 \
--create --topic orders-received --partitions 2 --replication-factor 1

# Set orders-received topic retention to 3 days
kafka-configs --alter --add-config retention.ms=259200000 \
--bootstrap-server 127.0.0.1:9092 --topic orders-received

# Create orders-confirmed topic
kafka-topics --zookeeper 127.0.0.1:2181 \
--create --topic orders-confirmed --partitions 2 --replication-factor 1

# Create orders-picked topic
kafka-topics --zookeeper 127.0.0.1:2181 \
--create --topic orders-picked --partitions 2 --replication-factor 1

# Create customer-notification topic
kafka-topics --zookeeper 127.0.0.1:2181 \
--create --topic customer-notifications --partitions 2 --replication-factor 1

# Create consumer-errors topic
kafka-topics --zookeeper 127.0.0.1:2181 \
--create --topic consumer-errors --partitions 2 --replication-factor 1

