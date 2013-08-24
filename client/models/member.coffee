App.Member = DS.Model.extend
  organization: DS.belongsTo 'App.Organization'

App.Member.toString = -> 
  "Member"