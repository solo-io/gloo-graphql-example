FROM node
COPY db.json index.js package.json  ./ 
RUN npm install
EXPOSE 8080
CMD [ "node", "index.js" ]