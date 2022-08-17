const { ApolloServer } = require('apollo-server');
const typeDefs = require('./schema');
const resolvers = require('./resolvers');

const UserAPI = require('./datasources/user-api');

const BASIC_LOGGING = {
  requestDidStart(requestContext) {
      console.log("request started");
      console.log(requestContext.request.query);
      console.log(requestContext.request.variables);
      return {
          didEncounterErrors(requestContext) {
              console.log("an error happened in response to query " + requestContext.request.query);
              console.log(requestContext.errors);
          }
      };
  },

  willSendResponse(requestContext) {
      console.log("response sent", requestContext.response);
  }
};

const server = new ApolloServer({
  typeDefs,
  plugins: [BASIC_LOGGING],
  resolvers,
  dataSources: () => {
    return {
      userAPI: new UserAPI(),
    };
  },
});

server.listen().then(() => {
  console.log(`
    🚀  Server is running!
    🔉  Listening on port 4000
    📭  Query at https://studio.apollographql.com/dev
  `);
});
