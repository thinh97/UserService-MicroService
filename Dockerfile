FROM iron/base

EXPOSE 6767
ADD userservice /
ENTRYPOINT ["./UserService"]