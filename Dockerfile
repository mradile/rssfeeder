FROM golang:1.12-stretch as gobuilder

#run as unpriviledged user
RUN addgroup --gid 990 app && adduser --disabled-password --uid 990 --gid 990 --gecos '' app

RUN mkdir -p /build /data
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build


FROM scratch

EXPOSE 3000
VOLUME /data
ENV DB /data

COPY --chown=990:990 --from=gobuilder /data /data
COPY --from=gobuilder /etc/passwd /etc/passwd
COPY --chown=990:990 --from=gobuilder /build/release/rssfeederd /rssfeederd
COPY --chown=990:990 --from=gobuilder /build/release/rssfeeder  /rssfeeder

USER 990:990

ENTRYPOINT ["./rssfeederd"]
CMD ["run"]
