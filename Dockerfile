FROM scratch
COPY password /
ENTRYPOINT ["/password"]
