FROM ubuntu:latest

## install all dependencies
RUN apt-get update
RUN apt-get install -y curl gpg

## create non-root user with home directory
RUN useradd -ms /bin/bash -u 10001 litecoinuser

## set user and work directory
USER litecoinuser
WORKDIR /home/litecoinuser

## Download linux binary (https://litecoin.org/)
RUN curl -O https://download.litecoin.org/litecoin-0.18.1/linux/litecoin-0.18.1-x86_64-linux-gnu.tar.gz
RUN ls -l
## import GPG public key (https://download.litecoin.org/README-HOWTO-GPG-VERIFY-TEAM-MEMBERS-KEY.txt)
RUN gpg --keyserver pgp.mit.edu --recv-key FE3348877809386C
## Download GPG signature
RUN curl -O https://download.litecoin.org/litecoin-0.18.1/linux/litecoin-0.18.1-linux-signatures.asc

## Verify downloaded file
RUN gpg --verify litecoin-0.18.1-linux-signatures.asc

## extract tar zip file
RUN tar -xzvf litecoin-0.18.1-x86_64-linux-gnu.tar.gz

## run daemon
ENTRYPOINT ["litecoin-0.18.1/bin/litecoind"]
