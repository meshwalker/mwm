FROM scratch
 COPY ./mwm /mwm
 ENTRYPOINT ["/mwm"]