FROM alpine

RUN mkdir -p /opt/promise
RUN mkdir -p /opt/promise/log
RUN touch /opt/promise/log/promise.log
WORKDIR /opt/promise
COPY fs ./
RUN chmod a+x ./director ./wait-for.sh

EXPOSE 80

# ENTRYPOINT ["./director"]