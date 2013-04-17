define(["ember", "templates", "ember-data"], function(Em, Templates, DS) {
  var Game, Store;

  Store = DS.Store.extend({
    revision: 11,
    adapter: 'DS.fixtureAdapter'
  });
  Game = DS.Model.extend();
  return Game.FIXTURES = [
    {
      id: 1,
      id: 2,
      id: 3
    }
  ];
});
