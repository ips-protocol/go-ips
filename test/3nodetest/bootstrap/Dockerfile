FROM zaqwsx_ipws-test-img

RUN ipws init -b=1024
ADD . /tmp/id
RUN mv -f /tmp/id/config /root/.ipws/config
RUN ipws id

ENV IPWS_PROF true
ENV IPWS_LOGGING_FMT nocolor

EXPOSE 4011 4012/udp
