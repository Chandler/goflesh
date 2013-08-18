# Might be useful later http://techblog.fundinggates.com/blog/2012/08/ember-handlebars-helpers-bound-and-unbound/
Handlebars.registerHelper 'avatar', (size, options) ->
  key = "doesn't matter yet"
  new Handlebars.SafeString(Utilities.avatarTag(key, size, options))

