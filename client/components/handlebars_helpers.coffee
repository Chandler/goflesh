# Might be useful later http://techblog.fundinggates.com/blog/2012/08/ember-handlebars-helpers-bound-and-unbound/
Handlebars.registerHelper 'avatar', (size, options) ->
  key = "doesn't matter"
  new Handlebars.SafeString(Utilities.avatarTag(key, size, options))


Ember.Handlebars.registerBoundHelper 'avatar2', (avatar, options) ->
  key = avatar
  size = options.hash.size
  klass = options.hash.class
  new Handlebars.SafeString(Utilities.avatar2Tag(key, size, klass))

