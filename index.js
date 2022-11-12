function getBitches() {
  let newWindow = $('<div class="window"></div>');
  let newTitle = $('<div class="title-bar"></div>');
  let newTitleText = $('<div class="title-bar-text"></div>');
  let newTitleControls = $('<div class="title-bar-controls"></div>');
  let minimize = $('<button arial-label="Minimize"></button>');
  let maximize = $('<button arial-label="Maximize"></button>');
  let close = $('<button arial-label="Close"></button>');
  let newContent = $('<div class="window-body"></div>');

  $(newTitleText).text('Person');
  $(newTitleControls).append(minimize, maximize, close);
  $(newTitle).append(newTitleText, newTitleControls);
  
  $(newContent).text('hello');

  $(newWindow).append(newTitle, newContent);
  $('body').append(newWindow);
}
