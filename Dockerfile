FROM golang:latest

# Install air
RUN go install github.com/cosmtrek/air@latest

# Intall go migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# https://github.com/otiai10/gosseract/blob/main/Dockerfile
RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata/

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-jpn
# See https://github.com/tesseract-ocr/tessdata for the list of available languages.
# If you want to download these traineddata via `wget`, don't forget to locate
# downloaded traineddata under ${TESSDATA_PREFIX}/tessdata.

RUN mkdir main

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

#May need to run chmod +x ./start.sh locally if mounting host to container
RUN chmod +x ./start.sh 

RUN go build -o ../main/main ./cmd/server

WORKDIR ../app
EXPOSE 8080

ENTRYPOINT [ "./start.sh" ] 