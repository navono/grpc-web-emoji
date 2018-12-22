const {EmojizeRequest, EmojizeReply} = require('./api/v1/emoji_pb');
const {EmojiServiceClient} = require('./api/v1/emoji_grpc_web_pb');

var client = new EmojiServiceClient('http://192.168.99.100:31380');
var editor = document.getElementById('editor');

window.emojize = function() {
  var request = new EmojizeRequest();
  request.setText(editor.innerText);

  client.emojize(request, {}, (err, response) => {
    editor.innerText = response.getEmojizedText();
  });
}