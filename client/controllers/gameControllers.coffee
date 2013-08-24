App.GamesNewController = NewController.extend
  editableRecordFields: ['name', 'slug']
  name: '',
  slug: '',

App.GameHomeController =  Ember.Controller.extend
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
      code = "JNHDB"
      $.post("/api/tag/" + code).done (e) ->
        user = App.User.find({user_id: 2})
        console.log user.get('status')
    

App.GamesController = Ember.ObjectController.extend
  selectedGame: null

App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'