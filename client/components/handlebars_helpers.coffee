# Might be useful later http://techblog.fundinggates.com/blog/2012/08/ember-handlebars-helpers-bound-and-unbound/
define ["handlebars", "utilities"], (Handlebars, Utilities) ->
  Handlebars.registerHelper 'avatar', (size, options) ->
    key = this.get('name')
    new Handlebars.SafeString(Utilities.avatarTag(key, size, options))

