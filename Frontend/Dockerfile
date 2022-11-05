FROM node:latest AS node
WORKDIR /app
COPY . .
RUN npm install
CMD ["npm", "run", "build"]

FROM nginx:latest
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=node /app/build/ /usr/share/nginx/html