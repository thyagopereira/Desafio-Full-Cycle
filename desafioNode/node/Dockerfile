FROM node:20.11.0

WORKDIR /usr/src/app

#Dockerize app
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# COPY files to generate node_modules
COPY package.json package.json
COPY package-lock.json package-lock.json

# RUN npm install
RUN npm install

COPY . .
EXPOSE 3000
CMD ["node","index.js"]