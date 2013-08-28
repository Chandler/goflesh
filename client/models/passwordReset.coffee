App.PasswordReset = DS.Model.extend
  user_id: DS.attr 'number'
  api_key: DS.attr 'string'
  code: DS.attr 'string'
  # expires: DS.attr 'string'

App.PasswordReset.toString = -> 
  "PasswordReset"