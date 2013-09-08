App.GamesNewController = NewController.extend
  editableRecordFields: ['name', 'slug']
  name: '',
  slug: '',

App.GameHomeController =  BaseController.extend
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
  
  registerTag: ->
    if @code != ''
      @clearErrors()
      currentPlayer = @get('game.currentPlayer')
      $.post("/api/tag/" + @code + "?player_id=" + currentPlayer.get('id'))
      .done (xhr, status, error) =>
        @set 'errors', "success!" 
      .fail (xhr, status, error) =>
        @set 'errors', JSON.stringify(xhr.responseText) 
    else
        @set 'errors', "human code empty"


App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'



App.GamePlayersController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'
  playersBinding: 'game.players'
