FROM 	debian:bullseye-slim

ENV 	PATH=/usr/local/go/bin:/go/bin:${PATH}
ENV 	GOBIN=/go/bin
ENV 	GOPATH=/go
COPY 	./.docker/init.sh /etc/init.sh
RUN 	set -eux; \
		apt-get update; \
		apt-get install -y --no-install-recommends \
			wget \
			software-properties-common \
			apt-transport-https \
		; \
		wget https://golang.org/dl/go1.17.linux-amd64.tar.gz -O tmp/go1.17.linux-amd64.tar.gz; \
		tar -zxvf /tmp/go1.17.linux-amd64.tar.gz -C /usr/local/; \
		mkdir 0775 /var/go; \
		chmod 0775 /etc/init.sh; \
		rm -rf /tmp/*; \
		rm -rf /var/lib/apt/lists/*

WORKDIR	/go

EXPOSE 	80

VOLUME 	["/source"]
CMD 	["/etc/init.sh"]
