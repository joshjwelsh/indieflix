# Use an official Ubuntu runtime as a parent image
FROM ubuntu:20.04

# Set the working directory in the container to /app
WORKDIR /app

# Add the current directory contents into the container at /app
ADD . /app

# Install dependencies
RUN apt-get update && apt-get install -y \
	python3 \
	python3-dev \
	python3-pip \
	libxml2-dev \
	libxslt1-dev \
	zlib1g-dev \
	libffi-dev \
	libssl-dev

# Install Scrapy
RUN pip3 install Scrapy

# Open bash when the container launches
CMD ["/bin/bash"]
