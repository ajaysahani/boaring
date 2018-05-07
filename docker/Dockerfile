FROM debian:jessie

ARG git_branch
RUN if [ -z "$git_branch" ]; then echo "ERROR: git_branch NOT SET. Usage: docker build . --build-arg git_branch=master"; exit 1; else : ; fi

# get golang 1.8.1
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
		libev4 libev-dev \
		gettext-base \
		wget \
		vim \
		tcpdump \
		netcat \
		python-pip \
		git-all \
	&& rm -rf /var/lib/apt/lists/*

RUN pip install cqlsh

ENV GOLANG_VERSION 1.10.1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 72d820dec546752e5a8303b33b009079c15c2390ce76d67cf514991646c6127b

RUN set -eux; \
    wget -O golang.tar.gz "$GOLANG_DOWNLOAD_URL" \
	&& echo "$GOLANG_DOWNLOAD_SHA256 golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# get and compile boaring
ENV BORING_HOME $GOPATH/src/github.com/ajaysahani/boaring

RUN go get -u github.com/Masterminds/glide
RUN go get -u github.com/golang/lint/golint

RUN git clone https://github.com/uber/cadence.git $CADENCE_HOME
RUN cd $BORING_HOME && git checkout $git_branch

EXPOSE 7933 7934 7935

COPY ./start.sh $BORING_HOME/start.sh
COPY ./config_template.yaml $BORING_HOME/config/docker_template.yaml
RUN chmod a+x $BORING_HOME/start.sh

WORKDIR $BORING_HOME
CMD ./start.sh $BORING_HOME