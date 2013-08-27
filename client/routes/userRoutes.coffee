App.UsersNewRoute = Ember.Route.extend
  model: ->
    App.User

App.UserRoute = Ember.Route.extend
  model: (params) ->
    App.User.find(params.user_id)