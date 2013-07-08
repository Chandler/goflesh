define ["ember", "User"], (Em, User) ->
  UsersShowRoute = Em.Route.extend
    model: (params) ->
      User.find(params.user_id)
