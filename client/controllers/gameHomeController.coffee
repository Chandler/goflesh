define ["ember-grid"], (GRID) ->
  GameHomeController = GRID.TableController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    organizationBinding: 'organization'
    contentBinding: 'players'
    
    #table controller stuff
    toolbar: [
        GRID.Filter
        GRID.ColumnSelector,
    ],

    columns: [
        GRID.column('name', { formatter: '{{avatar small}}', header: '' }),
        GRID.column('id', { header: '' }),

    ]