App.PasswordResetRoute = Ember.Route.extend
  model: (params) ->
    # console.log 'router'
    # result = App.PasswordReset.find(params)
    # console.log 'found object'
    # result
    console.log params['code']
    App.Auth.signIn
      data:
        api_key: params['code']