set(CMAKE_SYSTEM_NAME Linux)
set(CMAKE_SYSTEM_PROCESSOR arm)
set(CMAKE_C_FLAGS "-mfloat-abi=hard -march=armv7-a -mfpu=neon -mthumb")
set(CMAKE_C_COMPILER /usr/bin/arm-linux-gnueabihf-gcc)