define ["templates", "ember-data"], (Templates, DS) ->

  User = DS.Model.extend
    first_name: DS.attr 'string'
    last_name: DS.attr 'string'
    screen_name: DS.attr 'string'
    email: DS.attr 'string'
    password: DS.attr 'string'

  User.toString = -> 
    "User"

  User