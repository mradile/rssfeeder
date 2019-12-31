FROM alpine:3.9

EXPOSE 3000
VOLUME /data
ENV DB /data

RUN mkdir -p /build /data

RUN addgroup -g 990 rssfeeder && \
    adduser -D -u 990 -G rssfeeder rssfeeder

COPY rssfeederd /rssfeederd
RUN chown -R 990:990 /data

USER 990:990

ENTRYPOINT ["./rssfeederd"]
CMD ["run"]
