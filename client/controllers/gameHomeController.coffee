define ["ember-grid"], (GRID) ->
  GameHomeController = GRID.TableController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'gridModel'
    toolbar: [
        GRID.Filter
        GRID.ColumnSelector,
    ],

    columns: [
        GRID.column('name', { formatter: '{{avatar small}}', header: '' }),
        GRID.column('name', { header: 'Name' }),
        GRID.column('id', { header: 'Status' }),
        GRID.column('id', { header: 'Team' }),
        GRID.column('id', { header: 'Time to live' }),
    ]