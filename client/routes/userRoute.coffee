define ["ember", "User"], (Em, User) ->
  UserRoute = Em.Route.extend
    model: (params) ->
      User.find(params.user_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('user').set 'selectedUser', model
