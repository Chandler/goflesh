define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'

    becameError: (args)->
      # handle error case here
      alert 'there was an error!'
  Organization.toString = -> 
    "Organization"


  Organization