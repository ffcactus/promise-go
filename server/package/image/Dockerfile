FROM alpine

RUN mkdir -p /opt/promise
RUN touch /tmp/promise.log
WORKDIR /opt/promise
COPY fs ./

EXPOSE 80

# ENTRYPOINT ["./server"]
ENTRYPOINT ["sleep", "9999m"]