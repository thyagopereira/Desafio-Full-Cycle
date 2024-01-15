FROM node:20.11.0

WORKDIR /usr/src/app

# COPY files to generate node_modules
COPY package.json package.json
COPY package-lock.json package-lock.json

# RUN npm install
RUN npm install

COPY . .
EXPOSE 3000


CMD ["node","index.js"]