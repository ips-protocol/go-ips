FROM zaqwsx_ipws-test-img

RUN ipws init -b=1024
ADD . /tmp/test
RUN mv -f /tmp/test/config /root/.ipws/config
RUN ipws id
RUN chmod +x /tmp/test/run.sh

EXPOSE 4021 4022/udp

ENV IPWS_PROF true
ENV IPWS_LOGGING_FMT nocolor

ENTRYPOINT ["/bin/bash"]
CMD ["/tmp/test/run.sh"]
