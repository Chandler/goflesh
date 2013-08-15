define ["ember-grid"], (GRID) ->
  GameHomeController = GRID.TableController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'players'
    
    #table controller stuff
    toolbar: [
        GRID.Filter
        GRID.ColumnSelector,
    ],

    columns: [
        GRID.column('name', { formatter: '{{avatar small}}', header: '' }),
        GRID.column('get_user', { header: '' }),
        GRID.column('id', { header: '' }),
        GRID.column('id', { header: '' }),
    ]