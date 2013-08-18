App.GameHomeController =  Ember.Controller.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'game.players'
