FROM alpine

EXPOSE 8000
WORKDIR /app
# copy binary into image
COPY linman /app/
CMD ["./linman"]
