FROM node:22-alpine
WORKDIR /usr/app
COPY frontend /usr/app/
RUN npm install && npm run build
EXPOSE 3000

ENTRYPOINT ["npm", "run", "start"]