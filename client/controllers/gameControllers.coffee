App.GamesNewController = NewController.extend
  editableRecordFields: ['name', 'slug']
  name: '',
  slug: '',

App.GameHomeController =  Ember.Controller.extend
  code: ''
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'game.players'
  
  #this can be refactored into something like
  #selected: 'eventFeed'
  #showEventFeed = Ember.computed(selected, 'eventFeed')

  eventFeedSelected: false
  playerListSelected: true
  showPlayerList: ->
    @set 'eventFeedSelected', false
    @set 'playerListSelected', true
  showEventFeed: ->
    @set 'eventFeedSelected', true
    @set 'playerListSelected', false

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