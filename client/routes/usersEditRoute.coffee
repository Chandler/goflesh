define ["ember", "User"], (Em, User) ->
  UsersEditRoute = Ember.Route.extend
    model: (params) ->
      User.find(params.user_id)
