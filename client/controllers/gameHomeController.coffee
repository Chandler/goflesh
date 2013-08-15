define ["ember-grid"], (GRID) ->
  GameHomeController =  GRID.TableController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'gridModel'
    toolbar: [
        GRID.Filter
    ],

    columns: [
        GRID.column('name', { formatter: '{{avatar small}}', header: '' }),
        GRID.column('name', { header: '' }),
    ]