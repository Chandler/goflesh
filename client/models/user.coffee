App.User = DS.Model.extend(
  first_name: DS.attr 'string'
  last_name: DS.attr 'string'
  screen_name: DS.attr 'string'
  email: DS.attr 'string'
  phone: DS.attr 'string'
  avatar: DS.attr 'avatar'
  password: DS.attr 'string'
  players: DS.hasMany 'App.Player'
  organization: DS.belongsTo 'App.Organization'
)

App.User.toString = -> 
  "User"
