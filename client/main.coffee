#requireJS bootstraper.
require.config baseUrl: "public/js"
require ["jquery", "game", "app", "templates"], ($, game, app, templates) ->
  console.log "who"
  debugger

