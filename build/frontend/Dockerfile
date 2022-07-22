FROM node:latest
LABEL Maintainer="Souki Aoba <souki.aoba@ginco.co.jp>"

RUN mkdir -p /usr/src/iost-explorer
COPY ./frontend /usr/src/iost-explorer/frontend
WORKDIR /usr/src/iost-explorer/frontend
RUN npm install
RUN npm run build

EXPOSE 8080
CMD ["npm", "start"]
