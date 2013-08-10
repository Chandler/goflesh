define ["Utilities"], (Utilities) ->
  describe 'Utilities.avatarTag', ->
    it "returns a correct avatarTag", ->
      avatarTag = Utilities.avatarTag('asdfas','small', {})
      expect(avatarTag).toContain 'img'
