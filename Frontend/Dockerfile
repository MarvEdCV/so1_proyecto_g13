FROM node:19-alpine3.15
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm install
CMD ["npm", "run", "start"]