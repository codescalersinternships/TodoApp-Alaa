FROM node:14-alpine
WORKDIR /usr/src/app

COPY . .
RUN npm install

EXPOSE 5656
ENV HOST=0.0.0.0
CMD [ "npm", "start"]