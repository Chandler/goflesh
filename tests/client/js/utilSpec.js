define(["Utilities"], function(Utilities) {
  return describe('Utilities.avatarTag', function() {
    return it("returns a correct avatarTag", function() {
      var avatarTag;

      avatarTag = Utilities.avatarTag('asdfas', 'small', {});
      return expect(avatarTag).toContain('img');
    });
  });
});
