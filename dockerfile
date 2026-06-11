FROM golang:1.26

WORKDIR /app

RUN apt-get update && apt-get install -y \
    git \
    curl \
    vim

CMD ["bash"]