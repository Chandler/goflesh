define ["ember-data"], (DS) ->
  Member = DS.Model.extend
    organization: DS.belongsTo 'Em.App.Organization'
  
  Member.toString = -> 
    "Member"

  Member