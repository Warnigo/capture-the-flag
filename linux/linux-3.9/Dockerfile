FROM ubuntu:20.04

ARG GO_VERSION=1.20.6

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    sudo \
    vim \
    curl \
    git \
    bash \
    build-essential \
    ca-certificates \
    && curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz \
    && rm -rf /var/lib/apt/lists/*

ENV GOROOT=/usr/local/go
ENV GOPATH=/home/go_user/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN useradd -m -s /bin/bash go_user && \
    echo 'go_user ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers && \
    mkdir -p $GOPATH && chown -R go_user:go_user /home/go_user

USER go_user

WORKDIR /app

COPY --chown=go_user:go_user . .

RUN go build -o main .

CMD ["/bin/bash"]