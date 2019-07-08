FROM zaqwsx_ipws-test-img

RUN ipws init -b=1024
ADD . /tmp/id
RUN mv -f /tmp/id/config /root/.ipws/config
RUN ipws id

EXPOSE 4031 4032/udp

ENV IPWS_PROF true
ENV IPWS_LOGGING_FMT nocolor

ENTRYPOINT ["/bin/bash"]
CMD ["/tmp/id/run.sh"]
