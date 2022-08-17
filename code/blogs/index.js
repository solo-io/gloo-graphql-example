const jsonServer = require('json-server')
const server = jsonServer.create()
const router = jsonServer.router('db.json')
const middlewares = jsonServer.defaults()
const fs = require('fs');

let rawdata = fs.readFileSync('swagger.json');
let swagger = JSON.parse(rawdata);

server.use(middlewares)

server.get('/swagger.json', (req, res) => {
  res.jsonp(swagger)
})

server.use(router)
server.listen(8080, () => {
  console.log('JSON Server is running')
})