define ["ember", "UserModel"], (Em, UserModel) ->
  UsersEditRoute = Ember.Route.extend
    model: (params) ->
      UserModel.find(params.user_id)
