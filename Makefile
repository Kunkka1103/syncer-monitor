# 定义目标名称
TARGET = syncer-monitor

# 定义 Go 源文件
SRC = syncer-monitor.go

# 默认目标
all: build

# 编译命令
build:
	go build -o $(TARGET) $(SRC)

# 清理命令
clean:
	rm -f $(TARGET)

# 伪目标，不会与实际文件名冲突
.PHONY: all build clean

