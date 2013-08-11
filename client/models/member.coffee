define ["ember-data"], (DS) ->
  Member = DS.Model.extend
    organziation: DS.belongsTo 'Em.App.Organization'
  
  Member.toString = -> 
    "Member"

  Member