define ["ember", "User"], (Em, User) ->
 	UsersRoute = Em.Route.extend
    model: ->
      User
