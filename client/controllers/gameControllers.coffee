App.GamesNewController = NewController.extend
  editableRecordFields: ['name', 'slug']
  name: '',
  slug: '',

App.GameHomeController =  Ember.Controller.extend
  code: ''
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  playersBinding: 'game.players'
  eventsBinding: 'events'
  
  selectedList: 'eventList'
  
  selectList: (list) ->
    @set 'selectedList', list

  eventListSelected:  Ember.computed.equal('selectedList', 'eventList')
  playerListSelected: Ember.computed.equal('selectedList', 'playerList')


  showRegisterTag: (->
    currentUser.belongsToGame(@get('game.id')) && currentUser.isZombie() 
  ).property()

  registerTag: ->
    code = "VPCQG"
    currentPlayer = currentUser.playerForGame(@get('game.id'))
    $.post("/api/tag/" + code + "?player_id=" + currentPlayer.get('id')).done(e) ->
      console.log(e)


App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'