FROM node:20 as build

WORKDIR /workspace

COPY . .
RUN npm install

EXPOSE 3000

CMD ["npm", "start"]
