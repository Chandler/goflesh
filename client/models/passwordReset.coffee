define ["ember-data"], (DS) ->
  PasswordReset = DS.Model.extend
    code: DS.attr 'string'
    expires: DS.attr 'string'

  becameError: ->
  	console.log 'HABAAHBA'
  	# @transitionTo 'discovery'
  becameInvalid: (errors) ->
  	console.log 'other errors'

  PasswordReset.toString = -> 
    "PasswordReset"

  PasswordReset