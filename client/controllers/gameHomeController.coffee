define ["ember-grid"], (GRID) ->
  GameHomeController =  GRID.TableController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'game.players'
    toolbar: [
        GRID.Filter
    ],

    columns: [
        GRID.column('id', { formatter: '{{avatar small}}', header: '' }),
        GRID.column('user.first_name', { header: 'Name' }),
        GRID.column('id', { header: 'id' }),

    ]