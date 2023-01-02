FROM klakegg/hugo:0.101.0-onbuild AS hugo

FROM caddy:2
COPY --from=hugo /target /
RUN caddy file-server
