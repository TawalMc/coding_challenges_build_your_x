CC = gcc
CFLAGS = -Wall -g
TARGET = ccwc
SRCS = main.c ccwc.c
OBJS = $(SRCS:.c=.o)

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) $(OBJS) -o $(TARGET)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

main.o: main.c ccwc.h

file.o: ccwc.c ccwc.h

run_error: $(TARGET)
	./$(TARGET) ../test1.txt

run_empty: $(TARGET)
	./$(TARGET) -m -c ../test1.txt

run_full: $(TARGET)
	./$(TARGET) ../test.txt

clean:
	rm -f $(TARGET) $(OBJS)

.PHONY: all run_empty run_error run_full clean
