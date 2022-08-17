const { gql } = require('apollo-server');

const typeDefs = gql`
type User {
  id: Int!
  username: String
  lastname: String
  firstname: String
}

type Query {
  getUserDetails(username: String!): User
}
`;

module.exports = typeDefs;
