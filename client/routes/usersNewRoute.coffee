define ["ember", "User"], (Em, User) ->
 	UsersRoute = Ember.Route.extend
    model: ->
      User
