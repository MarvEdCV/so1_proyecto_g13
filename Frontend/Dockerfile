FROM node:latest AS node
WORKDIR /app
COPY . .
RUN npm install -g npm@8.19.3
CMD ["npm", "run", "start"]