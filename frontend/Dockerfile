FROM node:lts-alpine

RUN npm install -g serve

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN echo "VITE_BASE_API=http://13.48.114.197:8080/" > .env

RUN npm run build

EXPOSE 80
CMD [ "serve", "-s", "-p", "80", "dist" ]