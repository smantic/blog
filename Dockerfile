FROM klakegg/hugo as hugo 

COPY . . 
RUN hugo 

FROM caddy:2

COPY --from=hugo public /
RUN caddy file-server
