FROM registry.cn-beijing.aliyuncs.com/shannonai/debian
LABEL maintainer="hechenrui123@gmail.com"

WORKDIR /root
COPY entrypoint.sh entrypoint.sh
COPY parent.sh parent.sh
COPY child.sh child.sh
COPY ./memory-test memory-test
RUN chmod u+x memory-test && chmod u+x entrypoint.sh && chmod u+x parent.sh

EXPOSE 80

ENTRYPOINT ["./entrypoint.sh"]
CMD ["parent"] # set default arg for entrypoint
