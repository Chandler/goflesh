App.GamesNewController = NewController.extend
  editableRecordFields: ['name', 'slug']
  name: '',
  slug: '',

App.GameHomeController =  Ember.Controller.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'game.players'
    
    #this is a gross proof of concept
    #TODO figure out the right way to make a stateful view component
    #that doesn't rely on controller values like this.
    eventFeedSelected: false
    playerListSelected: true
    showPlayerList: ->
      @set 'eventFeedSelected', false
      @set 'playerListSelected', true
    showEventFeed: ->
      @set 'eventFeedSelected', true
      @set 'playerListSelected', false

    registerTag: ->
      code = "CR66W"
      # $.post("/api/tag/" + code).done (e) ->
      user = App.User.find({user_id: 2})
      
      console.log user.get('status')
    

App.GamesController = Ember.ObjectController.extend
  selectedGame: null

App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  contentBinding: 'controllers.game.content'