define ["ember-data"], (DS) ->
  PasswordReset = DS.Model.extend
    code: DS.attr 'string'
    expires: DS.attr 'string'

  PasswordReset.toString = -> 
    "PasswordReset"

  PasswordReset