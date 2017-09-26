
ssize_t write(int fd, const void *buf, size_t count);

ssize_t read(int fd, void *buf, size_t count);

off_t lseek(int fd, off_t offset, int whence);

int close(int fd);
