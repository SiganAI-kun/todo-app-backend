ARG VARIANT="1.21-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

ARG NODE_VERSION="20"
RUN su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get -y install --no-install-recommends jq default-mysql-client

RUN apt install wget

RUN mkdir -p /etc/apt/keyrings
RUN wget -O - https://packages.adoptium.net/artifactory/api/gpg/key/public | tee /etc/apt/keyrings/adoptium.asc
RUN echo "deb [signed-by=/etc/apt/keyrings/adoptium.asc] https://packages.adoptium.net/artifactory/dev $(awk -F '/^VERSION_CODENAME/{print$2}' /etc/os-release) main" | tee /etc/apt/sources.list.d/adoptium.list

WORKDIR /workspaces
COPY go.mod .
RUN go mod download

# RUN export DEBIAN_FRONTEND=noninteractive
# RUN apt-get -y install --no-install-recommends temurin-8-jdk graphviz fonts-noto-cjk