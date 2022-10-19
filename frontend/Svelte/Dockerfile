FROM node:14-alpine
WORKDIR /usr/src/app

COPY . .
RUN npm install
RUN npm run-script build

ENV HOST=0.0.0.0
CMD [ "npm", "start"]