App.UserRoute = Ember.Route.extend
  model: (params) ->
    App.User.find(params.user_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('user').set 'selectedUser', model
