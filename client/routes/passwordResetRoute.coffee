define ["ember", "PasswordReset"], (Em, PasswordReset) ->
  PasswordResetRoute = Em.Route.extend
    model: (params) ->
      PasswordReset.find(params)
