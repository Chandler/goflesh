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
  
  #TODO change these
  eventFeedSelected: false
  playerListSelected: true
  showPlayerList: ->
    @set 'eventFeedSelected', false
    @set 'playerListSelected', true
  showEventFeed: ->
    @set 'eventFeedSelected', true
    @set 'playerListSelected', false

  registerTag: ->
    code = "VPCQG"
    game_id = @get('game.id')
    current_player = App.Auth.get('user.players').filter (p) =>
      p.get('game.id') == game_id
    current_player = App.Auth.get('user.players').filter (p) =>
      p.get('game.id') == game_id
    player_id = current_player[0].get('id')
    $.post("/api/tag/" + code + "?player_id=" + player_id).done(e) ->
      console.log(e)

App.GamesController = Ember.ObjectController.extend
  selectedGame: null

App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'