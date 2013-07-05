define ["ember", "UserModel"], (Em, UserModel) ->
 	UsersRoute = Ember.Route.extend
    model: ->
      UserModel
