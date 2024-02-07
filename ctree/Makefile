CC = gcc
CFLAGS = -Wall -g
LDFLAGS =
INC_DIRS = -Isrc
SRC_DIR = src
TESTS_DIR = tests
OBJ_DIR = obj
BIN_DIR = bin

# Source files
SRCS = $(wildcard $(SRC_DIR)/*.c)
# Exclude main.c from the library build
LIB_SRCS = $(filter-out $(SRC_DIR)/main.c, $(SRCS))
# Object files for the library
LIB_OBJS = $(LIB_SRCS:$(SRC_DIR)/%.c=$(OBJ_DIR)/%.o)
# Test files
TEST_SRCS = $(wildcard $(TESTS_DIR)/*.c)
# Application binary
APP = $(BIN_DIR)/app
# Test binaries
TEST_BINS = $(TEST_SRCS:$(TESTS_DIR)/%.c=$(BIN_DIR)/%)

all: $(APP) tests

$(APP): $(OBJ_DIR)/main.o $(LIB_OBJS) | $(BIN_DIR)
	@mkdir -p $(BIN_DIR) # Ensure bin directory exists before linking
	$(CC) $^ -o $@ $(LDFLAGS)

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) $(INC_DIRS) -c $< -o $@

$(BIN_DIR)/%: $(TESTS_DIR)/%.c $(LIB_OBJS) | $(BIN_DIR)
	@mkdir -p $(BIN_DIR) # Ensure bin directory exists before creating test binaries
	$(CC) $(CFLAGS) $(INC_DIRS) $^ -o $@ $(LDFLAGS)

$(OBJ_DIR) $(BIN_DIR):
	mkdir -p $@

tests: $(TEST_BINS)

test: tests
	@for test_bin in $(TEST_BINS); do \
		echo Running $$test_bin; \
		./$$test_bin; \
	done


clean:
	rm -rf $(OBJ_DIR) $(BIN_DIR)

.PHONY: all tests clean
