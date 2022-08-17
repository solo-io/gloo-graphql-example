const { RESTDataSource } = require('apollo-datasource-rest');

class UserAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = 'http://127.0.0.1:8080';
  }

  async getUserDetails(username) {
    console.log("Searching user with username=", username)
    var  obj = await this.get("/users", {
      username: username
    })
   
    console.log(JSON.stringify(obj[0]))
    obj = (obj.length > 0) ? obj[0] : null ; 
    return obj;
  }
}

module.exports = UserAPI;
