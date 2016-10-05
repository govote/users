FROM scratch

COPY users /
COPY banner.txt /

EXPOSE 8081

CMD ["/users"]
