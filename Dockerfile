FROM golang:1.12-stretch
MAINTAINER IPWeb <contact@ipweb.io>

ENV SRC_DIR /go-ipfs

# Download packages first so they can be cached.
COPY go.mod go.sum $SRC_DIR/
RUN cd $SRC_DIR \
  && git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/go-ipws-config".insteadOf "https://github.com/ipweb-group/go-ipws-config" \
  && git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/interface-go-ipws-core".insteadOf "https://github.com/ipweb-group/interface-go-ipws-core" \
  && git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/ipw".insteadOf "https://github.com/ipweb-group/ipw" \
  && go mod download

COPY . $SRC_DIR

# Build the thing.
# Also: fix getting HEAD commit hash via git rev-parse.
RUN cd $SRC_DIR \
  && mkdir .git/objects \
  && make build

# Get su-exec, a very minimal tool for dropping privileges,
# and tini, a very minimal init daemon for containers
ENV SUEXEC_VERSION v0.2
ENV TINI_VERSION v0.16.1
RUN set -x \
  && cd /tmp \
  && git clone https://github.com/ncopa/su-exec.git \
  && cd su-exec \
  && git checkout -q $SUEXEC_VERSION \
  && make \
  && cd /tmp \
  && wget -q -O tini https://github.com/krallin/tini/releases/download/$TINI_VERSION/tini \
  && chmod +x tini

# Get the TLS CA certificates, they're not provided by busybox.
RUN apt-get update && apt-get install -y ca-certificates

# Install FUSE
RUN apt-get update && apt-get install -y fuse

# Now comes the actual target image, which aims to be as small as possible.
FROM busybox:1-glibc
MAINTAINER IPWeb <contact@ipweb.io>

# Get the ipws binary, entrypoint script, and TLS CAs from the build container.
ENV SRC_DIR /go-ipfs
COPY --from=0 $SRC_DIR/cmd/ipws/ipws /usr/local/bin/ipws
COPY --from=0 $SRC_DIR/bin/container_daemon /usr/local/bin/start_ipws
COPY --from=0 /tmp/su-exec/su-exec /sbin/su-exec
COPY --from=0 /tmp/tini /sbin/tini
COPY --from=0 /bin/fusermount /usr/local/bin/fusermount
COPY --from=0 /etc/ssl/certs /etc/ssl/certs

# Add suid bit on fusermount so it will run properly
RUN chmod 4755 /usr/local/bin/fusermount

# This shared lib (part of glibc) doesn't seem to be included with busybox.
COPY --from=0 /lib/x86_64-linux-gnu/libdl-2.24.so /lib/libdl.so.2

# Swarm TCP; should be exposed to the public
EXPOSE 4001
# Daemon API; must not be exposed publicly but to client services under you control
EXPOSE 5001
# Web Gateway; can be exposed publicly with a proxy, e.g. as https://ipfs.example.org
EXPOSE 8080
# Swarm Websockets; must be exposed publicly when the node is listening using the websocket transport (/ipX/.../tcp/8081/ws).
EXPOSE 8081

# Create the fs-repo directory and switch to a non-privileged user.
ENV IPWS_PATH /data/ipws
RUN mkdir -p $IPWS_PATH \
  && adduser -D -h $IPWS_PATH -u 1000 -G users ipws \
  && chown ipws:users $IPWS_PATH

# Create mount points for `ipws mount` command
RUN mkdir /ipws /ipns \
  && chown ipws:users /ipws /ipns

# Expose the fs-repo as a volume.
# start_ipws initializes an fs-repo if none is mounted.
# Important this happens after the USER directive so permission are correct.
VOLUME $IPWS_PATH

# The default logging level
ENV IPWS_LOGGING ""

# This just makes sure that:
# 1. There's an fs-repo, and initializes one if there isn't.
# 2. The API and Gateway are accessible from outside the container.
ENTRYPOINT ["/sbin/tini", "--", "/usr/local/bin/start_ipws"]

# Execute the daemon subcommand by default
CMD ["daemon", "--migrate=true"]
