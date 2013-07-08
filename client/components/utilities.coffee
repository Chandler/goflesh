define [], ->
  Utilities =
    avatarTag: (hash, size, options = {}) ->
      sizes =
        small: 50
        large: 100
        profile: 150
      px = sizes[size]
      cls = if options.cls? then options.cls else ""
      
      random = Math.random().toString(16).slice(2)
      hash = random + random + random + random

      "<img class=\"" + options.cls +  "\" src=\"http://www.gravatar.com/avatar/" + hash + "?s=" + px + "&d=identicon\"/>"

