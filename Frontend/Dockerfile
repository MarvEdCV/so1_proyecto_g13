FROM node:alpine
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm install -g npm@8.19.3
RUN npm install @craco/craco
CMD ["npm", "run", "start"]