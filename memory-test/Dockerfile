FROM registry.cn-beijing.aliyuncs.com/shannonai/debian
LABEL maintainer="taoshibo@shannonai.com"

WORKDIR /root
COPY entrypoint.sh entrypoint.sh
COPY ./simplehttpserver simplehttpserver
RUN chmod u+x simplehttpserver && chmod u+x entrypoint.sh

EXPOSE 80

ENTRYPOINT ["./entrypoint.sh"]
CMD ["simplehttpserver"] # set default arg for entrypoint
