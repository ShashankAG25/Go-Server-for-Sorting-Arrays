FROM debian:stretch-slim
ADD sortingServer /bin/sortingServer
CMD [ "/bin/sortingServer" ]
