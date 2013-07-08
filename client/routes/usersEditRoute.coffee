define ["ember", "User"], (Em, User) ->
  UsersEditRoute = Em.Route.extend
    model: (params) ->
      User.find(params.user_id)
