FROM busybox
ADD ./key_linux-amd64 /app
CMD ["/app"]
