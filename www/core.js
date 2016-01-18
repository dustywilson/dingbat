var s;
var wantBlocker = 2;

$(function(){
  updateBlocker();
  s = gotalk.connection().on('open', onConnect).on('close', onDisconnect);
});

function addBlocker() {
  wantBlocker++;
  updateBlocker();
}

function removeBlocker() {
  wantBlocker--;
  updateBlocker();
}

function updateBlocker() {
  if (wantBlocker < 0) wantBlocker = 0;
  if (wantBlocker > 0) return $("body").addClass("wantblocker");
  $("body").removeClass("wantblocker");
}

function logoClick() {
  alert("Logo clicked...");
}

function onSignIn(googleUser) {
  var id_token = googleUser.getAuthResponse().id_token;

  var profile = googleUser.getBasicProfile();
  console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName());
  console.log('Image URL: ' + profile.getImageUrl());
  console.log('Email: ' + profile.getEmail());

  s.request("client.init-google", id_token, function (err, result) {
    if (err) {
      console.error('client.init-google failed:', err);
      return signOut(); // or something else?
    }
    $("body").addClass("loggedin");
    removeBlocker();
    console.log(result);
    onLoggedInConnect();
  })
}

gotalk.handleNotification('connection.info', function(c) {
  console.log('connection.info.ID:', c.ID);
});

function signOut() {
  var auth2 = gapi.auth2.getAuthInstance();
  auth2.signOut().then(function () {
    console.log('User signed out.');
    $("body").removeClass("loggedin");
    addBlocker();
    s.request("client.signout-google", {}, function(err){
      location.reload();
    });
  });
}

function onConnect() {
  $("body").addClass("connected");
  removeBlocker();
}

function onDisconnect(err) {
  $("body").removeClass("connected");
  addBlocker();
  if (err != null && err.isGotalkProtocolError) return console.error(err);
}

function onLoggedInConnect() {
  s.request("person.create", {
    firstName: "Fred",
    lastName:  "Flintstone"
  }, function (err, result) {
    if (err) return console.error('create failed:', err);
    console.log('person.create result:', result);

    s.request("person.get", {ID:result.ID}, function (err, result) {
      if (err) return console.error('echo failed:', err);
      console.log('person.get result:', result);
    });
  });
}
