FROM node:17
WORKDIR /clientgrcp
COPY . .
RUN cd Client
RUN npm install
EXPOSE 3000
CMD ["node","index.js"] 