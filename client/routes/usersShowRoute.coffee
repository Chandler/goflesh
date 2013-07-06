define ["ember", "UserModel"], (Em, UserModel) ->
  UsersShowRoute = Ember.Route.extend
    model: (params) ->
      UserModel.find(params.user_id)
