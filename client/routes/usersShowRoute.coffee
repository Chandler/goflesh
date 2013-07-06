define ["ember", "User"], (Em, User) ->
  UsersShowRoute = Ember.Route.extend
    model: (params) ->
      User.find(params.user_id)
