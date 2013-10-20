App.Game = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  description: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  running_start_time: DS.attr 'string'
  players: DS.hasMany 'App.Player'

  showRegisterTag: (->
    return true
    # user = App.Auth.get('user')
    # player = @get('players').findProperty('user', user)
    # console.log "print"
    # console.log player.get('status')
    # console.log player
    # player
  ).property('players.@each.status')
  
  currentPlayer: (->
    user = App.Auth.get('user')
    @get('players').findProperty('user', user)
  ).property('players.@each.user')

App.Game.toString = -> 
  "Game"