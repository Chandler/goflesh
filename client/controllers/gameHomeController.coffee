define ["ember"], (Em) ->
  GameHomeController =  Em.Controller.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    
    #table controller stuff
    # toolbar: [
    #     GRID.Filter
    #     GRID.ColumnSelector,
    # ],

    # columns: [
    #     GRID.column('name', { formatter: '{{avatar small}}', header: '' }),
    #     GRID.column('id', { header: '' }),

    # ]