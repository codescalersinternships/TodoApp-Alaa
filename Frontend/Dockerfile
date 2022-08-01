FROM node:12.22.9

WORKDIR /usr/my-svelte-project/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3001

ENV HOST=0.0.0.0

CMD ["npm","run","dev"]