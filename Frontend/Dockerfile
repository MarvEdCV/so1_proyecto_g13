FROM node:19-alpine3.15
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm install
RUN npm install @craco/craco
CMD ["npm", "run", "start"]