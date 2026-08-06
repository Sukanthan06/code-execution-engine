FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    python3 \
    nodejs \
    gcc \
    g++ \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m -s /bin/bash runner

WORKDIR  /app

USER runner

CMD ["bash"]