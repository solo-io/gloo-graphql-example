const resolvers = {
  Query: {
    getUserDetails: (_, args, { dataSources }) => {
      return dataSources.userAPI.getUserDetails(args.username);
    },
  }
};

module.exports = resolvers;
