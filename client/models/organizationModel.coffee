define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    thumb_url: DS.attr 'string'
    member_count: DS.attr 'number'
    location: DS.attr 'string'
    show: DS.attr 'bool'


  Organization.FIXTURES = [
    {
      id: 1,
      name: 'University of Idaho HZ',
      thumb_url: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-ash3/c0.50.180.180/22176_305111860995_8060208_a.jpg',
      member_count: 199,
      location: 'Moscow, Idaho',
      show: true 
    },
    {
      id: 2,
      name: 'HZ Washington State University',
      thumb_url: 'https://sphotos-a.xx.fbcdn.net/hphotos-ash4/c356.0.604.604/s320x320/376136_10151643919862704_963634461_n.jpg',
      member_count: 387,
      location: 'Pullman, Washington'
    },
    {
      id: 3,
      name: 'HZ Stanford',
      thumb_url: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-ash3/41800_6192688417_1783896943_q.jpg',
      member_count: 612,
      location: 'Palo Alto, CA'
    }
  ]

  Organization