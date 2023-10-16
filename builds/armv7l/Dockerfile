FROM ubuntu:20.04

RUN apt-get update \
    && apt-get install -y gcc-multilib-arm-linux-gnueabihf binutils-arm-linux-gnueabihf g++-multilib-arm-linux-gnueabihf
RUN DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get install -y git cmake

RUN mkdir /workspace && cd /workspace \
    && git clone https://github.com/libjpeg-turbo/libjpeg-turbo.git \
    && mkdir libjpeg-turbo/build
COPY linux-arm-toolchain.cmake /workspace/libjpeg-turbo/build

RUN cd /workspace/libjpeg-turbo/build \
    && cmake -G"Unix Makefiles" -DCMAKE_TOOLCHAIN_FILE=linux-arm-toolchain.cmake -DREQUIRE_SIMD=1 .. ${1+"$@"} \
    && make -j 2 \
    && make install


WORKDIR /opt