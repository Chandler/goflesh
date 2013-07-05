define ["ember", "UserModel"], (Em, UserModel) ->
  UsersShowRoute = Ember.Route.extend
    model: (params) ->
    	console.log(params)
      UserModel.find(params.user_id)
